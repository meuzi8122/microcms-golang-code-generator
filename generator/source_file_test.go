package generator

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSourceFile(t *testing.T) {
	t.Run("生成された構造体をチェックする", func(t *testing.T) {
		node := GenerateSourceFile("../testdata/schema", "cms")
		assert.Equal(t, node.Name, ast.NewIdent("cms"))
	})
}
