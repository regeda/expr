package lexer

import "fmt"

type Typ uint8

const (
	typIllegal Typ = iota
	typPths
	typIdent
	// exported types
	TypInt
	TypStr
	TypTrue
	TypFalse
	TypInvoke
	TypVector
	TypOpAdd
	TypOpSub
	TypOpMul
	TypOpDiv
)

var typ2str = [...]string{
	"illegal", "pths", "ident",
	"int", "str", "true", "false", "invoke", "vector",
	"op_add", "op_sub", "op_mul", "op_div",
}

var typMathOp = map[byte]Typ{
	'+': TypOpAdd,
	'-': TypOpSub,
	'*': TypOpMul,
	'/': TypOpDiv,
}

func (t Typ) String() string {
	i := int(t)
	if i < len(typ2str) {
		return typ2str[i]
	}
	return fmt.Sprintf("_unknown_%d", i)
}

type Node struct {
	Typ
	Cap  uint
	DatN int64
	DatS string
}

func (n Node) setTyp(t Typ) Node {
	n.Typ = t
	return n
}

func (n Node) incCap() Node {
	n.Cap++
	return n
}

func (n Node) typeOf(tt ...Typ) bool {
	for _, t := range tt {
		if n.Typ == t {
			return true
		}
	}
	return false
}
