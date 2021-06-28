package lexer

import "strings"

type ast struct {
	rpn, ops, caps nvec
}

func (a *ast) reset() {
	a.rpn.reset()
	a.ops.reset()
	a.caps.reset()
}

func (a *ast) pushInt(n int64) {
	a.rpn.push(Node{Typ: TypInt, DatN: n})
}

func (a *ast) pushStr(s string) {
	a.rpn.push(Node{Typ: TypStr, DatS: s})
}

func (a *ast) pushIdent(s string) {
	a.rpn.push(Node{Typ: typIdent, DatS: strings.ToLower(s)})
}

func (a *ast) pushTrue() {
	a.rpn.push(Node{Typ: TypTrue})
}

func (a *ast) pushFalse() {
	a.rpn.push(Node{Typ: TypFalse})
}

func (a *ast) pushMathOp(c byte) {
	switch c {
	case '+', '-':
		a.rotateType(TypOpAdd, TypOpSub, TypOpMul, TypOpDiv)
	case '*', '/':
		a.rotateType(TypOpMul, TypOpDiv)
	}
	a.ops.push(Node{Typ: typMathOp[c]})
}

func (a *ast) pushInvoke() {
	n := a.rpn.pop().setTyp(TypInvoke)
	a.ops.push(n)
	a.caps.push(n)
}

func (a *ast) pushVector() {
	n := Node{Typ: TypVector}
	a.ops.push(n)
	a.caps.push(n)
}

func (a *ast) openPths() {
	a.ops.push(Node{Typ: typPths})
}

func (a *ast) closePths() {
	a.rotateBreakOn(typPths)
	a.ops.pop()
}

func (a *ast) rotateComma() {
	a.rotateBreakOn(a.incCaps().Typ)
}

func (a *ast) incCaps() Node {
	n := a.caps.pop().incCap()
	a.caps.push(n)
	return n
}

func (a *ast) popCaps() {
	n := a.caps.pop()
	a.rotateBreakOn(n.Typ)
	a.ops.pop()
	a.rpn.push(n)
}

func (a *ast) rotateType(t ...Typ) {
	for !a.ops.empty() {
		if !a.ops.top().typeOf(t...) {
			break
		}
		a.rpn.push(a.ops.pop())
	}
}

func (a *ast) rotateBreakOn(t ...Typ) {
	for !a.ops.empty() {
		if a.ops.top().typeOf(t...) {
			break
		}
		a.rpn.push(a.ops.pop())
	}
}
