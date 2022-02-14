// Code generated from /private/var/tmp/_bazel_ignas.anikevicius/c3129ead79ca509099c649d74caf832e/sandbox/darwin-sandbox/2595/execroot/__main__/internal/parser/ZserioParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // ZserioParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseZserioParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseZserioParserVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitPackageNameDefinition(ctx *PackageNameDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitLanguageDirective(ctx *LanguageDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSymbolDefinition(ctx *SymbolDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitConstDefinition(ctx *ConstDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitRuleGroupDefinition(ctx *RuleGroupDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitRuleDefinition(ctx *RuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSubtypeDeclaration(ctx *SubtypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitStructureDeclaration(ctx *StructureDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitStructureFieldDefinition(ctx *StructureFieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldAlignment(ctx *FieldAlignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldOffset(ctx *FieldOffsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldTypeId(ctx *FieldTypeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldArrayRange(ctx *FieldArrayRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldInitializer(ctx *FieldInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldOptionalClause(ctx *FieldOptionalClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFieldConstraint(ctx *FieldConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitChoiceDeclaration(ctx *ChoiceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitChoiceCases(ctx *ChoiceCasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitChoiceCase(ctx *ChoiceCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitChoiceDefault(ctx *ChoiceDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitChoiceFieldDefinition(ctx *ChoiceFieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitUnionDeclaration(ctx *UnionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitUnionFieldDefinition(ctx *UnionFieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitEnumItem(ctx *EnumItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBitmaskDeclaration(ctx *BitmaskDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBitmaskValue(ctx *BitmaskValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlTableDeclaration(ctx *SqlTableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlTableFieldDefinition(ctx *SqlTableFieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlConstraintDefinition(ctx *SqlConstraintDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlConstraint(ctx *SqlConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlWithoutRowId(ctx *SqlWithoutRowIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlDatabaseDefinition(ctx *SqlDatabaseDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitSqlDatabaseFieldDefinition(ctx *SqlDatabaseFieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitServiceDefinition(ctx *ServiceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitServiceMethodDefinition(ctx *ServiceMethodDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitPubsubDefinition(ctx *PubsubDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitPubsubMessageDefinition(ctx *PubsubMessageDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTopicDefinition(ctx *TopicDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitParameterDefinition(ctx *ParameterDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTemplateParameters(ctx *TemplateParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTemplateArguments(ctx *TemplateArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTemplateArgument(ctx *TemplateArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitInstantiateDeclaration(ctx *InstantiateDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBitwiseXorExpression(ctx *BitwiseXorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitDotExpression(ctx *DotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitValueofExpression(ctx *ValueofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitArrayExpression(ctx *ArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitNumbitsExpression(ctx *NumbitsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitLengthofExpression(ctx *LengthofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBitwiseOrExpression(ctx *BitwiseOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBitwiseAndExpression(ctx *BitwiseAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitIndexExpression(ctx *IndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTypeInstantiation(ctx *TypeInstantiationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBuiltinType(ctx *BuiltinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitDynamicLengthArgument(ctx *DynamicLengthArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitVarintType(ctx *VarintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFixedBitFieldType(ctx *FixedBitFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitDynamicBitFieldType(ctx *DynamicBitFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitFloatType(ctx *FloatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZserioParserVisitor) VisitExternType(ctx *ExternTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
