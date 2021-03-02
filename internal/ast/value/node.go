package value

import (
	"strconv"
	"strings"

	"github.com/regeda/expr/internal/ast"
)

func unescapeDquote(s string) string {
	return strings.Replace(s, `\"`, `"`, -1)
}

func Str(s string) *ast.Node {
	return &ast.Node{
		Token: ast.Node_STR,
		Data:  &ast.Node_S{S: unescapeDquote(s)},
	}
}

func Call(s string) *ast.Node {
	return &ast.Node{
		Token: ast.Node_CALL,
		Data:  &ast.Node_S{S: strings.ToLower(s)},
	}
}

func Atoi(s string) *ast.Node {
	n, _ := strconv.ParseInt(s, 10, 64)
	return Int(n)
}

func Int(n int64) *ast.Node {
	return &ast.Node{
		Token: ast.Node_INT,
		Data:  &ast.Node_I{I: n},
	}
}

func boolVal(b bool) *ast.Node {
	return &ast.Node{
		Token: ast.Node_BOOL,
		Data:  &ast.Node_B{B: b},
	}
}

var (
	True  = boolVal(true)
	False = boolVal(false)
)

func Bool(b bool) *ast.Node {
	if b {
		return True
	}
	return False
}

func Exit() *ast.Node {
	return &ast.Node{Token: ast.Node_EXIT}
}

func Arr() *ast.Node {
	return &ast.Node{Token: ast.Node_ARR}
}

func Nest(n *ast.Node, children ...*ast.Node) *ast.Node {
	n.Nested = append(n.Nested, children...)
	return n
}
