package visitor

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/parser"
)

var ErrImplicitLengthArraysNotSupported = errors.New("implicit length arrays are obsolete, and not supported")
var ErrAlignmentMustBeInteger = errors.New("alignment expression must be an integer")

type Visitor struct {
	parser.BaseZserioParserVisitor
}

func (c *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		fmt.Printf("syntax error near '%s'", t.GetText())
	default:
		return tree.Accept(c)
	}

	return fmt.Errorf("visit result not of a Node")
}

func (c *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	var children []interface{}

	for _, n := range node.GetChildren() {
		cr := c.Visit(n.(antlr.ParseTree))
		if err, isErr := cr.(error); isErr {
			return err
		}
		children = append(children, cr)
	}
	return children
}

func (v *Visitor) VisitPackageDeclaration(ctx *parser.PackageDeclarationContext) interface{} {
	pkg := ast.NewPackage(v.Visit(ctx.PackageNameDefinition().(antlr.ParseTree)).(string))
	pkg.Comment = getComment(ctx.GetParser(), ctx.GetStart())

	for _, imp := range ctx.AllImportDeclaration() {
		impNode := v.Visit(imp.(antlr.ParseTree))
		if i, ok := impNode.(*ast.Import); ok {
			pkg.Imports = append(pkg.Imports, i)
		}
	}

	for _, dir := range ctx.AllLanguageDirective() {
		dirNode := v.Visit(dir.(antlr.ParseTree))
		switch n := dirNode.(type) {
		case *ast.Const:
			pkg.Consts[n.Name] = n
		case *ast.Subtype:
			pkg.Subtypes[n.Name] = n
		case *ast.InstantiateType:
			pkg.InstantiatedTypes[n.Name] = n
		case *ast.Struct:
			pkg.Structs[n.Name] = n
		case *ast.Enum:
			pkg.Enums[n.Name] = n
		case *ast.Union:
			pkg.Unions[n.Name] = n
		case *ast.BitmaskType:
			pkg.Bitmasks[n.Name] = n
		case *ast.Choice:
			pkg.Choices[n.Name] = n
		}
	}
	return pkg
}

func (v *Visitor) VisitPackageNameDefinition(ctx *parser.PackageNameDefinitionContext) interface{} {
	// It would be nice of the grammar just reused qualifiedName here, but for some
	// reason they opted for "id (DOT id)*" again.
	id_nodes := ctx.AllId()
	var ids []string
	for _, n := range id_nodes {
		ids = append(ids, v.Visit(n.(antlr.ParseTree)).(string))
	}
	return strings.Join(ids, ".")
}

func getComment(p antlr.Parser, t antlr.Token) string {
	stream := p.GetTokenStream().(*antlr.CommonTokenStream)
	hidden := stream.GetHiddenTokensToLeft(t.GetTokenIndex(), antlr.TokenHiddenChannel)
	var comments []string

out:
	for _, hiddenToken := range hidden {
		switch hiddenToken.GetTokenType() {
		case parser.ZserioParserDOC_COMMENT, parser.ZserioParserMARKDOWN_COMMENT, parser.ZserioParserBLOCK_COMMENT, parser.ZserioParserLINE_COMMENT:
			comments = append(comments, hiddenToken.GetText())
		default:
			break out
		}
	}
	return strings.Join(comments, "\n")
}

func (v *Visitor) VisitImportDeclaration(ctx *parser.ImportDeclarationContext) interface{} {
	children := ctx.GetChildren()
	parts := []string{}
	for i := 1; i < len(children)-2; i++ {
		parts = append(parts, children[i].(antlr.ParseTree).GetText())
	}

	symbolName := parts[len(parts)-1]
	if ctx.MULTIPLY() != nil {
		symbolName = "*"
	}
	return &ast.Import{
		Package: strings.Join(parts[:len(parts)-1], ""),
		Type:    symbolName,
	}
}

func (v *Visitor) VisitLanguageDirective(ctx *parser.LanguageDirectiveContext) interface{} {
	// There can be only one child - return that directly.
	return v.VisitChildren(ctx).([]interface{})[0]
}

func (v *Visitor) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	// There can be only one child - return that directly.
	return v.VisitChildren(ctx).([]interface{})[0]
}

