package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var ExplicitExitAnalyzer = &analysis.Analyzer{
	Name: "explicitexit",
	Doc:  "check for explicit exit calls",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	expr := func(x *ast.ExprStmt) {
		// check that the expression represents a function call
		// with an explicit exit
		if call, ok := x.X.(*ast.CallExpr); ok {
			if isExplitExit(call) {
				pass.Reportf(x.Pos(), "explicit exit call")
			}
		}
	}
	for _, file := range pass.Files {
		if file.Name.String() != "main" {
			continue
		}
		// using the ast.Inspect function, we go through all the nodes of the AST
		ast.Inspect(file, func(node ast.Node) bool {
			if f, ok := node.(*ast.FuncDecl); ok {
				if f.Name.Name == "main" {
					if x, ok := node.(*ast.ExprStmt); ok {
						expr(x)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func isExplitExit(call *ast.CallExpr) bool {
	// check that the function call is exit
	if id, ok := call.Fun.(*ast.Ident); ok {
		if id.Name == "exit" {
			return true
		}
	}
	return false
}
