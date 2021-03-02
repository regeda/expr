
//line internal/lexer/lexer.go.rl:1
package lexer

import (
	"fmt"

	"github.com/regeda/expr/internal/ast"
	"github.com/regeda/expr/internal/ast/value"
	"github.com/regeda/expr/internal/ast/stack"
)


//line internal/lexer/lexer.go:15
const lexer_start int = 1
const lexer_first_final int = 57
const lexer_error int = 0

const lexer_en_invoke_rest int = 27
const lexer_en_main int = 1


//line internal/lexer/lexer.go.rl:14


type Lexer struct {
	data       []byte
	cs         int
	p, pe, eof int
	pb         int
	top        int
	stack      []int
	err        error
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
	l.err = nil
	l.nodes.Reset()

	l.nodes.Push(value.Exit())


//line internal/lexer/lexer.go:60
	{
	 l.cs = lexer_start
	 l.top = 0
	}

//line internal/lexer/lexer.go:66
	{
	if ( l.p) == ( l.pe) {
		goto _test_eof
	}
	goto _resume

_again:
	switch  l.cs {
	case 1:
		goto st1
	case 0:
		goto st0
	case 2:
		goto st2
	case 3:
		goto st3
	case 57:
		goto st57
	case 4:
		goto st4
	case 58:
		goto st58
	case 5:
		goto st5
	case 6:
		goto st6
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
	case 59:
		goto st59
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
	case 60:
		goto st60
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
	}

	if ( l.p)++; ( l.p) == ( l.pe) {
		goto _test_eof
	}
_resume:
	switch  l.cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 57:
		goto st_case_57
	case 4:
		goto st_case_4
	case 58:
		goto st_case_58
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
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
	case 59:
		goto st_case_59
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
	case 60:
		goto st_case_60
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
	}
	goto st_out
	st1:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch  l.data[( l.p)] {
		case 32:
			goto st1
		case 34:
			goto st2
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
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto st1
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto tr5
				}
			case  l.data[( l.p)] >= 65:
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
	st2:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof2
		}
	st_case_2:
		if  l.data[( l.p)] == 34 {
			goto tr10
		}
		goto tr9
tr9:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st3
	st3:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof3
		}
	st_case_3:
//line internal/lexer/lexer.go:391
		if  l.data[( l.p)] == 34 {
			goto tr12
		}
		goto st3
tr10:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st57
tr12:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st57
tr15:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
//line internal/lexer/lexer.go.rl:88
 { l.stack[ l.top] = 57;  l.top++; goto st27 } 
	goto st57
tr18:
//line internal/lexer/lexer.go.rl:88
 { l.stack[ l.top] = 57;  l.top++; goto st27 } 
	goto st57
tr23:
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st57
tr35:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st57
tr42:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st57
	st57:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof57
		}
	st_case_57:
//line internal/lexer/lexer.go:437
		goto st0
tr3:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st4
	st4:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof4
		}
	st_case_4:
//line internal/lexer/lexer.go:448
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st58
		}
		goto st0
tr4:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st58
	st58:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof58
		}
	st_case_58:
//line internal/lexer/lexer.go:462
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st58
		}
		goto st0
tr5:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st5
	st5:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof5
		}
	st_case_5:
//line internal/lexer/lexer.go:476
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
tr14:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
	goto st6
	st6:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof6
		}
	st_case_6:
//line internal/lexer/lexer.go:512
		switch  l.data[( l.p)] {
		case 32:
			goto st6
		case 40:
			goto tr18
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st6
		}
		goto st0
tr6:
//line internal/lexer/lexer.go.rl:60
 l.nodes.Push(l.nodes.Nest(value.Arr())) 
	goto st7
	st7:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof7
		}
	st_case_7:
//line internal/lexer/lexer.go:532
		switch  l.data[( l.p)] {
		case 32:
			goto st7
		case 34:
			goto st8
		case 43:
			goto tr21
		case 45:
			goto tr21
		case 93:
			goto tr23
		case 102:
			goto tr24
		case 116:
			goto tr25
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr22
			}
		case  l.data[( l.p)] >= 9:
			goto st7
		}
		goto st0
	st8:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof8
		}
	st_case_8:
		if  l.data[( l.p)] == 34 {
			goto tr27
		}
		goto tr26