func (v *Visitor) VisitSymbolDefinition(ctx *parser.SymbolDefinitionContext) interface{} {
	// There can be only one child - return that directly.
	return v.VisitChildren(ctx).([]interface{})[0]
}

func (v *Visitor) VisitConstDefinition(ctx *parser.ConstDefinitionContext) interface{} {
	return &ast.Const{
		Name:            ctx.Id().GetText(),
		Comment:         getComment(ctx.GetParser(), ctx.GetStart()),
		Type:            v.VisitTypeInstantiation(ctx.TypeInstantiation().(*parser.TypeInstantiationContext)).(*ast.TypeReference),
		ValueExpression: v.Visit(ctx.Expression()).(*ast.Expression),
	}
}

func (v *Visitor) VisitSubtypeDeclaration(ctx *parser.SubtypeDeclarationContext) interface{} {
	return &ast.Subtype{
		Name:    ctx.Id().GetText(),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
		Type:    v.VisitTypeReference(ctx.TypeReference().(*parser.TypeReferenceContext)).(*ast.TypeReference),
	}
}

func (v *Visitor) VisitStructureDeclaration(ctx *parser.StructureDeclarationContext) interface{} {
	typ := &ast.Struct{
		Name:    v.Visit(ctx.Id().(antlr.ParseTree)).(string),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
	}

	if tp := ctx.TemplateParameters(); tp != nil {
		typ.TemplateParameters = v.Visit(tp.(antlr.ParseTree)).([]string)
	}

	if tp := ctx.TypeParameters(); tp != nil {
		typ.TypeParameters = v.VisitTypeParameters(tp.(*parser.TypeParametersContext)).([]*ast.Parameter)
	}

	for _, fieldContext := range ctx.AllStructureFieldDefinition() {
		typ.Fields = append(typ.Fields, v.Visit(fieldContext.(antlr.ParseTree)).(*ast.Field))
	}

	for _, functionContext := range ctx.AllFunctionDefinition() {
		typ.Functions = append(typ.Functions, v.VisitFunctionDefinition(functionContext.(*parser.FunctionDefinitionContext)).(*ast.Function))
	}

	return typ
}

func (v *Visitor) VisitStructureFieldDefinition(ctx *parser.StructureFieldDefinitionContext) interface{} {
	field := v.Visit(ctx.FieldTypeId().(antlr.ParseTree)).(*ast.Field)
	field.Comment = getComment(ctx.GetParser(), ctx.GetStart())
	field.IsOptional = ctx.OPTIONAL() != nil
	if a := ctx.FieldAlignment(); a != nil {
		a, ok := v.Visit(a.(antlr.ParseTree)).(uint8)
		if !ok {
			return ErrAlignmentMustBeInteger
		}
		field.Alignment = a
	}

	if ctx.FieldInitializer() != nil {
		field.Initializer = v.VisitFieldInitializer(ctx.FieldInitializer().(*parser.FieldInitializerContext)).(*ast.Expression)
	}

	if ctx.FieldOptionalClause() != nil {
		field.OptionalClause = v.VisitFieldOptionalClause(ctx.FieldOptionalClause().(*parser.FieldOptionalClauseContext)).(*ast.Expression)
	}
	if ctx.FieldConstraint() != nil {
		field.Constraint = v.VisitFieldConstraint(ctx.FieldConstraint().(*parser.FieldConstraintContext)).(*ast.Expression)
	}
	return field
}

func (v *Visitor) VisitFieldAlignment(ctx *parser.FieldAlignmentContext) interface{} {
	r := v.Visit(ctx.Expression().(antlr.ParseTree))
	expr, ok := r.(*ast.Expression)
	if !ok {
		return r
	}
	if expr.Type != parser.ZserioParserDECIMAL_LITERAL {
		return ErrAlignmentMustBeInteger
	}
	value, err := strconv.ParseUint(expr.Text, 10, 8)
	if err != nil {
		return ErrAlignmentMustBeInteger
	}
	if value > math.MaxUint8 {
		return ErrAlignmentMustBeInteger
	}
	return uint8(value)
}

