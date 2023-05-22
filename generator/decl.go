package generator

import (
	"go/ast"
	"go/token"
)

/* type (A struct, B struct, ..)の形で、複数の構造体をまとめて宣言する。 */

func GenerateTypeDecl(specs []ast.Spec) ast.Decl {
	return &ast.GenDecl{
		Tok:   token.TYPE, // type句
		Specs: specs,
	}
}
