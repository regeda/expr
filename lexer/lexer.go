//line lexer/lexer.go.rl:1
package lexer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

//line lexer/lexer.go:11
const lexer_start int = 1
const lexer_first_final int = 95
const lexer_error int = 0

const lexer_en_pths_rest int = 23
const lexer_en_vector_rest int = 37
const lexer_en_invoke_rest int = 66
const lexer_en_main int = 1

//line lexer/lexer.go.rl:10

type Lexer struct {
	p, pe, pb, eof int
	cs, top        int
	data           []byte
	stack          [1024]int
	rpn, ops, caps nvec
}

func (l *Lexer) text() string {
	return string(l.data[l.pb:l.p])
}

func (l *Lexer) pushInt(n int64) {
	l.rpn.push(Node{Typ: TypInt, DatN: n})
}

func (l *Lexer) pushStr(s string) {
	l.rpn.push(Node{Typ: TypStr, DatS: s})
}

func (l *Lexer) pushIdent() {
	l.rpn.push(Node{Typ: typIdent, DatS: strings.ToLower(l.text())})
}

func (l *Lexer) pushTrue() {
	l.rpn.push(Node{Typ: TypTrue})
}

func (l *Lexer) pushFalse() {
	l.rpn.push(Node{Typ: TypFalse})
}

func (l *Lexer) pushMathOp() {
	switch l.text() {
	case "+":
		l.rotateTypeOf(TypOpAdd, TypOpSub, TypOpMul, TypOpDiv)
		l.ops.push(Node{Typ: TypOpAdd})
	case "-":
		l.rotateTypeOf(TypOpAdd, TypOpSub, TypOpMul, TypOpDiv)
		l.ops.push(Node{Typ: TypOpSub})
	case "*":
		l.rotateTypeOf(TypOpMul, TypOpDiv)
		l.ops.push(Node{Typ: TypOpMul})
	case "/":
		l.rotateTypeOf(TypOpMul, TypOpDiv)
		l.ops.push(Node{Typ: TypOpDiv})
	}
}

func (l *Lexer) pushInvoke() {
	n := l.rpn.pop().setTyp(TypInvoke)
	l.ops.push(n)
	l.caps.push(n)
}

func (l *Lexer) pushVector() {
	n := Node{Typ: TypVector}
	l.ops.push(n)
	l.caps.push(n)
}

func (l *Lexer) openPths() {
	l.ops.push(Node{Typ: typPths})
}

func (l *Lexer) closePths() {
	l.rotateUntil(typPths)
}

func (l *Lexer) rotateTypeOf(t ...Typ) {
	for !l.ops.empty() {
		if !l.ops.top().typeOf(t...) {
			break
		}
		l.rpn.push(l.ops.pop())
	}
}

func (l *Lexer) rotateComma() {
	n := l.caps.pop().incCap()
	l.rotateUntil(n.Typ)
	l.ops.push(n)
	l.caps.push(n)
}

func (l *Lexer) popCaps() {
	n := l.caps.pop()
	l.rotateUntil(n.Typ)
	l.rpn.push(n)
}

func (l *Lexer) rotateUntil(t Typ) {
	for !l.ops.empty() {
		n := l.ops.pop()
		if n.Typ == t {
			break
		}
		l.rpn.push(n)
	}
}

func (l *Lexer) Parse(input []byte) ([]Node, error) {
	l.rpn.reset()
	l.ops.reset()
	l.caps.reset()

	l.data = input
	l.p = 0
	l.pb = 0
	l.pe = len(input)
	l.eof = len(input)

	var perr error
	var n64 int64
	var str string

//line lexer/lexer.go:142
	{
		l.cs = lexer_start
		l.top = 0
	}

//line lexer/lexer.go:148
	{
		if (l.p) == (l.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch l.cs {
		case 1:
			goto st1
		case 0:
			goto st0
		case 2:
			goto st2
		case 95:
			goto st95
		case 96:
			goto st96
		case 3:
			goto st3
		case 97:
			goto st97
		case 4:
			goto st4
		case 5:
			goto st5
		case 6:
			goto st6
		case 98:
			goto st98
		case 7:
			goto st7
		case 8:
			goto st8
		case 9:
			goto st9
		case 10:
			goto st10
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 14:
			goto st14
		case 15:
			goto st15
		case 16:
			goto st16
		case 17:
			goto st17
		case 18:
			goto st18
		case 19:
			goto st19
		case 99:
			goto st99
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 100:
			goto st100
		case 23:
			goto st23
		case 24:
			goto st24
		case 101:
			goto st101
		case 25:
			goto st25
		case 26:
			goto st26
		case 27:
			goto st27
		case 28:
			goto st28
		case 29:
			goto st29
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 35:
			goto st35
		case 36:
			goto st36
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 40:
			goto st40
		case 41:
			goto st41
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 102:
			goto st102
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 66:
			goto st66
		case 67:
			goto st67
		case 68:
			goto st68
		case 69:
			goto st69
		case 103:
			goto st103
		case 70:
			goto st70
		case 71:
			goto st71
		case 72:
			goto st72
		case 73:
			goto st73
		case 74:
			goto st74
		case 75:
			goto st75
		case 76:
			goto st76
		case 77:
			goto st77
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 82:
			goto st82
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 86:
			goto st86
		case 87:
			goto st87
		case 88:
			goto st88
		case 89:
			goto st89
		case 90:
			goto st90
		case 91:
			goto st91
		case 92:
			goto st92
		case 93:
			goto st93
		case 94:
			goto st94
		}

		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof
		}
	_resume:
		switch l.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 3:
			goto st_case_3
		case 97:
			goto st_case_97
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 98:
			goto st_case_98
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 99:
			goto st_case_99
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 100:
			goto st_case_100
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 101:
			goto st_case_101
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 102:
			goto st_case_102
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 103:
			goto st_case_103
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		}
		goto st_out
	st1:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch l.data[(l.p)] {
		case 32:
			goto st1
		case 34:
			goto tr2
		case 40:
			goto tr3
		case 43:
			goto tr4
		case 45:
			goto tr4
		case 91:
			goto tr7
		case 95:
			goto tr6
		case 102:
			goto tr8
		case 116:
			goto tr9
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st1
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr6
				}
			case l.data[(l.p)] >= 65:
				goto tr6
			}
		default:
			goto tr5
		}
		goto st0
	st_case_0:
	st0:
		l.cs = 0
		goto _out
	tr2:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st2
	st2:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof2
		}
	st_case_2:
