package generator

import (
	"go/ast"

	"github.com/gertd/go-pluralize"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateField(name string, kind string) *ast.Field {
	pluralize := pluralize.NewClient()

	var ftype string

	if kind == "text" || kind == "textArea" {
		ftype = "string"
	} else if kind == "number" {
		ftype = "int"
	} else if kind == "boolean" {
		ftype = "bool"
	} else if kind == "image" {
		ftype = "Image"
	} else if kind == "select" {
		ftype = "[]DomainValue"
	} else if kind == "relation" {
		ftype = cases.Title(language.Und).String(pluralize.Singular(name))
	} else if kind == "relationList" {
		ftype = "[]" + cases.Title(language.Und).String(pluralize.Singular(name))
	} else {
		panic("[ERROR] プロパティを生成できませんでした kind = " + kind)
	}

	return &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(name)},
		Type:  ast.NewIdent(ftype),
		Tag:   nil,
	}
}
