package generator

import (
	"go/ast"
)

func GenerateTypeSpec(schema Schema) *ast.TypeSpec {
	var fields []*ast.Field

	for _, afield := range schema.ApiFields {
		fields = append(fields, GenerateField(afield.FieldId, afield.Kind))
	}

	stype := ast.StructType{
		Fields: &ast.FieldList{
			List: fields,
		},
	}

	return &ast.TypeSpec{
		Name: ast.NewIdent(schema.Name),
		Type: &stype,
	}
}

func GenerateImageTypeSpec() *ast.TypeSpec {
	stype := ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				GenerateField("url", "text"),
				GenerateField("width", "number"),
				GenerateField("height", "number"),
			},
		},
	}

	return &ast.TypeSpec{
		Name: ast.NewIdent("Image"),
		Type: &stype,
	}

}

func GenerateDomainValueTypeSpec() *ast.TypeSpec {
	stype := ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				GenerateField("id", "text"),
				GenerateField("value", "text"),
			},
		},
	}

	return &ast.TypeSpec{
		Name: ast.NewIdent("DomainValue"),
		Type: &stype,
	}
}
