
import (
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
)

func foo() {
	ast.New()
	parser.New()
	scanner.New()
	token.New()
}
