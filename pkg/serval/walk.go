package serval

import (
	"github.com/microsoft/typescript-go/internal/ast"
)

// WalkAST visits every node in the AST tree, calling fn for each.
// If fn returns true, traversal stops.
func WalkAST(node *ast.Node, fn func(node *ast.Node) bool) {
	if fn(node) {
		return
	}
	node.ForEachChild(func(child *ast.Node) bool {
		WalkAST(child, fn)
		return false
	})
}

// FindNodes returns all nodes matching the predicate.
func FindNodes(root *ast.Node, predicate func(*ast.Node) bool) []*ast.Node {
	var results []*ast.Node
	WalkAST(root, func(node *ast.Node) bool {
		if predicate(node) {
			results = append(results, node)
		}
		return false
	})
	return results
}

// Re-export AST helper functions for node type checking.
var (
	IsVariableStatement    = ast.IsVariableStatement
	IsFunctionDeclaration  = ast.IsFunctionDeclaration
	IsExportDeclaration    = ast.IsExportDeclaration
	IsImportDeclaration    = ast.IsImportDeclaration
	IsCallExpression       = ast.IsCallExpression
	IsIdentifier           = ast.IsIdentifier
	IsStringLiteral        = ast.IsStringLiteral
	IsObjectLiteralExpression = ast.IsObjectLiteralExpression
	IsPropertyAssignment   = ast.IsPropertyAssignment
	IsArrowFunction        = ast.IsArrowFunction
	IsTypeAliasDeclaration = ast.IsTypeAliasDeclaration
	IsInterfaceDeclaration = ast.IsInterfaceDeclaration
	IsClassDeclaration     = ast.IsClassDeclaration
)
