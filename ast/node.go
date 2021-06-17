package ast

import (
	"strings"
)

var (
	True  = Bool(true)
	False = Bool(false)
)

func Str(s string) *Node {
	return &Node{
		Token: Node_STR,
		Data:  &Node_S{S: s},
	}
}

func Ident(s string) *Node {
	return &Node{
		Token: Node_IDENT,
		Data:  &Node_S{S: strings.ToLower(s)},
	}
}

func Call(s string) *Node {
	return &Node{
		Token: Node_CALL,
		Data:  &Node_S{S: strings.ToLower(s)},
	}
}

func Int(n int64) *Node {
	return &Node{
		Token: Node_INT,
		Data:  &Node_I{I: n},
	}
}

func Bool(b bool) *Node {
	return &Node{
		Token: Node_BOOL,
		Data:  &Node_B{B: b},
	}
}

func Exit() *Node {
	return &Node{Token: Node_EXIT}
}

func Arr() *Node {
	return &Node{Token: Node_ARR}
}

func (n *Node) Nest(nn ...*Node) *Node {
	n.Nested = append(n.Nested, nn...)
	return n
}
