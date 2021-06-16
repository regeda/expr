package ast

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/regeda/expr/tokenz"
)

var (
	errTkPunctWrongLen = errors.New("The token Punct should contain 1 byte of data")
)

var (
	valueAfterPunct = []byte{'[', '(', ','}
	commaAfterPunct = []byte{']', ')'}
)

type Builder struct {
	st stack
}

func (b *Builder) Build(tokens []tokenz.Token) (*Node, error) {
	b.st.reset()

	b.st.push(Exit())

	btk := tokenz.Token{Tk: tokenz.TkNone}

	for _, t := range tokens {
		switch t.Tk {
		case tokenz.TkInt:
			v, err := strconv.ParseInt(string(t.Dat), 10, 64)
			if err != nil {
				return nil, err
			}
			if !expectValueAfter(btk) {
				return nil, fmt.Errorf("unexpected integer after %v", btk)
			}
			b.st.nest(Int(v))
		case tokenz.TkStr:
			v, err := strconv.Unquote(string(t.Dat))
			if err != nil {
				return nil, errors.Wrapf(err, "strconv.Unquote %s", t.Dat)
			}
			if !expectValueAfter(btk) {
				return nil, fmt.Errorf("unexpected string after %v", btk)
			}
			b.st.nest(Str(v))
		case tokenz.TkIdent:
			if !expectValueAfter(btk) {
				return nil, fmt.Errorf("unexpected ident after %v", btk)
			}
			b.st.push(Ident(string(t.Dat)))
		case tokenz.TkTrue:
			if !expectValueAfter(btk) {
				return nil, fmt.Errorf("unexpected TRUE after %v", btk)
			}
			b.st.nest(True)
		case tokenz.TkFalse:
			if !expectValueAfter(btk) {
				return nil, fmt.Errorf("unexpected FALSE after %v", btk)
			}
			b.st.nest(False)
		case tokenz.TkPunct:
			if !t.DatLen(1) {
				return nil, errTkPunctWrongLen
			}
			switch t.Dat[0] {
			case '[':
				if !expectValueAfter(btk) {
					return nil, fmt.Errorf("unexpected array after %v", btk)
				}
				b.st.push(b.st.nest(Arr()))
			case ']':
				switch btk.Tk {
				case tokenz.TkInt, tokenz.TkStr, tokenz.TkTrue, tokenz.TkFalse, tokenz.TkPunct:
				default:
					return nil, fmt.Errorf("unexpected array closing after %v", btk)
				}
				n := b.st.pop()
				if n.Token != Node_ARR {
					return nil, fmt.Errorf("stack error: expected array, got %v", n.Token)
				}
			case '(':
				switch btk.Tk {
				case tokenz.TkIdent:
					n := b.st.pop()
					n.Token = Node_CALL
					b.st.push(b.st.nest(n))
				default:
					return nil, fmt.Errorf("unexpected invokation after %v", btk)
				}
			case ')':
				switch btk.Tk {
				case tokenz.TkInt, tokenz.TkStr, tokenz.TkTrue, tokenz.TkFalse, tokenz.TkPunct:
				default:
					return nil, fmt.Errorf("unexpected invokation closing after %v", btk)
				}
				n := b.st.pop()
				if n.Token != Node_CALL {
					return nil, fmt.Errorf("stack error: expected invokation, got %v", n.Token)
				}
			case ',':
				switch btk.Tk {
				case tokenz.TkInt, tokenz.TkStr, tokenz.TkTrue, tokenz.TkFalse:
					n := b.st.top()
					if n.Token != Node_CALL && n.Token != Node_ARR {
						return nil, fmt.Errorf("unexpected comma after %v", btk)
					}
				case tokenz.TkPunct:
					if bytes.IndexByte(commaAfterPunct, btk.Dat[0]) == -1 {
						return nil, fmt.Errorf("unexpected comma after %v", btk)
					}
				default:
					return nil, fmt.Errorf("unexpected comma after %v", btk)
				}
			default:
				return nil, fmt.Errorf("unexpected punct %s after %v", t.Dat, btk)
			}
		default:
			return nil, fmt.Errorf("unexpected token %v", t)
		}

		btk = t
	}

	if b.st.len() != 1 {
		return nil, fmt.Errorf("unexpected stack length %v", b.st.len())
	}

	return b.st[0], nil
}

func expectValueAfter(t tokenz.Token) bool {
	switch t.Tk {
	case tokenz.TkNone:
		return true
	case tokenz.TkPunct:
		return bytes.IndexByte(valueAfterPunct, t.Dat[0]) != -1
	default:
		return false
	}
}
