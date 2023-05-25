package generator

import "go/ast"

func GenerateSourceFile(dname string, pname string) *ast.File {
	schemas := GenerateSchemas(dname)

	specs := []ast.Spec{
		GenerateImageTypeSpec(),
		GenerateDomainValueTypeSpec(),
	}

	for _, schema := range schemas {
		specs = append(specs, GenerateTypeSpec(schema))
	}

	return &ast.File{
		Name: ast.NewIdent(pname),
		Decls: []ast.Decl{
			GenerateTypeDecl(specs),
		},
	}
}
