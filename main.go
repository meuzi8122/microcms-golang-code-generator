package main

import (
	"go/ast"
	"go/format"
	"go/token"
	"os"

	"github.com/meuzi8122/microcms-golang-code-generator/generator"
)

func main() {

	schemas := generator.GenerateSchemas("../ブログスキーマ")

	specs := []ast.Spec{
		generator.GenerateImageTypeSpec(),
		generator.GenerateDomainValueTypeSpec(),
	}

	for _, schema := range schemas {
		specs = append(specs, generator.GenerateTypeSpec(schema))
	}

	dname := "../cms"

	os.Mkdir(dname, 0777)
	f, _ := os.Create(dname + "/model.go")

	fset := token.NewFileSet()

	node := generator.GenerateSourceFile("cms", []ast.Decl{
		generator.GenerateTypeDecl(specs),
	})

	format.Node(f, fset, node)

}
