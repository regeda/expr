package tokenz

import "strconv"

type Tk uint8

const (
	TkNone Tk = iota
	TkInt
	TkStr
	TkTrue
	TkFalse
	TkIdent
	TkPunct
)

var tk2str = [...]string{
	"none", "int", "str", "true", "false", "ident", "punct",
}

type Token struct {
	Tk
	Dat []byte
}

func (t *Token) DatLen(n int) bool {
	return len(t.Dat) == n
}

func (t Token) String() string {
	i := int(t.Tk)
	var s string
	if i < len(tk2str) {
		s = tk2str[i]
	} else {
		s = "_unknown_" + strconv.Itoa(i)
	}
	if len(t.Dat) > 0 {
		s += "=" + string(t.Dat)
	}
	return s
}