tr26:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st9
	st9:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof9
		}
	st_case_9:
//line internal/lexer/lexer.go:576
		if  l.data[( l.p)] == 34 {
			goto tr29
		}
		goto st9
tr27:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st10
tr29:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st10
tr33:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st10
tr40:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st10
	st10:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof10
		}
	st_case_10:
//line internal/lexer/lexer.go:604
		switch  l.data[( l.p)] {
		case 32:
			goto st10
		case 44:
			goto st11
		case 93:
			goto tr23
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st10
		}
		goto st0
tr34:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st11
tr41:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st11
	st11:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof11
		}
	st_case_11:
//line internal/lexer/lexer.go:630
		switch  l.data[( l.p)] {
		case 32:
			goto st11
		case 34:
			goto st8
		case 43:
			goto tr21
		case 45:
			goto tr21
		case 102:
			goto tr24
		case 116:
			goto tr25
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr22
			}
		case  l.data[( l.p)] >= 9:
			goto st11
		}
		goto st0
tr21:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st12
	st12:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof12
		}
	st_case_12:
//line internal/lexer/lexer.go:663
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st13
		}
		goto st0
tr22:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st13
	st13:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof13
		}
	st_case_13:
//line internal/lexer/lexer.go:677
		switch  l.data[( l.p)] {
		case 32:
			goto tr33
		case 44:
			goto tr34
		case 93:
			goto tr35
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st13
			}
		case  l.data[( l.p)] >= 9:
			goto tr33
		}
		goto st0
tr24:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st14
	st14:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof14
		}
	st_case_14:
//line internal/lexer/lexer.go:704
		if  l.data[( l.p)] == 97 {
			goto st15
		}
		goto st0
	st15:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof15
		}
	st_case_15:
		if  l.data[( l.p)] == 108 {
			goto st16
		}
		goto st0
	st16:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof16
		}
	st_case_16:
		if  l.data[( l.p)] == 115 {
			goto st17
		}
		goto st0
	st17:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof17
		}
	st_case_17:
		if  l.data[( l.p)] == 101 {
			goto st18
		}
		goto st0
	st18:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof18
		}
	st_case_18:
		switch  l.data[( l.p)] {
		case 32:
			goto tr40
		case 44:
			goto tr41
		case 93:
			goto tr42
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto tr40
		}
		goto st0
tr25:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st19
	st19:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof19
		}
	st_case_19:
//line internal/lexer/lexer.go:762
		if  l.data[( l.p)] == 114 {
			goto st20
		}
		goto st0
	st20:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof20
		}
	st_case_20:
		if  l.data[( l.p)] == 117 {
			goto st17
		}
		goto st0
tr7:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st21
	st21:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof21
		}
	st_case_21:
//line internal/lexer/lexer.go:785
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		case 97:
			goto st22
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 98 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st22:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof22
		}
	st_case_22:
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		case 108:
			goto st23
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st23:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof23
		}
	st_case_23:
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		case 115:
			goto st24
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st24:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		case 101:
			goto st59
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st59:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof59
		}
	st_case_59:
		if  l.data[( l.p)] == 95 {
			goto st5
		}
		switch {
		case  l.data[( l.p)] < 65:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st5
			}
		case  l.data[( l.p)] > 90:
			if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		goto st0
tr8:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st25
	st25:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof25
		}
	st_case_25:
//line internal/lexer/lexer.go:943
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		case 114:
			goto st26
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st26:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  l.data[( l.p)] {
		case 32:
			goto tr14
		case 40:
			goto tr15
		case 95:
			goto st5
		case 117:
			goto st24
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr14
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st5
				}
			case  l.data[( l.p)] >= 65:
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st27:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof27
		}
	st_case_27:
		switch  l.data[( l.p)] {
		case 32:
			goto st27
		case 34:
			goto st28
		case 41:
			goto tr51
		case 43:
			goto tr52
		case 45:
			goto tr52
		case 91:
			goto tr55
		case 95:
			goto tr54
		case 102:
			goto tr56
		case 116:
			goto tr57
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto st27
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto tr54
				}
			case  l.data[( l.p)] >= 65:
				goto tr54
			}
		default:
			goto tr53
		}
		goto st0
	st28:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof28
		}
	st_case_28:
		if  l.data[( l.p)] == 34 {
			goto tr59
		}
		goto tr58
tr58:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st29
	st29:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof29
		}
	st_case_29:
//line internal/lexer/lexer.go:1066
		if  l.data[( l.p)] == 34 {
			goto tr61
		}
		goto st29
tr59:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st30
tr61:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st30
tr69:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
//line internal/lexer/lexer.go.rl:88
 { l.stack[ l.top] = 30;  l.top++; goto st27 } 
	goto st30
tr72:
//line internal/lexer/lexer.go.rl:88
 { l.stack[ l.top] = 30;  l.top++; goto st27 } 
	goto st30
tr77:
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st30
tr65:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st30
tr89:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st30
tr102:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st30
tr96:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st30
	st30:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof30
		}
	st_case_30:
//line internal/lexer/lexer.go:1120
		switch  l.data[( l.p)] {
		case 32:
			goto st30
		case 41:
			goto tr51
		case 44:
			goto st31
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st30
		}
		goto st0
tr51:
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
//line internal/lexer/lexer.go.rl:90
 { l.top--;  l.cs =  l.stack[ l.top];goto _again } 
	goto st60
tr66:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
//line internal/lexer/lexer.go.rl:90
 { l.top--;  l.cs =  l.stack[ l.top];goto _again } 
	goto st60
tr103:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
//line internal/lexer/lexer.go.rl:90
 { l.top--;  l.cs =  l.stack[ l.top];goto _again } 
	goto st60
	st60:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof60
		}
	st_case_60:
//line internal/lexer/lexer.go:1160
		goto st0
tr67:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st31
tr104:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st31
	st31:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof31
		}
	st_case_31:
//line internal/lexer/lexer.go:1175
		switch  l.data[( l.p)] {
		case 32:
			goto st31
		case 34:
			goto st28
		case 43:
			goto tr52
		case 45:
			goto tr52
		case 91:
			goto tr55
		case 95:
			goto tr54
		case 102:
			goto tr56
		case 116:
			goto tr57
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto st31
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto tr54
				}
			case  l.data[( l.p)] >= 65:
				goto tr54
			}
		default:
			goto tr53
		}
		goto st0
tr52:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st32
	st32:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof32
		}
	st_case_32:
//line internal/lexer/lexer.go:1221
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st33
		}
		goto st0
tr53:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st33
	st33:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof33
		}
	st_case_33:
//line internal/lexer/lexer.go:1235
		switch  l.data[( l.p)] {
		case 32:
			goto tr65
		case 41:
			goto tr66
		case 44:
			goto tr67
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st33
			}
		case  l.data[( l.p)] >= 9:
			goto tr65
		}
		goto st0
tr54:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st34
	st34:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof34
		}
	st_case_34:
//line internal/lexer/lexer.go:1262
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
tr68:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
	goto st35
	st35:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof35
		}
	st_case_35:
//line internal/lexer/lexer.go:1298
		switch  l.data[( l.p)] {
		case 32:
			goto st35
		case 40:
			goto tr72
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st35
		}
		goto st0
tr55:
//line internal/lexer/lexer.go.rl:60
 l.nodes.Push(l.nodes.Nest(value.Arr())) 
	goto st36
	st36:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof36
		}
	st_case_36:
//line internal/lexer/lexer.go:1318
		switch  l.data[( l.p)] {
		case 32:
			goto st36
		case 34:
			goto st37
		case 43:
			goto tr75
		case 45:
			goto tr75
		case 93:
			goto tr77
		case 102:
			goto tr78
		case 116:
			goto tr79
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr76
			}
		case  l.data[( l.p)] >= 9:
			goto st36
		}
		goto st0
	st37:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof37
		}
	st_case_37:
		if  l.data[( l.p)] == 34 {
			goto tr81
		}
		goto tr80
