//line internal/lexer/lexer.go.rl:1
package lexer

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/regeda/expr/internal/ast"
	"github.com/regeda/expr/internal/ast/stack"
	"github.com/regeda/expr/internal/ast/value"
)

//line internal/lexer/lexer.go:17
const lexer_start int = 1
const lexer_first_final int = 60
const lexer_error int = 0

const lexer_en_invoke_rest int = 28
const lexer_en_main int = 1

//line internal/lexer/lexer.go.rl:16

type Lexer struct {
	data       []byte
	cs         int
	p, pe, eof int
	pb         int
	top        int
	stack      []int
	nodes      stack.Stack
}

func New() *Lexer {
	return &Lexer{
		stack: make([]int, 1024),
	}
}

func (l *Lexer) text() string {
	return string(l.data[l.pb:l.p])
}

func (l *Lexer) Parse(input []byte) (*ast.Node, error) {
	l.data = input
	l.p = 0
	l.pb = 0
	l.pe = len(input)
	l.eof = len(input)
	l.nodes.Reset()

	l.nodes.Push(value.Exit())

	var perr error
	var n64 int64
	var str string

//line internal/lexer/lexer.go:64
	{
		l.cs = lexer_start
		l.top = 0
	}

//line internal/lexer/lexer.go:70
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
		case 60:
			goto st60
		case 3:
			goto st3
		case 4:
			goto st4
		case 61:
			goto st61
		case 5:
			goto st5
		case 6:
			goto st6
		case 62:
			goto st62
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
		case 20:
			goto st20
		case 21:
			goto st21
		case 22:
			goto st22
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 63:
			goto st63
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
		case 64:
			goto st64
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
		case 60:
			goto st_case_60
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 61:
			goto st_case_61
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 62:
			goto st_case_62
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
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 63:
			goto st_case_63
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
		case 64:
			goto st_case_64
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
		case 43:
			goto tr3
		case 45:
			goto tr3
		case 91:
			goto tr6
		case 95:
			goto tr5
		case 102:
			goto tr7
		case 116:
			goto tr8
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
					goto tr5
				}
			case l.data[(l.p)] >= 65:
				goto tr5
			}
		default:
			goto tr4
		}
		goto st0
	st_case_0:
	st0:
		l.cs = 0
		goto _out
	tr2:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st2
	st2:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof2
		}
	st_case_2:
//line internal/lexer/lexer.go:402
		switch l.data[(l.p)] {
		case 34:
			goto st60
		case 92:
			goto st3
		}
		goto st2
	st60:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof60
		}
	st_case_60:
		goto st0
	st3:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof3
		}
	st_case_3:
		goto st2
	tr3:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st4
	st4:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof4
		}
	st_case_4:
//line internal/lexer/lexer.go:431
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st61
		}
		goto st0
	tr4:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st61
	st61:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof61
		}
	st_case_61:
//line internal/lexer/lexer.go:445
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st61
		}
		goto st0
	tr5:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st5
	st5:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof5
		}
	st_case_5:
//line internal/lexer/lexer.go:459
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
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
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	tr13:
//line internal/lexer/lexer.go.rl:63
		l.nodes.Push(l.nodes.Nest(value.Call(l.text())))
		goto st6
	st6:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof6
		}
	st_case_6:
//line internal/lexer/lexer.go:495
		switch l.data[(l.p)] {
		case 32:
			goto st6
		case 40:
			goto tr17
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st6
		}
		goto st0
	tr14:
//line internal/lexer/lexer.go.rl:63
		l.nodes.Push(l.nodes.Nest(value.Call(l.text())))
//line internal/lexer/lexer.go.rl:108
		{
			l.stack[l.top] = 62
			l.top++
			goto st28
		}
		goto st62
	tr17:
//line internal/lexer/lexer.go.rl:108
		{
			l.stack[l.top] = 62
			l.top++
			goto st28
		}
		goto st62
	tr22:
//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st62
	tr30:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 62
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st62
	tr36:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 62
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st62
	tr43:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st62
	st62:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof62
		}
	st_case_62:
//line internal/lexer/lexer.go:556
		goto st0
	tr6:
