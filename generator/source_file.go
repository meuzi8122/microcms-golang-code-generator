package generator

import "go/ast"

func GenerateSourceFile(name string, decls []ast.Decl) *ast.File {
	return &ast.File{
		Name:  ast.NewIdent(name),
		Decls: decls,
	}
}