func (v *Visitor) VisitFieldInitializer(ctx *parser.FieldInitializerContext) interface{} {
	return v.Visit(ctx.Expression()).(*ast.Expression)
}

func (v *Visitor) VisitFieldOptionalClause(ctx *parser.FieldOptionalClauseContext) interface{} {
	return v.Visit(ctx.Expression()).(*ast.Expression)
}

func (v *Visitor) VisitFieldConstraint(ctx *parser.FieldConstraintContext) interface{} {
	return v.Visit(ctx.Expression()).(*ast.Expression)
}

func (v *Visitor) VisitParameterDefinition(ctx *parser.ParameterDefinitionContext) interface{} {
	return &ast.Parameter{
		Name: ctx.Id().GetText(),
		Type: v.VisitTypeReference(ctx.TypeReference().(*parser.TypeReferenceContext)).(*ast.TypeReference),
	}
}

func (v *Visitor) VisitFunctionType(ctx *parser.FunctionTypeContext) interface{} {
	return v.VisitTypeReference(ctx.TypeReference().(*parser.TypeReferenceContext))
}

func (v *Visitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) interface{} {
	return &ast.Function{
		Name:       ctx.FunctionName().GetText(),
		Comment:    getComment(ctx.GetParser(), ctx.GetStart()),
		Result:     v.Visit(ctx.FunctionBody().(*parser.FunctionBodyContext).Expression()).(*ast.Expression),
		ReturnType: v.VisitFunctionType(ctx.FunctionType().(*parser.FunctionTypeContext)).(*ast.TypeReference),
	}
}

func (v *Visitor) VisitTypeParameters(ctx *parser.TypeParametersContext) interface{} {
	var typeParams []*ast.Parameter
	for _, paramDefinition := range ctx.AllParameterDefinition() {
		typeParams = append(typeParams, v.VisitParameterDefinition(paramDefinition.(*parser.ParameterDefinitionContext)).(*ast.Parameter))
	}
	return typeParams
}

func (v *Visitor) VisitChoiceCase(ctx *parser.ChoiceCaseContext) interface{} {
	return &ast.ChoiceCaseExpression{
		Condition: v.Visit(ctx.Expression()).(*ast.Expression),
		Comment:   getComment(ctx.GetParser(), ctx.GetStart()),
	}
}

func (v *Visitor) VisitChoiceCases(ctx *parser.ChoiceCasesContext) interface{} {
	typ := &ast.ChoiceCase{}

	for _, choiceCaseContext := range ctx.AllChoiceCase() {
		typ.Conditions = append(typ.Conditions, v.VisitChoiceCase(choiceCaseContext.(*parser.ChoiceCaseContext)).(*ast.ChoiceCaseExpression))
	}

	if tp := ctx.ChoiceFieldDefinition(); tp != nil {
		typ.Field = v.VisitChoiceFieldDefinition(tp.(*parser.ChoiceFieldDefinitionContext)).(*ast.Field)
	}
	return typ
}

func (v *Visitor) VisitChoiceDefault(ctx *parser.ChoiceDefaultContext) interface{} {
	typ := &ast.ChoiceCase{}
	if tp := ctx.ChoiceFieldDefinition(); tp != nil {
		typ.Field = v.VisitChoiceFieldDefinition(tp.(*parser.ChoiceFieldDefinitionContext)).(*ast.Field)
	}
	return typ
}

func (v *Visitor) VisitChoiceFieldDefinition(ctx *parser.ChoiceFieldDefinitionContext) interface{} {
	return v.VisitFieldTypeId(ctx.FieldTypeId().(*parser.FieldTypeIdContext)).(*ast.Field)
}

