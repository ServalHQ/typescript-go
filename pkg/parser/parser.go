// Package parser re-exports the TypeScript parser from internal/parser.
package parser

import (
	"github.com/microsoft/typescript-go/internal/parser"
	"github.com/microsoft/typescript-go/pkg/ast"
	"github.com/microsoft/typescript-go/pkg/core"
)

// ParseSourceFile parses a TypeScript or JavaScript source file and returns the AST.
func ParseSourceFile(opts ast.SourceFileParseOptions, sourceText string, scriptKind core.ScriptKind) *ast.SourceFile {
	return parser.ParseSourceFile(opts, sourceText, scriptKind)
}

// ParseIsolatedEntityName parses a qualified name (e.g., "A.B.C") in isolation.
func ParseIsolatedEntityName(text string) *ast.Node {
	return parser.ParseIsolatedEntityName(text)
}
