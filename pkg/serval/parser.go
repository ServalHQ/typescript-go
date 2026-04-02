// Package serval provides thin public wrappers around typescript-go internals.
//
// This package exists in the ServalHQ fork to expose internal APIs for use
// in the Serval backend. Upstream typescript-go keeps all packages internal.
//
// RULES:
// - Never modify files in internal/ — keep upstream mergeable
// - Only add files in pkg/serval/ — our additions
// - Sync upstream: git fetch upstream && git merge upstream/main
package serval

import (
	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/parser"
)

// ParseSourceFile parses TypeScript/JavaScript source code into an AST.
func ParseSourceFile(fileName string, sourceText string) *ast.SourceFile {
	opts := ast.SourceFileParseOptions{
		FileName: fileName,
	}
	return parser.ParseSourceFile(opts, sourceText, core.ScriptKindTS)
}

// ParseTSX parses a TSX file into an AST.
func ParseTSX(fileName string, sourceText string) *ast.SourceFile {
	opts := ast.SourceFileParseOptions{
		FileName: fileName,
	}
	return parser.ParseSourceFile(opts, sourceText, core.ScriptKindTSX)
}

// ParseJS parses a JavaScript file into an AST.
func ParseJS(fileName string, sourceText string) *ast.SourceFile {
	opts := ast.SourceFileParseOptions{
		FileName: fileName,
	}
	return parser.ParseSourceFile(opts, sourceText, core.ScriptKindJS)
}

// Re-export key types so consumers don't need to import internal/ directly.
type (
	SourceFile = ast.SourceFile
	Node       = ast.Node
)