func (v *Visitor) VisitChoiceDeclaration(ctx *parser.ChoiceDeclarationContext) interface{} {
	typ := &ast.Choice{
		Name:       v.Visit(ctx.Id().(antlr.ParseTree)).(string),
		Comment:    getComment(ctx.GetParser(), ctx.GetStart()),
		Expression: v.Visit(ctx.Expression().(antlr.ParseTree)).(*ast.Expression),
	}

	if tp := ctx.TemplateParameters(); tp != nil {
		typ.TemplateParameters = v.Visit(tp.(antlr.ParseTree)).([]string)
	}

	if tp := ctx.TypeParameters(); tp != nil {
		typ.TypeParameters = v.VisitTypeParameters(tp.(*parser.TypeParametersContext)).([]*ast.Parameter)
	}

	// visit all choice cases
	for _, choiceCasesContext := range ctx.AllChoiceCases() {
		typ.Cases = append(typ.Cases, v.VisitChoiceCases(choiceCasesContext.(*parser.ChoiceCasesContext)).(*ast.ChoiceCase))
	}

	// visit the default choice
	if ctx.ChoiceDefault() != nil {
		typ.DefaultCase = v.VisitChoiceDefault(ctx.ChoiceDefault().(*parser.ChoiceDefaultContext)).(*ast.ChoiceCase)
	}

	for _, functionContext := range ctx.AllFunctionDefinition() {
		typ.Functions = append(typ.Functions, v.VisitFunctionDefinition(functionContext.(*parser.FunctionDefinitionContext)).(*ast.Function))
	}

	return typ
}

func (v *Visitor) VisitUnionDeclaration(ctx *parser.UnionDeclarationContext) interface{} {
	typ := &ast.Union{
		Name:    ctx.Id().GetText(),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
	}

	if tp := ctx.TemplateParameters(); tp != nil {
		typ.TemplateParameters = v.Visit(tp.(antlr.ParseTree)).([]string)
	}

	if tp := ctx.TypeParameters(); tp != nil {
		typ.TypeParameters = v.VisitTypeParameters(tp.(*parser.TypeParametersContext)).([]*ast.Parameter)
	}

	for _, unionField := range ctx.AllUnionFieldDefinition() {
		unionFieldDefinitionContext := unionField.(*parser.UnionFieldDefinitionContext)
		choiceFieldDefinitionContext := unionFieldDefinitionContext.ChoiceFieldDefinition().(*parser.ChoiceFieldDefinitionContext)
		typ.Fields = append(typ.Fields, v.VisitChoiceFieldDefinition(choiceFieldDefinitionContext).(*ast.Field))
	}

	for _, functionContext := range ctx.AllFunctionDefinition() {
		typ.Functions = append(typ.Functions, v.VisitFunctionDefinition(functionContext.(*parser.FunctionDefinitionContext)).(*ast.Function))
	}
	return typ
}

func (v *Visitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	typ := &ast.Enum{
		Name:    ctx.Id().GetText(),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
		Type:    v.VisitTypeInstantiation(ctx.TypeInstantiation().(*parser.TypeInstantiationContext)).(*ast.TypeReference),
	}
	for _, enumItem := range ctx.AllEnumItem() {
		typ.Items = append(typ.Items, v.VisitEnumItem(enumItem.(*parser.EnumItemContext)).(*ast.EnumItem))
	}
	return typ
}

func (v *Visitor) VisitEnumItem(ctx *parser.EnumItemContext) interface{} {

	typ := &ast.EnumItem{
		Name:    ctx.Id().GetText(),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
	}

	// an enum item can have an optional expression, specifying the value
	if ctx.Expression() != nil {
		typ.Expression = v.Visit(ctx.Expression()).(*ast.Expression)
	}

	return typ
}

func (v *Visitor) VisitInstantiateDeclaration(ctx *parser.InstantiateDeclarationContext) interface{} {
	return &ast.InstantiateType{
		Name:    v.Visit(ctx.Id().(antlr.ParseTree)).(string),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
		Type:    v.VisitTypeReference(ctx.TypeReference().(*parser.TypeReferenceContext)).(*ast.TypeReference),
	}
}

func (v *Visitor) VisitBitmaskDeclaration(ctx *parser.BitmaskDeclarationContext) interface{} {
	typ := &ast.BitmaskType{
		Name:    ctx.Id().GetText(),
		Comment: getComment(ctx.GetParser(), ctx.GetStart()),
		Type:    v.VisitTypeInstantiation(ctx.TypeInstantiation().(*parser.TypeInstantiationContext)).(*ast.TypeReference),
	}
	for _, bitmaskValue := range ctx.AllBitmaskValue() {
		typ.Values = append(typ.Values, v.VisitBitmaskValue(bitmaskValue.(*parser.BitmaskValueContext)).(*ast.BitmaskValue))
	}
	return typ
}