tr80:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st38
	st38:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof38
		}
	st_case_38:
//line internal/lexer/lexer.go:1362
		if  l.data[( l.p)] == 34 {
			goto tr83
		}
		goto st38
tr81:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st39
tr83:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st39
tr87:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st39
tr94:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st39
	st39:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof39
		}
	st_case_39:
//line internal/lexer/lexer.go:1390
		switch  l.data[( l.p)] {
		case 32:
			goto st39
		case 44:
			goto st40
		case 93:
			goto tr77
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st39
		}
		goto st0
tr88:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st40
tr95:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st40
	st40:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof40
		}
	st_case_40:
//line internal/lexer/lexer.go:1416
		switch  l.data[( l.p)] {
		case 32:
			goto st40
		case 34:
			goto st37
		case 43:
			goto tr75
		case 45:
			goto tr75
		case 102:
			goto tr78
		case 116:
			goto tr79
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr76
			}
		case  l.data[( l.p)] >= 9:
			goto st40
		}
		goto st0
tr75:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st41
	st41:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof41
		}
	st_case_41:
//line internal/lexer/lexer.go:1449
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st42
		}
		goto st0
tr76:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st42
	st42:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof42
		}
	st_case_42:
//line internal/lexer/lexer.go:1463
		switch  l.data[( l.p)] {
		case 32:
			goto tr87
		case 44:
			goto tr88
		case 93:
			goto tr89
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st42
			}
		case  l.data[( l.p)] >= 9:
			goto tr87
		}
		goto st0
tr78:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st43
	st43:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof43
		}
	st_case_43:
//line internal/lexer/lexer.go:1490
		if  l.data[( l.p)] == 97 {
			goto st44
		}
		goto st0
	st44:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof44
		}
	st_case_44:
		if  l.data[( l.p)] == 108 {
			goto st45
		}
		goto st0
	st45:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof45
		}
	st_case_45:
		if  l.data[( l.p)] == 115 {
			goto st46
		}
		goto st0
	st46:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof46
		}
	st_case_46:
		if  l.data[( l.p)] == 101 {
			goto st47
		}
		goto st0
	st47:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch  l.data[( l.p)] {
		case 32:
			goto tr94
		case 44:
			goto tr95
		case 93:
			goto tr96
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto tr94
		}
		goto st0
tr79:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st48
	st48:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof48
		}
	st_case_48:
//line internal/lexer/lexer.go:1548
		if  l.data[( l.p)] == 114 {
			goto st49
		}
		goto st0
	st49:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof49
		}
	st_case_49:
		if  l.data[( l.p)] == 117 {
			goto st46
		}
		goto st0
tr56:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st50
	st50:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof50
		}
	st_case_50:
//line internal/lexer/lexer.go:1571
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		case 97:
			goto st51
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 98 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st51:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		case 108:
			goto st52
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st52:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		case 115:
			goto st53
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st53:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		case 101:
			goto st54
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st54:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch  l.data[( l.p)] {
		case 32:
			goto tr102
		case 41:
			goto tr103
		case 44:
			goto tr104
		case 95:
			goto st34
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr102
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
tr57:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st55
	st55:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof55
		}
	st_case_55:
