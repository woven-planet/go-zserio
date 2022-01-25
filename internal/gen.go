//go:generate antlr -Dlanguage=Go -visitor -no-listener -package parser -o . parser/ZserioLexer.g4 parser/ZserioParser.g4

package internal