//line lexer/lexer.go:638
		switch l.data[(l.p)] {
		case 34:
			goto st95
		case 92:
			goto st3
		}
		goto st2
	st95:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof95
		}
	st_case_95:
		if l.data[(l.p)] == 32 {
			goto tr206
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr206
		}
		goto st0
	tr7:
//line lexer/lexer.go.rl:141
		l.pushVector()
//line lexer/lexer.go.rl:192
		{
			l.stack[l.top] = 96
			l.top++
			goto st37
		}
		goto st96
	tr206:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 96
				goto _out
			}
		}
		l.pushStr(str)

		goto st96
	tr212:
//line lexer/lexer.go.rl:161
		l.pushFalse()
		goto st96
	tr213:
//line lexer/lexer.go.rl:160
		l.pushTrue()
		goto st96
	st96:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof96
		}
	st_case_96:
//line lexer/lexer.go:688
		if l.data[(l.p)] == 32 {
			goto st96
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st96
		}
		goto st0
	st3:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof3
		}
	st_case_3:
		goto st2
	tr3:
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 97
			l.top++
			goto st23
		}
		goto st97
	tr14:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 97
			l.top++
			goto st23
		}
		goto st97
	tr25:
//line lexer/lexer.go.rl:144
		l.pushIdent()
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 97
			l.top++
			goto st66
		}
		goto st97
	tr28:
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 97
			l.top++
			goto st66
		}
		goto st97
	tr210:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 97
				goto _out
			}
		}
		l.pushInt(n64)

		goto st97
	st97:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof97
		}
	st_case_97:
//line lexer/lexer.go:745
		switch l.data[(l.p)] {
		case 32:
			goto st97
		case 45:
			goto tr209
		case 47:
			goto tr209
		}
		switch {
		case l.data[(l.p)] > 13:
			if 42 <= l.data[(l.p)] && l.data[(l.p)] <= 43 {
				goto tr209
			}
		case l.data[(l.p)] >= 9:
			goto st97
		}
		goto st0
	tr209:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st4
	tr211:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 4
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st4
	st4:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof4
		}
	st_case_4:
//line lexer/lexer.go:784
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 43:
			goto tr15
		case 45:
			goto tr15
		case 95:
			goto tr17
		case 102:
			goto tr18
		case 116:
			goto tr19
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr13
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr17
				}
			case l.data[(l.p)] >= 65:
				goto tr17
			}
		default:
			goto tr16
		}
		goto st0
	tr13:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
		goto st5
	st5:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof5
		}
	st_case_5:
//line lexer/lexer.go:828
		switch l.data[(l.p)] {
		case 32:
			goto st5
		case 40:
			goto tr3
		case 43:
			goto tr4
		case 45:
			goto tr4
		case 95:
			goto tr6
		case 102:
			goto tr21
		case 116:
			goto tr22
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st5
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr6
				}
			case l.data[(l.p)] >= 65:
				goto tr6
			}
		default:
			goto tr5
		}
		goto st0
	tr4:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st6
	tr15:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st6
	st6:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof6
		}
	st_case_6:
//line lexer/lexer.go:878
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st98
		}
		goto st0
	tr5:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st98
	tr16:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st98
	st98:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof98
		}
	st_case_98:
//line lexer/lexer.go:898
		switch l.data[(l.p)] {
		case 32:
			goto tr210
		case 45:
			goto tr211
		case 47:
			goto tr211
		}
		switch {
		case l.data[(l.p)] < 42:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr210
			}
		case l.data[(l.p)] > 43:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st98
			}
		default:
			goto tr211
		}
		goto st0
	tr6:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st7
	tr17:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st7
	st7:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof7
		}
	st_case_7:
//line lexer/lexer.go:935
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	tr24:
//line lexer/lexer.go.rl:144
		l.pushIdent()
		goto st8
	st8:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof8
		}
	st_case_8:
//line lexer/lexer.go:971
		switch l.data[(l.p)] {
		case 32:
			goto st8
		case 40:
			goto tr28
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st8
		}
		goto st0
	tr21:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st9
	tr18:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st9
	st9:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof9
		}
	st_case_9:
//line lexer/lexer.go:997
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 97:
			goto st10
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st10:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 108:
			goto st11
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st11:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 115:
			goto st12
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st12:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 101:
			goto st13
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st13:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof13
		}
	st_case_13:
		if l.data[(l.p)] == 95 {
			goto st7
		}
		switch {
		case l.data[(l.p)] < 65:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st7
			}
		case l.data[(l.p)] > 90:
			if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	tr22:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st14
	tr19:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st14
	st14:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof14
		}
	st_case_14:
//line lexer/lexer.go:1161
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 114:
			goto st15
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st15:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof15
		}
	st_case_15:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 117:
			goto st12
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	tr8:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st16
	st16:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof16
		}
	st_case_16:
//line lexer/lexer.go:1232
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 97:
			goto st17
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st17:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 108:
			goto st18
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st18:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 115:
			goto st19
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st19:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 101:
			goto st99
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st99:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof99
		}
	st_case_99:
		switch l.data[(l.p)] {
		case 32:
			goto tr212
		case 95:
			goto st7
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr212
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	tr9:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st20
	st20:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof20
		}
	st_case_20:
//line lexer/lexer.go:1398
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 114:
			goto st21
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st21:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 117:
			goto st22
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st22:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch l.data[(l.p)] {
		case 32:
			goto tr24
		case 40:
			goto tr25
		case 95:
			goto st7
		case 101:
			goto st100
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr24
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	st100:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof100
		}
	st_case_100:
		switch l.data[(l.p)] {
		case 32:
			goto tr213
		case 95:
			goto st7
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr213
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st7
				}
			case l.data[(l.p)] >= 65:
				goto st7
			}
		default:
			goto st7
		}
		goto st0
	tr51:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
		goto st23
	st23:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof23
		}
	st_case_23:
//line lexer/lexer.go:1531
		switch l.data[(l.p)] {
		case 32:
			goto st23
		case 40:
			goto tr42
		case 43:
			goto tr43
		case 45:
			goto tr43
		case 95:
			goto tr45
		case 102:
			goto tr46
		case 116:
			goto tr47
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st23
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr45
				}
			case l.data[(l.p)] >= 65:
				goto tr45
			}
		default:
			goto tr44
		}
		goto st0
	tr42:
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 24
			l.top++
			goto st23
		}
		goto st24
	tr52:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 24
			l.top++
			goto st23
		}
		goto st24
	tr63:
//line lexer/lexer.go.rl:144
		l.pushIdent()
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 24
			l.top++
			goto st66
		}
		goto st24
	tr66:
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 24
			l.top++
			goto st66
		}
		goto st24
	tr59:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 24
				goto _out
			}
		}
		l.pushInt(n64)

		goto st24
	st24:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof24
		}
	st_case_24:
//line lexer/lexer.go:1609
		switch l.data[(l.p)] {
		case 32:
			goto st24
		case 41:
			goto tr49
		case 45:
			goto tr50
		case 47:
			goto tr50
		}
		switch {
		case l.data[(l.p)] > 13:
			if 42 <= l.data[(l.p)] && l.data[(l.p)] <= 43 {
				goto tr50
			}
		case l.data[(l.p)] >= 9:
			goto st24
		}
		goto st0
	tr49:
//line lexer/lexer.go.rl:143
		l.closePths()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st101
	tr60:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 101
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:143
		l.closePths()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st101
	st101:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof101
		}
	st_case_101:
//line lexer/lexer.go:1654
		goto st0
	tr50:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st25
	tr61:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 25
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st25
	st25:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof25
		}
	st_case_25:
//line lexer/lexer.go:1677
		switch l.data[(l.p)] {
		case 32:
			goto tr51
		case 40:
			goto tr52
		case 43:
			goto tr53
		case 45:
			goto tr53
		case 95:
			goto tr55
		case 102:
			goto tr56
		case 116:
			goto tr57
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr51
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr55
				}
			case l.data[(l.p)] >= 65:
				goto tr55
			}
		default:
			goto tr54
		}
		goto st0
	tr43:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st26
	tr53:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st26
	st26:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof26
		}
	st_case_26:
//line lexer/lexer.go:1727
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st27
		}
		goto st0
	tr44:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st27
	tr54:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st27
	st27:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof27
		}
	st_case_27:
//line lexer/lexer.go:1747
		switch l.data[(l.p)] {
		case 32:
			goto tr59
		case 41:
			goto tr60
		case 45:
			goto tr61
		case 47:
			goto tr61
		}
		switch {
		case l.data[(l.p)] < 42:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr59
			}
		case l.data[(l.p)] > 43:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st27
			}
		default:
			goto tr61
		}
		goto st0
	tr45:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st28
	tr55:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st28
	st28:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof28
		}
	st_case_28:
//line lexer/lexer.go:1786
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	tr62:
//line lexer/lexer.go.rl:144
		l.pushIdent()
		goto st29
	st29:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof29
		}
	st_case_29:
//line lexer/lexer.go:1822
		switch l.data[(l.p)] {
		case 32:
			goto st29
		case 40:
			goto tr66
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st29
		}
		goto st0
	tr46:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st30
	tr56:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st30
	st30:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof30
		}
	st_case_30:
//line lexer/lexer.go:1848
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		case 97:
			goto st31
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st31:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof31
		}
	st_case_31:
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		case 108:
			goto st32
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st32:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof32
		}
	st_case_32:
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		case 115:
			goto st33
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st33:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof33
		}
	st_case_33:
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		case 101:
			goto st34
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st34:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof34
		}
	st_case_34:
		if l.data[(l.p)] == 95 {
			goto st28
		}
		switch {
		case l.data[(l.p)] < 65:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st28
			}
		case l.data[(l.p)] > 90:
			if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	tr47:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st35
	tr57:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st35
	st35:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof35
		}
	st_case_35:
//line lexer/lexer.go:2012
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		case 114:
			goto st36
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st36:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof36
		}
	st_case_36:
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 40:
			goto tr63
		case 95:
			goto st28
		case 117:
			goto st33
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr62
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st28
				}
			case l.data[(l.p)] >= 65:
				goto st28
			}
		default:
			goto st28
		}
		goto st0
	st37:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof37
		}
	st_case_37:
		switch l.data[(l.p)] {
		case 32:
			goto st37
		case 34:
			goto tr73
		case 40:
			goto tr74
		case 43:
			goto tr75
		case 45:
			goto tr75
		case 91:
			goto tr78
		case 93:
			goto tr79
		case 95:
			goto tr77
		case 102:
			goto tr80
		case 116:
			goto tr81
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st37
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr77
				}
			case l.data[(l.p)] >= 65:
				goto tr77
			}
		default:
			goto tr76
		}
		goto st0
	tr91:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st38
	tr73:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st38
	st38:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof38
		}
	st_case_38:
//line lexer/lexer.go:2134
		switch l.data[(l.p)] {
		case 34:
			goto st39
		case 92:
			goto st65
		}
		goto st38
	st39:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch l.data[(l.p)] {
		case 32:
			goto tr85
		case 44:
			goto tr86
		case 93:
			goto tr87
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr85
		}
		goto st0
	tr96:
//line lexer/lexer.go.rl:141
		l.pushVector()
//line lexer/lexer.go.rl:192
		{
			l.stack[l.top] = 40
			l.top++
			goto st37
		}
		goto st40
	tr78:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:141
		l.pushVector()
//line lexer/lexer.go.rl:192
		{
			l.stack[l.top] = 40
			l.top++
			goto st37
		}
		goto st40
	tr85:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 40
				goto _out
			}
		}
		l.pushStr(str)

		goto st40
	tr130:
//line lexer/lexer.go.rl:161
		l.pushFalse()
		goto st40
	tr136:
//line lexer/lexer.go.rl:160
		l.pushTrue()
		goto st40
	st40:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof40
		}
	st_case_40:
//line lexer/lexer.go:2197
		switch l.data[(l.p)] {
		case 32:
			goto st40
		case 44:
			goto tr89
		case 93:
			goto tr79
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st40
		}
		goto st0
	tr86:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 41
				goto _out
			}
		}
		l.pushStr(str)

//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st41
	tr89:
//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st41
	tr114:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 41
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st41
	tr131:
//line lexer/lexer.go.rl:161
		l.pushFalse()
//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st41
	tr137:
//line lexer/lexer.go.rl:160
		l.pushTrue()
//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st41
	st41:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof41
		}
	st_case_41:
//line lexer/lexer.go:2256
		switch l.data[(l.p)] {
		case 32:
			goto st41
		case 34:
			goto tr91
		case 40:
			goto tr92
		case 43:
			goto tr93
		case 45:
			goto tr93
		case 91:
			goto tr96
		case 95:
			goto tr95
		case 102:
			goto tr97
		case 116:
			goto tr98
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st41
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr95
				}
			case l.data[(l.p)] >= 65:
				goto tr95
			}
		default:
			goto tr94
		}
		goto st0
	tr92:
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 42
			l.top++
			goto st23
		}
		goto st42
	tr102:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 42
			l.top++
			goto st23
		}
		goto st42
	tr117:
//line lexer/lexer.go.rl:144
		l.pushIdent()
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 42
			l.top++
			goto st66
		}
		goto st42
	tr120:
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 42
			l.top++
			goto st66
		}
		goto st42
	tr112:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 42
				goto _out
			}
		}
		l.pushInt(n64)

		goto st42
	tr74:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 42
			l.top++
			goto st23
		}
		goto st42
	st42:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof42
		}
	st_case_42:
//line lexer/lexer.go:2346
		switch l.data[(l.p)] {
		case 32:
			goto st42
		case 44:
			goto tr89
		case 47:
			goto tr100
		case 93:
			goto tr79
		}
		switch {
		case l.data[(l.p)] > 13:
			if 42 <= l.data[(l.p)] && l.data[(l.p)] <= 45 {
				goto tr100
			}
		case l.data[(l.p)] >= 9:
			goto st42
		}
		goto st0
	tr100:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st43
	tr113:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 43
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st43
	st43:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof43
		}
	st_case_43:
//line lexer/lexer.go:2387
		switch l.data[(l.p)] {
		case 32:
			goto tr101
		case 40:
			goto tr102
		case 43:
			goto tr103
		case 45:
			goto tr103
		case 95:
			goto tr105
		case 102:
			goto tr106
		case 116:
			goto tr107
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr101
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr105
				}
			case l.data[(l.p)] >= 65:
				goto tr105
			}
		default:
			goto tr104
		}
		goto st0
	tr101:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
		goto st44
	st44:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof44
		}
	st_case_44:
//line lexer/lexer.go:2431
		switch l.data[(l.p)] {
		case 32:
			goto st44
		case 40:
			goto tr92
		case 43:
			goto tr93
		case 45:
			goto tr93
		case 95:
			goto tr95
		case 102:
			goto tr109
		case 116:
			goto tr110
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st44
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr95
				}
			case l.data[(l.p)] >= 65:
				goto tr95
			}
		default:
			goto tr94
		}
		goto st0
	tr93:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st45
	tr103:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st45
	tr75:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st45
	st45:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof45
		}
	st_case_45:
//line lexer/lexer.go:2487
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st46
		}
		goto st0
	tr94:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st46
	tr104:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st46
	tr76:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st46
	st46:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof46
		}
	st_case_46:
//line lexer/lexer.go:2513
		switch l.data[(l.p)] {
		case 32:
			goto tr112
		case 44:
			goto tr114
		case 47:
			goto tr113
		case 93:
			goto tr115
		}
		switch {
		case l.data[(l.p)] < 42:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr112
			}
		case l.data[(l.p)] > 45:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st46
			}
		default:
			goto tr113
		}
		goto st0
	tr79:
//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st102
	tr87:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 102
				goto _out
			}
		}
		l.pushStr(str)

//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st102
	tr115:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 102
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st102
	tr132:
//line lexer/lexer.go.rl:161
		l.pushFalse()
//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st102
	tr138:
//line lexer/lexer.go.rl:160
		l.pushTrue()
//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st102
	st102:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof102
		}
	st_case_102:
//line lexer/lexer.go:2593
		goto st0
	tr95:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st47
	tr105:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st47
	tr77:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st47
	st47:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof47
		}
	st_case_47:
//line lexer/lexer.go:2616
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	tr116:
//line lexer/lexer.go.rl:144
		l.pushIdent()
		goto st48
	st48:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof48
		}
	st_case_48:
//line lexer/lexer.go:2652
		switch l.data[(l.p)] {
		case 32:
			goto st48
		case 40:
			goto tr120
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st48
		}
		goto st0
	tr109:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st49
	tr106:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st49
	st49:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof49
		}
	st_case_49:
//line lexer/lexer.go:2678
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 97:
			goto st50
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st50:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 108:
			goto st51
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st51:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 115:
			goto st52
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st52:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 101:
			goto st53
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st53:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof53
		}
	st_case_53:
		if l.data[(l.p)] == 95 {
			goto st47
		}
		switch {
		case l.data[(l.p)] < 65:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st47
			}
		case l.data[(l.p)] > 90:
			if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	tr110:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st54
	tr107:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st54
	st54:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof54
		}
	st_case_54:
//line lexer/lexer.go:2842
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 114:
			goto st55
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st55:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 117:
			goto st52
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	tr97:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st56
	tr80:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st56
	st56:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof56
		}
	st_case_56:
//line lexer/lexer.go:2919
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 97:
			goto st57
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st57:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 108:
			goto st58
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st58:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 115:
			goto st59
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st59:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 101:
			goto st60
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st60:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch l.data[(l.p)] {
		case 32:
			goto tr130
		case 44:
			goto tr131
		case 93:
			goto tr132
		case 95:
			goto st47
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr130
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	tr98:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st61
	tr81:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st61
	st61:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof61
		}
	st_case_61:
//line lexer/lexer.go:3095
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 114:
			goto st62
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st62:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 117:
			goto st63
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st63:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch l.data[(l.p)] {
		case 32:
			goto tr116
		case 40:
			goto tr117
		case 95:
			goto st47
		case 101:
			goto st64
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr116
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st64:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch l.data[(l.p)] {
		case 32:
			goto tr136
		case 44:
			goto tr137
		case 93:
			goto tr138
		case 95:
			goto st47
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr136
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st47
				}
			case l.data[(l.p)] >= 65:
				goto st47
			}
		default:
			goto st47
		}
		goto st0
	st65:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof65
		}
	st_case_65:
		goto st38
	st66:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch l.data[(l.p)] {
		case 32:
			goto st66
		case 34:
			goto tr140
		case 40:
			goto tr141
		case 41:
			goto tr142
		case 43:
			goto tr143
		case 45:
			goto tr143
		case 91:
			goto tr146
		case 95:
			goto tr145
		case 102:
			goto tr147
		case 116:
			goto tr148
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st66
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr145
				}
			case l.data[(l.p)] >= 65:
				goto tr145
			}
		default:
			goto tr144
		}
		goto st0
	tr158:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st67
	tr140:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st67
	st67:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof67
		}
	st_case_67:
//line lexer/lexer.go:3289
		switch l.data[(l.p)] {
		case 34:
			goto st68
		case 92:
			goto st94
		}
		goto st67
	st68:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch l.data[(l.p)] {
		case 32:
			goto tr152
		case 41:
			goto tr153
		case 44:
			goto tr154
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr152
		}
		goto st0
	tr163:
//line lexer/lexer.go.rl:141
		l.pushVector()
//line lexer/lexer.go.rl:192
		{
			l.stack[l.top] = 69
			l.top++
			goto st37
		}
		goto st69
	tr146:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:141
		l.pushVector()
//line lexer/lexer.go.rl:192
		{
			l.stack[l.top] = 69
			l.top++
			goto st37
		}
		goto st69
	tr152:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 69
				goto _out
			}
		}
		l.pushStr(str)

		goto st69
	tr197:
//line lexer/lexer.go.rl:161
		l.pushFalse()
		goto st69
	tr203:
//line lexer/lexer.go.rl:160
		l.pushTrue()
		goto st69
	st69:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof69
		}
	st_case_69:
//line lexer/lexer.go:3352
		switch l.data[(l.p)] {
		case 32:
			goto st69
		case 41:
			goto tr142
		case 44:
			goto tr156
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st69
		}
		goto st0
	tr142:
//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st103
	tr153:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 103
				goto _out
			}
		}
		l.pushStr(str)

//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st103
	tr180:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 103
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st103
	tr198:
//line lexer/lexer.go.rl:161
		l.pushFalse()
//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st103
	tr204:
//line lexer/lexer.go.rl:160
		l.pushTrue()
//line lexer/lexer.go.rl:139
		l.popCaps()
//line lexer/lexer.go.rl:135
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st103
	st103:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof103
		}
	st_case_103:
//line lexer/lexer.go:3421
		goto st0
	tr154:
//line lexer/lexer.go.rl:145

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 70
				goto _out
			}
		}
		l.pushStr(str)

//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st70
	tr156:
//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st70
	tr182:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 70
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st70
	tr199:
//line lexer/lexer.go.rl:161
		l.pushFalse()
//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st70
	tr205:
//line lexer/lexer.go.rl:160
		l.pushTrue()
//line lexer/lexer.go.rl:138
		l.rotateComma()
		goto st70
	st70:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof70
		}
	st_case_70:
//line lexer/lexer.go:3469
		switch l.data[(l.p)] {
		case 32:
			goto st70
		case 34:
			goto tr158
		case 40:
			goto tr159
		case 43:
			goto tr160
		case 45:
			goto tr160
		case 91:
			goto tr163
		case 95:
			goto tr162
		case 102:
			goto tr164
		case 116:
			goto tr165
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st70
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr162
				}
			case l.data[(l.p)] >= 65:
				goto tr162
			}
		default:
			goto tr161
		}
		goto st0
	tr159:
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 71
			l.top++
			goto st23
		}
		goto st71
	tr169:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 71
			l.top++
			goto st23
		}
		goto st71
	tr184:
//line lexer/lexer.go.rl:144
		l.pushIdent()
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 71
			l.top++
			goto st66
		}
		goto st71
	tr187:
//line lexer/lexer.go.rl:140
		l.pushInvoke()
//line lexer/lexer.go.rl:193
		{
			l.stack[l.top] = 71
			l.top++
			goto st66
		}
		goto st71
	tr179:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 71
				goto _out
			}
		}
		l.pushInt(n64)

		goto st71
	tr141:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:142
		l.openPths()
//line lexer/lexer.go.rl:194
		{
			l.stack[l.top] = 71
			l.top++
			goto st23
		}
		goto st71
	st71:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof71
		}
	st_case_71:
//line lexer/lexer.go:3559
		switch l.data[(l.p)] {
		case 32:
			goto st71
		case 41:
			goto tr142
		case 44:
			goto tr156
		case 47:
			goto tr167
		}
		switch {
		case l.data[(l.p)] > 13:
			if 42 <= l.data[(l.p)] && l.data[(l.p)] <= 45 {
				goto tr167
			}
		case l.data[(l.p)] >= 9:
			goto st71
		}
		goto st0
	tr167:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st72
	tr181:
//line lexer/lexer.go.rl:153

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 72
				goto _out
			}
		}
		l.pushInt(n64)

//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st72
	st72:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof72
		}
	st_case_72:
//line lexer/lexer.go:3600
		switch l.data[(l.p)] {
		case 32:
			goto tr168
		case 40:
			goto tr169
		case 43:
			goto tr170
		case 45:
			goto tr170
		case 95:
			goto tr172
		case 102:
			goto tr173
		case 116:
			goto tr174
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr168
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr172
				}
			case l.data[(l.p)] >= 65:
				goto tr172
			}
		default:
			goto tr171
		}
		goto st0
	tr168:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
		goto st73
	st73:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof73
		}
	st_case_73:
//line lexer/lexer.go:3644
		switch l.data[(l.p)] {
		case 32:
			goto st73
		case 40:
			goto tr159
		case 43:
			goto tr160
		case 45:
			goto tr160
		case 95:
			goto tr162
		case 102:
			goto tr176
		case 116:
			goto tr177
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st73
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto tr162
				}
			case l.data[(l.p)] >= 65:
				goto tr162
			}
		default:
			goto tr161
		}
		goto st0
	tr160:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st74
	tr170:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st74
	tr143:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st74
	st74:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof74
		}
	st_case_74:
//line lexer/lexer.go:3700
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st75
		}
		goto st0
	tr161:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st75
	tr171:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st75
	tr144:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st75
	st75:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof75
		}
	st_case_75:
//line lexer/lexer.go:3726
		switch l.data[(l.p)] {
		case 32:
			goto tr179
		case 41:
			goto tr180
		case 44:
			goto tr182
		case 47:
			goto tr181
		}
		switch {
		case l.data[(l.p)] < 42:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr179
			}
		case l.data[(l.p)] > 45:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st75
			}
		default:
			goto tr181
		}
		goto st0
	tr162:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st76
	tr172:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st76
	tr145:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st76
	st76:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof76
		}
	st_case_76:
//line lexer/lexer.go:3771
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	tr183:
//line lexer/lexer.go.rl:144
		l.pushIdent()
		goto st77
	st77:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof77
		}
	st_case_77:
//line lexer/lexer.go:3807
		switch l.data[(l.p)] {
		case 32:
			goto st77
		case 40:
			goto tr187
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st77
		}
		goto st0
	tr176:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st78
	tr173:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st78
	st78:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof78
		}
	st_case_78:
//line lexer/lexer.go:3833
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 97:
			goto st79
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st79:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 108:
			goto st80
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st80:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 115:
			goto st81
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st81:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 101:
			goto st82
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st82:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof82
		}
	st_case_82:
		if l.data[(l.p)] == 95 {
			goto st76
		}
		switch {
		case l.data[(l.p)] < 65:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st76
			}
		case l.data[(l.p)] > 90:
			if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	tr177:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st83
	tr174:
//line lexer/lexer.go.rl:162
		l.pushMathOp()
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st83
	st83:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof83
		}
	st_case_83:
//line lexer/lexer.go:3997
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 114:
			goto st84
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st84:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 117:
			goto st81
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	tr164:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st85
	tr147:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st85
	st85:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof85
		}
	st_case_85:
//line lexer/lexer.go:4074
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 97:
			goto st86
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st86:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof86
		}
	st_case_86:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 108:
			goto st87
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st87:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 115:
			goto st88
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st88:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 101:
			goto st89
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st89:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch l.data[(l.p)] {
		case 32:
			goto tr197
		case 41:
			goto tr198
		case 44:
			goto tr199
		case 95:
			goto st76
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr197
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	tr165:
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st90
	tr148:
//line lexer/lexer.go.rl:137
		l.caps.push(l.caps.pop().incCap())
//line lexer/lexer.go.rl:136
		l.pb = l.p
		goto st90
	st90:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof90
		}
	st_case_90:
//line lexer/lexer.go:4250
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 114:
			goto st91
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st91:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 117:
			goto st92
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st92:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch l.data[(l.p)] {
		case 32:
			goto tr183
		case 40:
			goto tr184
		case 95:
			goto st76
		case 101:
			goto st93
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr183
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st93:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch l.data[(l.p)] {
		case 32:
			goto tr203
		case 41:
			goto tr204
		case 44:
			goto tr205
		case 95:
			goto st76
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr203
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st76
				}
			case l.data[(l.p)] >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
	st94:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof94
		}
	st_case_94:
		goto st67
	st_out:
	_test_eof1:
		l.cs = 1
		goto _test_eof
	_test_eof2:
		l.cs = 2
		goto _test_eof
	_test_eof95:
		l.cs = 95
		goto _test_eof
	_test_eof96:
		l.cs = 96
		goto _test_eof
	_test_eof3:
		l.cs = 3
		goto _test_eof
	_test_eof97:
		l.cs = 97
		goto _test_eof
	_test_eof4:
		l.cs = 4
		goto _test_eof
	_test_eof5:
		l.cs = 5
		goto _test_eof
	_test_eof6:
		l.cs = 6
		goto _test_eof
	_test_eof98:
		l.cs = 98
		goto _test_eof
	_test_eof7:
		l.cs = 7
		goto _test_eof
	_test_eof8:
		l.cs = 8
		goto _test_eof
	_test_eof9:
		l.cs = 9
		goto _test_eof
	_test_eof10:
		l.cs = 10
		goto _test_eof
	_test_eof11:
		l.cs = 11
		goto _test_eof
	_test_eof12:
		l.cs = 12
		goto _test_eof
	_test_eof13:
		l.cs = 13
		goto _test_eof
	_test_eof14:
		l.cs = 14
		goto _test_eof
	_test_eof15:
		l.cs = 15
		goto _test_eof
	_test_eof16:
		l.cs = 16
		goto _test_eof
	_test_eof17:
		l.cs = 17
		goto _test_eof
	_test_eof18:
		l.cs = 18
		goto _test_eof
	_test_eof19:
		l.cs = 19
		goto _test_eof
	_test_eof99:
		l.cs = 99
		goto _test_eof
	_test_eof20:
		l.cs = 20
		goto _test_eof
	_test_eof21:
		l.cs = 21
		goto _test_eof
	_test_eof22:
		l.cs = 22
		goto _test_eof
	_test_eof100:
		l.cs = 100
		goto _test_eof
	_test_eof23:
		l.cs = 23
		goto _test_eof
	_test_eof24:
		l.cs = 24
		goto _test_eof
	_test_eof101:
		l.cs = 101
		goto _test_eof
	_test_eof25:
		l.cs = 25
		goto _test_eof
	_test_eof26:
		l.cs = 26
		goto _test_eof
	_test_eof27:
		l.cs = 27
		goto _test_eof
	_test_eof28:
		l.cs = 28
		goto _test_eof
	_test_eof29:
		l.cs = 29
		goto _test_eof
	_test_eof30:
		l.cs = 30
		goto _test_eof
	_test_eof31:
		l.cs = 31
		goto _test_eof
	_test_eof32:
		l.cs = 32
		goto _test_eof
	_test_eof33:
		l.cs = 33
		goto _test_eof
	_test_eof34:
		l.cs = 34
		goto _test_eof
	_test_eof35:
		l.cs = 35
		goto _test_eof
	_test_eof36:
		l.cs = 36
		goto _test_eof
	_test_eof37:
		l.cs = 37
		goto _test_eof
	_test_eof38:
		l.cs = 38
		goto _test_eof
	_test_eof39:
		l.cs = 39
		goto _test_eof
	_test_eof40:
		l.cs = 40
		goto _test_eof
	_test_eof41:
		l.cs = 41
		goto _test_eof
	_test_eof42:
		l.cs = 42
		goto _test_eof
	_test_eof43:
		l.cs = 43
		goto _test_eof
	_test_eof44:
		l.cs = 44
		goto _test_eof
	_test_eof45:
		l.cs = 45
		goto _test_eof
	_test_eof46:
		l.cs = 46
		goto _test_eof
	_test_eof102:
		l.cs = 102
		goto _test_eof
	_test_eof47:
		l.cs = 47
		goto _test_eof
	_test_eof48:
		l.cs = 48
		goto _test_eof
	_test_eof49:
		l.cs = 49
		goto _test_eof
	_test_eof50:
		l.cs = 50
		goto _test_eof
	_test_eof51:
		l.cs = 51
		goto _test_eof
	_test_eof52:
		l.cs = 52
		goto _test_eof
	_test_eof53:
		l.cs = 53
		goto _test_eof
	_test_eof54:
		l.cs = 54
		goto _test_eof
	_test_eof55:
		l.cs = 55
		goto _test_eof
	_test_eof56:
		l.cs = 56
		goto _test_eof
	_test_eof57:
		l.cs = 57
		goto _test_eof
	_test_eof58:
		l.cs = 58
		goto _test_eof
	_test_eof59:
		l.cs = 59
		goto _test_eof
	_test_eof60:
		l.cs = 60
		goto _test_eof
	_test_eof61:
		l.cs = 61
		goto _test_eof
	_test_eof62:
		l.cs = 62
		goto _test_eof
	_test_eof63:
		l.cs = 63
		goto _test_eof
	_test_eof64:
		l.cs = 64
		goto _test_eof
	_test_eof65:
		l.cs = 65
		goto _test_eof
	_test_eof66:
		l.cs = 66
		goto _test_eof
	_test_eof67:
		l.cs = 67
		goto _test_eof
	_test_eof68:
		l.cs = 68
		goto _test_eof
	_test_eof69:
		l.cs = 69
		goto _test_eof
	_test_eof103:
		l.cs = 103
		goto _test_eof
	_test_eof70:
		l.cs = 70
		goto _test_eof
	_test_eof71:
		l.cs = 71
		goto _test_eof
	_test_eof72:
		l.cs = 72
		goto _test_eof
	_test_eof73:
		l.cs = 73
		goto _test_eof
	_test_eof74:
		l.cs = 74
		goto _test_eof
	_test_eof75:
		l.cs = 75
		goto _test_eof
	_test_eof76:
		l.cs = 76
		goto _test_eof
	_test_eof77:
		l.cs = 77
		goto _test_eof
	_test_eof78:
		l.cs = 78
		goto _test_eof
	_test_eof79:
		l.cs = 79
		goto _test_eof
	_test_eof80:
		l.cs = 80
		goto _test_eof
	_test_eof81:
		l.cs = 81
		goto _test_eof
	_test_eof82:
		l.cs = 82
		goto _test_eof
	_test_eof83:
		l.cs = 83
		goto _test_eof
	_test_eof84:
		l.cs = 84
		goto _test_eof
	_test_eof85:
		l.cs = 85
		goto _test_eof
	_test_eof86:
		l.cs = 86
		goto _test_eof
	_test_eof87:
		l.cs = 87
		goto _test_eof
	_test_eof88:
		l.cs = 88
		goto _test_eof
	_test_eof89:
		l.cs = 89
		goto _test_eof
	_test_eof90:
		l.cs = 90
		goto _test_eof
	_test_eof91:
		l.cs = 91
		goto _test_eof
	_test_eof92:
		l.cs = 92
		goto _test_eof
	_test_eof93:
		l.cs = 93
		goto _test_eof
	_test_eof94:
		l.cs = 94
		goto _test_eof

	_test_eof:
		{
		}
		if (l.p) == (l.eof) {
			switch l.cs {
			case 95:
//line lexer/lexer.go.rl:145

				str, perr = strconv.Unquote(l.text())
				if perr != nil {
					perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
					{
						(l.p)++
						l.cs = 0
						goto _out
					}
				}
				l.pushStr(str)

			case 98:
//line lexer/lexer.go.rl:153

				n64, perr = strconv.ParseInt(l.text(), 10, 64)
				if perr != nil {
					{
						(l.p)++
						l.cs = 0
						goto _out
					}
				}
				l.pushInt(n64)

			case 100:
//line lexer/lexer.go.rl:160
				l.pushTrue()
			case 99:
//line lexer/lexer.go.rl:161
				l.pushFalse()
//line lexer/lexer.go:4517
			}
		}

	_out:
		{
		}
	}

//line lexer/lexer.go.rl:212

	if l.top > 0 {
		return nil, fmt.Errorf("stack parsing error at %d", l.pb)
	}

	if l.cs < 95 {
		if perr != nil {
			return nil, fmt.Errorf("parsing error at %d: %w", l.pb, perr)
		}
		return nil, fmt.Errorf("token parsing error at %d", l.pb)
	}

	for !l.ops.empty() {
		l.rpn.push(l.ops.pop())
	}

	return l.rpn, nil
}