func (v *Visitor) VisitBitmaskValue(ctx *parser.BitmaskValueContext) interface{} {
	typ := &ast.BitmaskValue{
		Name: ctx.Id().GetText(),
	}

	// a bitmask value can have an optional expression, specifying the value
	if ctx.Expression() != nil {
		typ.Expression = v.Visit(ctx.Expression()).(*ast.Expression)
	}

	return typ
}

func (v *Visitor) VisitParenthesizedExpression(ctx *parser.ParenthesizedExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitArrayExpression(ctx *parser.ArrayExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitDotExpression(ctx *parser.DotExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Operand2: &ast.Expression{
			Type: parser.UnevaluatableExpressionType,
			Text: ctx.Id().GetText(),
		},
		Type: ctx.GetOperator().GetTokenType(),
		Text: ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitLengthofExpression(ctx *parser.LengthofExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitValueofExpression(ctx *parser.ValueofExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitNumbitsExpression(ctx *parser.NumbitsExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression()).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	expr := &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}

	// right-shift token is not correctly identified as right shift, but as
	//"greater then". If this is the case, choose the correct token.
	if ctx.GetOperator().GetTokenType() == parser.ZserioParserGT {
		expr.Type = parser.ZserioParserRSHIFT
	}
	return expr
}
func (v *Visitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}
func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}
func (v *Visitor) VisitBitwiseAndExpression(ctx *parser.BitwiseAndExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}
func (v *Visitor) VisitBitwiseXorExpression(ctx *parser.BitwiseXorExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}
func (v *Visitor) VisitBitwiseOrExpression(ctx *parser.BitwiseOrExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}
func (v *Visitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}
func (v *Visitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitTernaryExpression(ctx *parser.TernaryExpressionContext) interface{} {
	return &ast.Expression{
		Operand1: v.Visit(ctx.Expression(0)).(*ast.Expression),
		Operand2: v.Visit(ctx.Expression(1)).(*ast.Expression),
		Operand3: v.Visit(ctx.Expression(2)).(*ast.Expression),
		Type:     ctx.GetOperator().GetTokenType(),
		Text:     ctx.GetOperator().GetText(),
	}
}

func (v *Visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) interface{} {
	token := ctx.Literal().GetStart()
	return &ast.Expression{
		Type: token.GetTokenType(),
		Text: token.GetText(),
	}
}

func (v *Visitor) VisitIndexExpression(ctx *parser.IndexExpressionContext) interface{} {
	return &ast.Expression{
		Type: ctx.INDEX().GetSymbol().GetTokenType(),
		Text: ctx.INDEX().GetText(),
	}
}

func (v *Visitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) interface{} {
	return &ast.Expression{
		Type: ctx.Id().GetStart().GetTokenType(),
		Text: ctx.Id().GetText(),
	}
}

func (v *Visitor) VisitFieldTypeId(ctx *parser.FieldTypeIdContext) interface{} {
	field := &ast.Field{
		Name: ctx.Id().GetText(),
		Type: v.Visit(ctx.TypeInstantiation().(antlr.ParseTree)).(*ast.TypeReference),
	}

	if ctx.FieldArrayRange() != nil {
		field.Array = &ast.Array{}
		expressionContext := ctx.FieldArrayRange().(*parser.FieldArrayRangeContext).Expression()
		if expressionContext != nil {
			field.Array.Length = v.Visit(expressionContext).(*ast.Expression)
		}
		if ctx.IMPLICIT() != nil {
			// http://zserio.org/doc/ZserioLanguageOverview.html#implicit-length-arrays
			return ErrImplicitLengthArraysNotSupported
		}
		if ctx.PACKED() != nil {
			field.Array.IsPacked = true
		}
	}
	return field
}

func (v *Visitor) VisitTypeReference(ctx *parser.TypeReferenceContext) interface{} {
	if t := ctx.BuiltinType(); t != nil {
		typ := &ast.TypeReference{
			IsBuiltin: true,
			Name:      t.GetText(),
		}
		if ix := strings.Index(typ.Name, ":"); ix != -1 {
			bits, _ := strconv.ParseUint(typ.Name[ix+1:], 10, 8)
			typ.Name = typ.Name[:ix]
			typ.Bits = int(bits)
			if typ.Bits == 8 || typ.Bits == 16 || typ.Bits == 32 || typ.Bits == 64 {
				switch typ.Name {
				case "bit":
					typ.Name = fmt.Sprintf("uint%d", typ.Bits)
					typ.Bits = 0
				case "uint", "int":
					typ.Name = fmt.Sprintf("%s%d", typ.Name, typ.Bits)
					typ.Bits = 0
				}
			}
		}
		return typ
	}

	var templateArguments []*ast.TypeReference
	if ta := ctx.TemplateArguments(); ta != nil {
		templateArguments = v.Visit(ta.(antlr.ParseTree)).([]*ast.TypeReference)
	}
	name := v.Visit(ctx.QualifiedName().(antlr.ParseTree)).(string)
	pkg := ""
	if ix := strings.LastIndex(name, "."); ix != -1 {
		pkg = name[:ix]
		name = name[ix+1:]
	}
	return &ast.TypeReference{
		IsBuiltin:         false,
		Package:           pkg,
		Name:              name,
		TemplateArguments: templateArguments,
	}
}

func (v *Visitor) VisitDynamicLengthArgument(ctx *parser.DynamicLengthArgumentContext) interface{} {
	return v.Visit(ctx.Expression().(antlr.ParseTree)).(*ast.Expression)
}

func (v *Visitor) VisitTypeInstantiation(ctx *parser.TypeInstantiationContext) interface{} {
	typ := v.Visit(ctx.TypeReference().(antlr.ParseTree)).(*ast.TypeReference)
	if ctx.TypeArguments() != nil {
		typ.TypeArguments = v.VisitTypeArguments(ctx.TypeArguments().(*parser.TypeArgumentsContext)).([]*ast.Expression)
	}
	if ctx.DynamicLengthArgument() != nil {
		typ.LengthExpression = v.Visit(ctx.DynamicLengthArgument().(antlr.ParseTree)).(*ast.Expression)
	}
	return typ
}

func (v *Visitor) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	var params []*ast.Expression
	for _, ta := range ctx.AllTypeArgument() {
		params = append(params, v.VisitTypeArgument(ta.(*parser.TypeArgumentContext)).(*ast.Expression))
	}
	return params
}

func (v *Visitor) VisitTypeArgument(ctx *parser.TypeArgumentContext) interface{} {
	if ctx.EXPLICIT() != nil {
		token := ctx.Id().GetStart()
		return &ast.Expression{
			Type: token.GetTokenType(),
			Text: token.GetText(),
		}
		/*
				Type:     ctx.GetOperator().GetTokenType(),
				Text:     ctx.GetOperator().GetText(),
			}
			final AstLocation location = new AstLocation(ctx.getStart());
			final Token expressionToken = ctx.id().ID().getSymbol();

			return new Expression(location, currentPackage, expressionToken.getType(),
					expressionToken.getText(), Expression.ExpressionFlag.IS_EXPLICIT);
		*/
	} else {
		return v.Visit(ctx.Expression()).(*ast.Expression)
	}
}

func (v *Visitor) VisitTemplateArguments(ctx *parser.TemplateArgumentsContext) interface{} {
	var arguments []*ast.TypeReference
	for _, ta := range ctx.AllTemplateArgument() {
		arguments = append(arguments, v.Visit(ta.(antlr.ParseTree)).(*ast.TypeReference))
	}
	return arguments
}

func (v *Visitor) VisitTemplateArgument(ctx *parser.TemplateArgumentContext) interface{} {
	return v.Visit(ctx.TypeReference().(antlr.ParseTree))
}

func (v *Visitor) VisitTemplateParameters(ctx *parser.TemplateParametersContext) interface{} {
	var ids []string
	for _, id := range ctx.AllId() {
		ids = append(ids, id.GetText())
	}
	return ids
}

func (v *Visitor) VisitId(ctx *parser.IdContext) interface{} {
	id_node := ctx.GetChild(0)
	return v.Visit(id_node.(antlr.ParseTree))
}

func (v *Visitor) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("Error!")
	// return nil
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}
