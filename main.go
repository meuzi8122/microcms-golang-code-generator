package main

import (
	"go/format"
	"go/token"
	"os"

	"github.com/meuzi8122/microcms-golang-code-generator/generator"
)

func main() {
	dname := "../cms"

	os.Mkdir(dname, 0777)
	f, _ := os.Create(dname + "/model.go")

	fset := token.NewFileSet()

	node := generator.GenerateSourceFile("../ブログスキーマ", "cms")

	format.Node(f, fset, node)

}