//line internal/lexer/lexer.go:1741
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		case 114:
			goto st56
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st56:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch  l.data[( l.p)] {
		case 32:
			goto tr68
		case 40:
			goto tr69
		case 95:
			goto st34
		case 117:
			goto st53
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr68
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st34
				}
			case  l.data[( l.p)] >= 65:
				goto st34
			}
		default:
			goto st34
		}
		goto st0
	st_out:
	_test_eof1:  l.cs = 1; goto _test_eof
	_test_eof2:  l.cs = 2; goto _test_eof
	_test_eof3:  l.cs = 3; goto _test_eof
	_test_eof57:  l.cs = 57; goto _test_eof
	_test_eof4:  l.cs = 4; goto _test_eof
	_test_eof58:  l.cs = 58; goto _test_eof
	_test_eof5:  l.cs = 5; goto _test_eof
	_test_eof6:  l.cs = 6; goto _test_eof
	_test_eof7:  l.cs = 7; goto _test_eof
	_test_eof8:  l.cs = 8; goto _test_eof
	_test_eof9:  l.cs = 9; goto _test_eof
	_test_eof10:  l.cs = 10; goto _test_eof
	_test_eof11:  l.cs = 11; goto _test_eof
	_test_eof12:  l.cs = 12; goto _test_eof
	_test_eof13:  l.cs = 13; goto _test_eof
	_test_eof14:  l.cs = 14; goto _test_eof
	_test_eof15:  l.cs = 15; goto _test_eof
	_test_eof16:  l.cs = 16; goto _test_eof
	_test_eof17:  l.cs = 17; goto _test_eof
	_test_eof18:  l.cs = 18; goto _test_eof
	_test_eof19:  l.cs = 19; goto _test_eof
	_test_eof20:  l.cs = 20; goto _test_eof
	_test_eof21:  l.cs = 21; goto _test_eof
	_test_eof22:  l.cs = 22; goto _test_eof
	_test_eof23:  l.cs = 23; goto _test_eof
	_test_eof24:  l.cs = 24; goto _test_eof
	_test_eof59:  l.cs = 59; goto _test_eof
	_test_eof25:  l.cs = 25; goto _test_eof
	_test_eof26:  l.cs = 26; goto _test_eof
	_test_eof27:  l.cs = 27; goto _test_eof
	_test_eof28:  l.cs = 28; goto _test_eof
	_test_eof29:  l.cs = 29; goto _test_eof
	_test_eof30:  l.cs = 30; goto _test_eof
	_test_eof60:  l.cs = 60; goto _test_eof
	_test_eof31:  l.cs = 31; goto _test_eof
	_test_eof32:  l.cs = 32; goto _test_eof
	_test_eof33:  l.cs = 33; goto _test_eof
	_test_eof34:  l.cs = 34; goto _test_eof
	_test_eof35:  l.cs = 35; goto _test_eof
	_test_eof36:  l.cs = 36; goto _test_eof
	_test_eof37:  l.cs = 37; goto _test_eof
	_test_eof38:  l.cs = 38; goto _test_eof
	_test_eof39:  l.cs = 39; goto _test_eof
	_test_eof40:  l.cs = 40; goto _test_eof
	_test_eof41:  l.cs = 41; goto _test_eof
	_test_eof42:  l.cs = 42; goto _test_eof
	_test_eof43:  l.cs = 43; goto _test_eof
	_test_eof44:  l.cs = 44; goto _test_eof
	_test_eof45:  l.cs = 45; goto _test_eof
	_test_eof46:  l.cs = 46; goto _test_eof
	_test_eof47:  l.cs = 47; goto _test_eof
	_test_eof48:  l.cs = 48; goto _test_eof
	_test_eof49:  l.cs = 49; goto _test_eof
	_test_eof50:  l.cs = 50; goto _test_eof
	_test_eof51:  l.cs = 51; goto _test_eof
	_test_eof52:  l.cs = 52; goto _test_eof
	_test_eof53:  l.cs = 53; goto _test_eof
	_test_eof54:  l.cs = 54; goto _test_eof
	_test_eof55:  l.cs = 55; goto _test_eof
	_test_eof56:  l.cs = 56; goto _test_eof

	_test_eof: {}
	if ( l.p) == ( l.eof) {
		switch  l.cs {
		case 58:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
		case 59:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go:1874
		}
	}

	_out: {}
	}

//line internal/lexer/lexer.go.rl:96


	if l.top > 0 {
		return nil, fmt.Errorf("stack parse error: %s", l.data)
	}

	if l.cs < 57 {
		return nil, fmt.Errorf("token parse error: %s", l.data)
	}

	return l.nodes.Top(), nil
}