//line internal/lexer/lexer.go.rl:64
		l.nodes.Push(l.nodes.Nest(value.Arr()))
		goto st7
	st7:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof7
		}
	st_case_7:
//line internal/lexer/lexer.go:567
		switch l.data[(l.p)] {
		case 32:
			goto st7
		case 34:
			goto tr19
		case 43:
			goto tr20
		case 45:
			goto tr20
		case 93:
			goto tr22
		case 102:
			goto tr23
		case 116:
			goto tr24
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto tr21
			}
		case l.data[(l.p)] >= 9:
			goto st7
		}
		goto st0
	tr19:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st8
	st8:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof8
		}
	st_case_8:
//line internal/lexer/lexer.go:602
		switch l.data[(l.p)] {
		case 34:
			goto st9
		case 92:
			goto st21
		}
		goto st8
	st9:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch l.data[(l.p)] {
		case 32:
			goto tr28
		case 44:
			goto tr29
		case 93:
			goto tr30
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr28
		}
		goto st0
	tr28:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 10
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

		goto st10
	tr34:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 10
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

		goto st10
	tr41:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
		goto st10
	st10:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof10
		}
	st_case_10:
//line internal/lexer/lexer.go:657
		switch l.data[(l.p)] {
		case 32:
			goto st10
		case 44:
			goto st11
		case 93:
			goto tr22
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st10
		}
		goto st0
	tr29:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 11
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

		goto st11
	tr35:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 11
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

		goto st11
	tr42:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
		goto st11
	st11:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof11
		}
	st_case_11:
//line internal/lexer/lexer.go:700
		switch l.data[(l.p)] {
		case 32:
			goto st11
		case 34:
			goto tr19
		case 43:
			goto tr20
		case 45:
			goto tr20
		case 102:
			goto tr23
		case 116:
			goto tr24
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto tr21
			}
		case l.data[(l.p)] >= 9:
			goto st11
		}
		goto st0
	tr20:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st12
	st12:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof12
		}
	st_case_12:
//line internal/lexer/lexer.go:733
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st13
		}
		goto st0
	tr21:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st13
	st13:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof13
		}
	st_case_13:
//line internal/lexer/lexer.go:747
		switch l.data[(l.p)] {
		case 32:
			goto tr34
		case 44:
			goto tr35
		case 93:
			goto tr36
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st13
			}
		case l.data[(l.p)] >= 9:
			goto tr34
		}
		goto st0
	tr23:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st14
	st14:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof14
		}
	st_case_14:
//line internal/lexer/lexer.go:774
		if l.data[(l.p)] == 97 {
			goto st15
		}
		goto st0
	st15:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof15
		}
	st_case_15:
		if l.data[(l.p)] == 108 {
			goto st16
		}
		goto st0
	st16:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof16
		}
	st_case_16:
		if l.data[(l.p)] == 115 {
			goto st17
		}
		goto st0
	st17:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof17
		}
	st_case_17:
		if l.data[(l.p)] == 101 {
			goto st18
		}
		goto st0
	st18:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch l.data[(l.p)] {
		case 32:
			goto tr41
		case 44:
			goto tr42
		case 93:
			goto tr43
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr41
		}
		goto st0
	tr24:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st19
	st19:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof19
		}
	st_case_19:
//line internal/lexer/lexer.go:832
		if l.data[(l.p)] == 114 {
			goto st20
		}
		goto st0
	st20:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof20
		}
	st_case_20:
		if l.data[(l.p)] == 117 {
			goto st17
		}
		goto st0
	st21:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof21
		}
	st_case_21:
		goto st8
	tr7:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st22
	st22:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof22
		}
	st_case_22:
//line internal/lexer/lexer.go:861
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
		case 97:
			goto st23
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr13
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st23:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
		case 108:
			goto st24
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
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st24:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
		case 115:
			goto st25
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
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st25:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
		case 101:
			goto st63
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
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st63:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof63
		}
	st_case_63:
		if l.data[(l.p)] == 95 {
			goto st5
		}
		switch {
		case l.data[(l.p)] < 65:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st5
			}
		case l.data[(l.p)] > 90:
			if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	tr8:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st26
	st26:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof26
		}
	st_case_26:
//line internal/lexer/lexer.go:1019
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
		case 114:
			goto st27
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
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st27:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch l.data[(l.p)] {
		case 32:
			goto tr13
		case 40:
			goto tr14
		case 95:
			goto st5
		case 117:
			goto st25
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
					goto st5
				}
			case l.data[(l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st28:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch l.data[(l.p)] {
		case 32:
			goto st28
		case 34:
			goto tr51
		case 41:
			goto tr52
		case 43:
			goto tr53
		case 45:
			goto tr53
		case 91:
			goto tr56
		case 95:
			goto tr55
		case 102:
			goto tr57
		case 116:
			goto tr58
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st28
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
	tr51:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st29
	st29:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof29
		}
	st_case_29:
//line internal/lexer/lexer.go:1133
		switch l.data[(l.p)] {
		case 34:
			goto st30
		case 92:
			goto st59
		}
		goto st29
	st30:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch l.data[(l.p)] {
		case 32:
			goto tr62
		case 41:
			goto tr63
		case 44:
			goto tr64
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr62
		}
		goto st0
	tr72:
//line internal/lexer/lexer.go.rl:63
		l.nodes.Push(l.nodes.Nest(value.Call(l.text())))
//line internal/lexer/lexer.go.rl:108
		{
			l.stack[l.top] = 31
			l.top++
			goto st28
		}
		goto st31
	tr75:
//line internal/lexer/lexer.go.rl:108
		{
			l.stack[l.top] = 31
			l.top++
			goto st28
		}
		goto st31
	tr80:
//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st31
	tr62:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 31
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

		goto st31
	tr88:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 31
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st31
	tr68:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 31
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

		goto st31
	tr94:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 31
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st31
	tr107:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
		goto st31
	tr101:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
		goto st31
	st31:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof31
		}
	st_case_31:
//line internal/lexer/lexer.go:1233
		switch l.data[(l.p)] {
		case 32:
			goto st31
		case 41:
			goto tr52
		case 44:
			goto st32
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st31
		}
		goto st0
	tr52:
//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
//line internal/lexer/lexer.go.rl:110
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st64
	tr63:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 64
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
//line internal/lexer/lexer.go.rl:110
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st64
	tr69:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 64
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
//line internal/lexer/lexer.go.rl:110
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st64
	tr108:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
//line internal/lexer/lexer.go.rl:61
		l.nodes.Pop()
//line internal/lexer/lexer.go.rl:110
		{
			l.top--
			l.cs = l.stack[l.top]
			goto _again
		}
		goto st64
	st64:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof64
		}
	st_case_64:
//line internal/lexer/lexer.go:1294
		goto st0
	tr64:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 32
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

		goto st32
	tr70:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 32
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

		goto st32
	tr109:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
		goto st32
	st32:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof32
		}
	st_case_32:
//line internal/lexer/lexer.go:1326
		switch l.data[(l.p)] {
		case 32:
			goto st32
		case 34:
			goto tr51
		case 43:
			goto tr53
		case 45:
			goto tr53
		case 91:
			goto tr56
		case 95:
			goto tr55
		case 102:
			goto tr57
		case 116:
			goto tr58
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto st32
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
	tr53:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st33
	st33:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof33
		}
	st_case_33:
//line internal/lexer/lexer.go:1372
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st34
		}
		goto st0
	tr54:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st34
	st34:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof34
		}
	st_case_34:
//line internal/lexer/lexer.go:1386
		switch l.data[(l.p)] {
		case 32:
			goto tr68
		case 41:
			goto tr69
		case 44:
			goto tr70
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st34
			}
		case l.data[(l.p)] >= 9:
			goto tr68
		}
		goto st0
	tr55:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st35
	st35:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof35
		}
	st_case_35:
//line internal/lexer/lexer.go:1413
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	tr71:
//line internal/lexer/lexer.go.rl:63
		l.nodes.Push(l.nodes.Nest(value.Call(l.text())))
		goto st36
	st36:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof36
		}
	st_case_36:
//line internal/lexer/lexer.go:1449
		switch l.data[(l.p)] {
		case 32:
			goto st36
		case 40:
			goto tr75
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st36
		}
		goto st0
	tr56:
//line internal/lexer/lexer.go.rl:64
		l.nodes.Push(l.nodes.Nest(value.Arr()))
		goto st37
	st37:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof37
		}
	st_case_37:
//line internal/lexer/lexer.go:1469
		switch l.data[(l.p)] {
		case 32:
			goto st37
		case 34:
			goto tr77
		case 43:
			goto tr78
		case 45:
			goto tr78
		case 93:
			goto tr80
		case 102:
			goto tr81
		case 116:
			goto tr82
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto tr79
			}
		case l.data[(l.p)] >= 9:
			goto st37
		}
		goto st0
	tr77:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st38
	st38:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof38
		}
	st_case_38:
//line internal/lexer/lexer.go:1504
		switch l.data[(l.p)] {
		case 34:
			goto st39
		case 92:
			goto st51
		}
		goto st38
	st39:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch l.data[(l.p)] {
		case 32:
			goto tr86
		case 44:
			goto tr87
		case 93:
			goto tr88
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr86
		}
		goto st0
	tr86:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 40
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

		goto st40
	tr92:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 40
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

		goto st40
	tr99:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
		goto st40
	st40:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof40
		}
	st_case_40:
//line internal/lexer/lexer.go:1559
		switch l.data[(l.p)] {
		case 32:
			goto st40
		case 44:
			goto st41
		case 93:
			goto tr80
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto st40
		}
		goto st0
	tr87:
//line internal/lexer/lexer.go.rl:65

		str, perr = strconv.Unquote(l.text())
		if perr != nil {
			perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
			{
				(l.p)++
				l.cs = 41
				goto _out
			}
		}
		l.nodes.Nest(value.Str(str))

		goto st41
	tr93:
//line internal/lexer/lexer.go.rl:73

		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			{
				(l.p)++
				l.cs = 41
				goto _out
			}
		}
		l.nodes.Nest(value.Int(n64))

		goto st41
	tr100:
//line internal/lexer/lexer.go.rl:80
		l.nodes.Nest(value.Bool(l.text() == "true"))
		goto st41
	st41:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof41
		}
	st_case_41:
//line internal/lexer/lexer.go:1602
		switch l.data[(l.p)] {
		case 32:
			goto st41
		case 34:
			goto tr77
		case 43:
			goto tr78
		case 45:
			goto tr78
		case 102:
			goto tr81
		case 116:
			goto tr82
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto tr79
			}
		case l.data[(l.p)] >= 9:
			goto st41
		}
		goto st0
	tr78:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st42
	st42:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof42
		}
	st_case_42:
//line internal/lexer/lexer.go:1635
		if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
			goto st43
		}
		goto st0
	tr79:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st43
	st43:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof43
		}
	st_case_43:
//line internal/lexer/lexer.go:1649
		switch l.data[(l.p)] {
		case 32:
			goto tr92
		case 44:
			goto tr93
		case 93:
			goto tr94
		}
		switch {
		case l.data[(l.p)] > 13:
			if 48 <= l.data[(l.p)] && l.data[(l.p)] <= 57 {
				goto st43
			}
		case l.data[(l.p)] >= 9:
			goto tr92
		}
		goto st0
	tr81:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st44
	st44:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof44
		}
	st_case_44:
//line internal/lexer/lexer.go:1676
		if l.data[(l.p)] == 97 {
			goto st45
		}
		goto st0
	st45:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof45
		}
	st_case_45:
		if l.data[(l.p)] == 108 {
			goto st46
		}
		goto st0
	st46:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof46
		}
	st_case_46:
		if l.data[(l.p)] == 115 {
			goto st47
		}
		goto st0
	st47:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof47
		}
	st_case_47:
		if l.data[(l.p)] == 101 {
			goto st48
		}
		goto st0
	st48:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch l.data[(l.p)] {
		case 32:
			goto tr99
		case 44:
			goto tr100
		case 93:
			goto tr101
		}
		if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
			goto tr99
		}
		goto st0
	tr82:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st49
	st49:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof49
		}
	st_case_49:
//line internal/lexer/lexer.go:1734
		if l.data[(l.p)] == 114 {
			goto st50
		}
		goto st0
	st50:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof50
		}
	st_case_50:
		if l.data[(l.p)] == 117 {
			goto st47
		}
		goto st0
	st51:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof51
		}
	st_case_51:
		goto st38
	tr57:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st52
	st52:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof52
		}
	st_case_52:
//line internal/lexer/lexer.go:1763
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		case 97:
			goto st53
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 98 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st53:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		case 108:
			goto st54
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st54:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		case 115:
			goto st55
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st55:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		case 101:
			goto st56
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st56:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch l.data[(l.p)] {
		case 32:
			goto tr107
		case 41:
			goto tr108
		case 44:
			goto tr109
		case 95:
			goto st35
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr107
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	tr58:
//line internal/lexer/lexer.go.rl:59
		l.pb = l.p
		goto st57
	st57:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof57
		}
	st_case_57:
//line internal/lexer/lexer.go:1933
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		case 114:
			goto st58
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st58:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch l.data[(l.p)] {
		case 32:
			goto tr71
		case 40:
			goto tr72
		case 95:
			goto st35
		case 117:
			goto st55
		}
		switch {
		case l.data[(l.p)] < 48:
			if 9 <= l.data[(l.p)] && l.data[(l.p)] <= 13 {
				goto tr71
			}
		case l.data[(l.p)] > 57:
			switch {
			case l.data[(l.p)] > 90:
				if 97 <= l.data[(l.p)] && l.data[(l.p)] <= 122 {
					goto st35
				}
			case l.data[(l.p)] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st59:
		if (l.p)++; (l.p) == (l.pe) {
			goto _test_eof59
		}
	st_case_59:
		goto st29
	st_out:
	_test_eof1:
		l.cs = 1
		goto _test_eof
	_test_eof2:
		l.cs = 2
		goto _test_eof
	_test_eof60:
		l.cs = 60
		goto _test_eof
	_test_eof3:
		l.cs = 3
		goto _test_eof
	_test_eof4:
		l.cs = 4
		goto _test_eof
	_test_eof61:
		l.cs = 61
		goto _test_eof
	_test_eof5:
		l.cs = 5
		goto _test_eof
	_test_eof6:
		l.cs = 6
		goto _test_eof
	_test_eof62:
		l.cs = 62
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
	_test_eof20:
		l.cs = 20
		goto _test_eof
	_test_eof21:
		l.cs = 21
		goto _test_eof
	_test_eof22:
		l.cs = 22
		goto _test_eof
	_test_eof23:
		l.cs = 23
		goto _test_eof
	_test_eof24:
		l.cs = 24
		goto _test_eof
	_test_eof25:
		l.cs = 25
		goto _test_eof
	_test_eof63:
		l.cs = 63
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
	_test_eof64:
		l.cs = 64
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

	_test_eof:
		{
		}
		if (l.p) == (l.eof) {
			switch l.cs {
			case 60:
//line internal/lexer/lexer.go.rl:65

				str, perr = strconv.Unquote(l.text())
				if perr != nil {
					perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
					{
						(l.p)++
						l.cs = 0
						goto _out
					}
				}
				l.nodes.Nest(value.Str(str))

			case 61:
//line internal/lexer/lexer.go.rl:73

				n64, perr = strconv.ParseInt(l.text(), 10, 64)
				if perr != nil {
					{
						(l.p)++
						l.cs = 0
						goto _out
					}
				}
				l.nodes.Nest(value.Int(n64))

			case 63:
//line internal/lexer/lexer.go.rl:80
				l.nodes.Nest(value.Bool(l.text() == "true"))
//line internal/lexer/lexer.go:2092
			}
		}

	_out:
		{
		}
	}

//line internal/lexer/lexer.go.rl:116

	if l.top > 0 {
		return nil, fmt.Errorf("stack parsing error at %d", l.pb)
	}

	if l.cs < 60 {
		if perr != nil {
			return nil, fmt.Errorf("parsing error at %d: %w", l.pb, perr)
		}
		return nil, fmt.Errorf("token parsing error at %d", l.pb)
	}

	return l.nodes.Top(), nil
}
