
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
const lexer_first_final int = 61
const lexer_error int = 0

const lexer_en_invoke_rest int = 29
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
	case 61:
		goto st61
	case 4:
		goto st4
	case 5:
		goto st5
	case 62:
		goto st62
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
	case 25:
		goto st25
	case 26:
		goto st26
	case 63:
		goto st63
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
	case 64:
		goto st64
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
	case 60:
		goto st60
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
	case 61:
		goto st_case_61
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 62:
		goto st_case_62
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
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 63:
		goto st_case_63
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
	case 64:
		goto st_case_64
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
	case 60:
		goto st_case_60
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
		switch  l.data[( l.p)] {
		case 34:
			goto tr10
		case 92:
			goto tr11
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
//line internal/lexer/lexer.go:410
		switch  l.data[( l.p)] {
		case 34:
			goto tr13
		case 92:
			goto st4
		}
		goto st3
tr10:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st61
tr13:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st61
tr17:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
//line internal/lexer/lexer.go.rl:90
 { l.stack[ l.top] = 61;  l.top++; goto st29 } 
	goto st61
tr20:
//line internal/lexer/lexer.go.rl:90
 { l.stack[ l.top] = 61;  l.top++; goto st29 } 
	goto st61
tr25:
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st61
tr39:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st61
tr46:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st61
	st61:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof61
		}
	st_case_61:
//line internal/lexer/lexer.go:459
		goto st0
tr11:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st4
	st4:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof4
		}
	st_case_4:
//line internal/lexer/lexer.go:470
		goto st3
tr3:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st5
	st5:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof5
		}
	st_case_5:
//line internal/lexer/lexer.go:481
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st62
		}
		goto st0
tr4:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st62
	st62:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof62
		}
	st_case_62:
//line internal/lexer/lexer.go:495
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st62
		}
		goto st0
tr5:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st6
	st6:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof6
		}
	st_case_6:
//line internal/lexer/lexer.go:509
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
tr16:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
	goto st7
	st7:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof7
		}
	st_case_7:
//line internal/lexer/lexer.go:545
		switch  l.data[( l.p)] {
		case 32:
			goto st7
		case 40:
			goto tr20
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st7
		}
		goto st0
tr6:
//line internal/lexer/lexer.go.rl:60
 l.nodes.Push(l.nodes.Nest(value.Arr())) 
	goto st8
	st8:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof8
		}
	st_case_8:
//line internal/lexer/lexer.go:565
		switch  l.data[( l.p)] {
		case 32:
			goto st8
		case 34:
			goto st9
		case 43:
			goto tr23
		case 45:
			goto tr23
		case 93:
			goto tr25
		case 102:
			goto tr26
		case 116:
			goto tr27
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr24
			}
		case  l.data[( l.p)] >= 9:
			goto st8
		}
		goto st0
	st9:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch  l.data[( l.p)] {
		case 34:
			goto tr29
		case 92:
			goto tr30
		}
		goto tr28
tr28:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st10
	st10:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof10
		}
	st_case_10:
//line internal/lexer/lexer.go:612
		switch  l.data[( l.p)] {
		case 34:
			goto tr32
		case 92:
			goto st22
		}
		goto st10
tr29:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st11
tr32:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st11
tr37:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st11
tr44:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st11
	st11:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof11
		}
	st_case_11:
//line internal/lexer/lexer.go:643
		switch  l.data[( l.p)] {
		case 32:
			goto st11
		case 44:
			goto st12
		case 93:
			goto tr25
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st11
		}
		goto st0
tr38:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st12
tr45:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st12
	st12:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof12
		}
	st_case_12:
//line internal/lexer/lexer.go:669
		switch  l.data[( l.p)] {
		case 32:
			goto st12
		case 34:
			goto st9
		case 43:
			goto tr23
		case 45:
			goto tr23
		case 102:
			goto tr26
		case 116:
			goto tr27
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr24
			}
		case  l.data[( l.p)] >= 9:
			goto st12
		}
		goto st0
tr23:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st13
	st13:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof13
		}
	st_case_13:
//line internal/lexer/lexer.go:702
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st14
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
//line internal/lexer/lexer.go:716
		switch  l.data[( l.p)] {
		case 32:
			goto tr37
		case 44:
			goto tr38
		case 93:
			goto tr39
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st14
			}
		case  l.data[( l.p)] >= 9:
			goto tr37
		}
		goto st0
tr26:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st15
	st15:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof15
		}
	st_case_15:
//line internal/lexer/lexer.go:743
		if  l.data[( l.p)] == 97 {
			goto st16
		}
		goto st0
	st16:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof16
		}
	st_case_16:
		if  l.data[( l.p)] == 108 {
			goto st17
		}
		goto st0
	st17:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof17
		}
	st_case_17:
		if  l.data[( l.p)] == 115 {
			goto st18
		}
		goto st0
	st18:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof18
		}
	st_case_18:
		if  l.data[( l.p)] == 101 {
			goto st19
		}
		goto st0
	st19:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  l.data[( l.p)] {
		case 32:
			goto tr44
		case 44:
			goto tr45
		case 93:
			goto tr46
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto tr44
		}
		goto st0
tr27:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st20
	st20:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof20
		}
	st_case_20:
//line internal/lexer/lexer.go:801
		if  l.data[( l.p)] == 114 {
			goto st21
		}
		goto st0
	st21:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof21
		}
	st_case_21:
		if  l.data[( l.p)] == 117 {
			goto st18
		}
		goto st0
tr30:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st22
	st22:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof22
		}
	st_case_22:
//line internal/lexer/lexer.go:824
		goto st10
tr7:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st23
	st23:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof23
		}
	st_case_23:
//line internal/lexer/lexer.go:835
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		case 97:
			goto st24
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 98 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st24:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof24
		}
	st_case_24:
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		case 108:
			goto st25
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st25:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		case 115:
			goto st26
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st26:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		case 101:
			goto st63
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st63:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof63
		}
	st_case_63:
		if  l.data[( l.p)] == 95 {
			goto st6
		}
		switch {
		case  l.data[( l.p)] < 65:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st6
			}
		case  l.data[( l.p)] > 90:
			if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		goto st0
tr8:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st27
	st27:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof27
		}
	st_case_27:
//line internal/lexer/lexer.go:993
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		case 114:
			goto st28
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st28:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof28
		}
	st_case_28:
		switch  l.data[( l.p)] {
		case 32:
			goto tr16
		case 40:
			goto tr17
		case 95:
			goto st6
		case 117:
			goto st26
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr16
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st6
				}
			case  l.data[( l.p)] >= 65:
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st29:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof29
		}
	st_case_29:
		switch  l.data[( l.p)] {
		case 32:
			goto st29
		case 34:
			goto st30
		case 41:
			goto tr55
		case 43:
			goto tr56
		case 45:
			goto tr56
		case 91:
			goto tr59
		case 95:
			goto tr58
		case 102:
			goto tr60
		case 116:
			goto tr61
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto st29
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto tr58
				}
			case  l.data[( l.p)] >= 65:
				goto tr58
			}
		default:
			goto tr57
		}
		goto st0
	st30:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof30
		}
	st_case_30:
		switch  l.data[( l.p)] {
		case 34:
			goto tr63
		case 92:
			goto tr64
		}
		goto tr62
tr62:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st31
	st31:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof31
		}
	st_case_31:
//line internal/lexer/lexer.go:1119
		switch  l.data[( l.p)] {
		case 34:
			goto tr66
		case 92:
			goto st60
		}
		goto st31
tr63:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st32
tr66:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st32
tr75:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
//line internal/lexer/lexer.go.rl:90
 { l.stack[ l.top] = 32;  l.top++; goto st29 } 
	goto st32
tr78:
//line internal/lexer/lexer.go.rl:90
 { l.stack[ l.top] = 32;  l.top++; goto st29 } 
	goto st32
tr83:
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st32
tr71:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st32
tr97:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st32
tr110:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st32
tr104:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
	goto st32
	st32:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof32
		}
	st_case_32:
//line internal/lexer/lexer.go:1176
		switch  l.data[( l.p)] {
		case 32:
			goto st32
		case 41:
			goto tr55
		case 44:
			goto st33
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st32
		}
		goto st0
tr55:
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
//line internal/lexer/lexer.go.rl:92
 { l.top--;  l.cs =  l.stack[ l.top];goto _again } 
	goto st64
tr72:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
//line internal/lexer/lexer.go.rl:92
 { l.top--;  l.cs =  l.stack[ l.top];goto _again } 
	goto st64
tr111:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go.rl:57
 l.nodes.Pop() 
//line internal/lexer/lexer.go.rl:92
 { l.top--;  l.cs =  l.stack[ l.top];goto _again } 
	goto st64
	st64:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof64
		}
	st_case_64:
//line internal/lexer/lexer.go:1216
		goto st0
tr73:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st33
tr112:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st33
	st33:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof33
		}
	st_case_33:
//line internal/lexer/lexer.go:1231
		switch  l.data[( l.p)] {
		case 32:
			goto st33
		case 34:
			goto st30
		case 43:
			goto tr56
		case 45:
			goto tr56
		case 91:
			goto tr59
		case 95:
			goto tr58
		case 102:
			goto tr60
		case 116:
			goto tr61
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto st33
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto tr58
				}
			case  l.data[( l.p)] >= 65:
				goto tr58
			}
		default:
			goto tr57
		}
		goto st0
tr56:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st34
	st34:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof34
		}
	st_case_34:
//line internal/lexer/lexer.go:1277
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st35
		}
		goto st0
tr57:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st35
	st35:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof35
		}
	st_case_35:
//line internal/lexer/lexer.go:1291
		switch  l.data[( l.p)] {
		case 32:
			goto tr71
		case 41:
			goto tr72
		case 44:
			goto tr73
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st35
			}
		case  l.data[( l.p)] >= 9:
			goto tr71
		}
		goto st0
tr58:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st36
	st36:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof36
		}
	st_case_36:
//line internal/lexer/lexer.go:1318
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
tr74:
//line internal/lexer/lexer.go.rl:59
 l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) 
	goto st37
	st37:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof37
		}
	st_case_37:
//line internal/lexer/lexer.go:1354
		switch  l.data[( l.p)] {
		case 32:
			goto st37
		case 40:
			goto tr78
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st37
		}
		goto st0
tr59:
//line internal/lexer/lexer.go.rl:60
 l.nodes.Push(l.nodes.Nest(value.Arr())) 
	goto st38
	st38:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof38
		}
	st_case_38:
//line internal/lexer/lexer.go:1374
		switch  l.data[( l.p)] {
		case 32:
			goto st38
		case 34:
			goto st39
		case 43:
			goto tr81
		case 45:
			goto tr81
		case 93:
			goto tr83
		case 102:
			goto tr84
		case 116:
			goto tr85
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr82
			}
		case  l.data[( l.p)] >= 9:
			goto st38
		}
		goto st0
	st39:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof39
		}
	st_case_39:
		switch  l.data[( l.p)] {
		case 34:
			goto tr87
		case 92:
			goto tr88
		}
		goto tr86
tr86:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st40
	st40:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof40
		}
	st_case_40:
//line internal/lexer/lexer.go:1421
		switch  l.data[( l.p)] {
		case 34:
			goto tr90
		case 92:
			goto st52
		}
		goto st40
tr87:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st41
tr90:
//line internal/lexer/lexer.go.rl:61
 l.nodes.Nest(value.Str(l.text())) 
	goto st41
tr95:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st41
tr102:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st41
	st41:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof41
		}
	st_case_41:
//line internal/lexer/lexer.go:1452
		switch  l.data[( l.p)] {
		case 32:
			goto st41
		case 44:
			goto st42
		case 93:
			goto tr83
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto st41
		}
		goto st0
tr96:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
	goto st42
tr103:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
	goto st42
	st42:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof42
		}
	st_case_42:
//line internal/lexer/lexer.go:1478
		switch  l.data[( l.p)] {
		case 32:
			goto st42
		case 34:
			goto st39
		case 43:
			goto tr81
		case 45:
			goto tr81
		case 102:
			goto tr84
		case 116:
			goto tr85
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto tr82
			}
		case  l.data[( l.p)] >= 9:
			goto st42
		}
		goto st0
tr81:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st43
	st43:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof43
		}
	st_case_43:
//line internal/lexer/lexer.go:1511
		if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
			goto st44
		}
		goto st0
tr82:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st44
	st44:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof44
		}
	st_case_44:
//line internal/lexer/lexer.go:1525
		switch  l.data[( l.p)] {
		case 32:
			goto tr95
		case 44:
			goto tr96
		case 93:
			goto tr97
		}
		switch {
		case  l.data[( l.p)] > 13:
			if 48 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 57 {
				goto st44
			}
		case  l.data[( l.p)] >= 9:
			goto tr95
		}
		goto st0
tr84:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st45
	st45:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof45
		}
	st_case_45:
//line internal/lexer/lexer.go:1552
		if  l.data[( l.p)] == 97 {
			goto st46
		}
		goto st0
	st46:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof46
		}
	st_case_46:
		if  l.data[( l.p)] == 108 {
			goto st47
		}
		goto st0
	st47:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof47
		}
	st_case_47:
		if  l.data[( l.p)] == 115 {
			goto st48
		}
		goto st0
	st48:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof48
		}
	st_case_48:
		if  l.data[( l.p)] == 101 {
			goto st49
		}
		goto st0
	st49:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof49
		}
	st_case_49:
		switch  l.data[( l.p)] {
		case 32:
			goto tr102
		case 44:
			goto tr103
		case 93:
			goto tr104
		}
		if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
			goto tr102
		}
		goto st0
tr85:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st50
	st50:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof50
		}
	st_case_50:
//line internal/lexer/lexer.go:1610
		if  l.data[( l.p)] == 114 {
			goto st51
		}
		goto st0
	st51:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof51
		}
	st_case_51:
		if  l.data[( l.p)] == 117 {
			goto st48
		}
		goto st0
tr88:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st52
	st52:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof52
		}
	st_case_52:
//line internal/lexer/lexer.go:1633
		goto st40
tr60:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st53
	st53:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof53
		}
	st_case_53:
//line internal/lexer/lexer.go:1644
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		case 97:
			goto st54
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 98 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
	st54:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		case 108:
			goto st55
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
	st55:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		case 115:
			goto st56
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
	st56:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		case 101:
			goto st57
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
	st57:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch  l.data[( l.p)] {
		case 32:
			goto tr110
		case 41:
			goto tr111
		case 44:
			goto tr112
		case 95:
			goto st36
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr110
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
tr61:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st58
	st58:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof58
		}
	st_case_58:
//line internal/lexer/lexer.go:1814
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		case 114:
			goto st59
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
	st59:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch  l.data[( l.p)] {
		case 32:
			goto tr74
		case 40:
			goto tr75
		case 95:
			goto st36
		case 117:
			goto st56
		}
		switch {
		case  l.data[( l.p)] < 48:
			if 9 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 13 {
				goto tr74
			}
		case  l.data[( l.p)] > 57:
			switch {
			case  l.data[( l.p)] > 90:
				if 97 <=  l.data[( l.p)] &&  l.data[( l.p)] <= 122 {
					goto st36
				}
			case  l.data[( l.p)] >= 65:
				goto st36
			}
		default:
			goto st36
		}
		goto st0
tr64:
//line internal/lexer/lexer.go.rl:55
 l.pb = l.p 
	goto st60
	st60:
		if ( l.p)++; ( l.p) == ( l.pe) {
			goto _test_eof60
		}
	st_case_60:
//line internal/lexer/lexer.go:1885
		goto st31
	st_out:
	_test_eof1:  l.cs = 1; goto _test_eof
	_test_eof2:  l.cs = 2; goto _test_eof
	_test_eof3:  l.cs = 3; goto _test_eof
	_test_eof61:  l.cs = 61; goto _test_eof
	_test_eof4:  l.cs = 4; goto _test_eof
	_test_eof5:  l.cs = 5; goto _test_eof
	_test_eof62:  l.cs = 62; goto _test_eof
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
	_test_eof25:  l.cs = 25; goto _test_eof
	_test_eof26:  l.cs = 26; goto _test_eof
	_test_eof63:  l.cs = 63; goto _test_eof
	_test_eof27:  l.cs = 27; goto _test_eof
	_test_eof28:  l.cs = 28; goto _test_eof
	_test_eof29:  l.cs = 29; goto _test_eof
	_test_eof30:  l.cs = 30; goto _test_eof
	_test_eof31:  l.cs = 31; goto _test_eof
	_test_eof32:  l.cs = 32; goto _test_eof
	_test_eof64:  l.cs = 64; goto _test_eof
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
	_test_eof57:  l.cs = 57; goto _test_eof
	_test_eof58:  l.cs = 58; goto _test_eof
	_test_eof59:  l.cs = 59; goto _test_eof
	_test_eof60:  l.cs = 60; goto _test_eof

	_test_eof: {}
	if ( l.p) == ( l.eof) {
		switch  l.cs {
		case 62:
//line internal/lexer/lexer.go.rl:62
 l.nodes.Nest(value.Atoi(l.text())) 
		case 63:
//line internal/lexer/lexer.go.rl:63
 l.nodes.Nest(value.Bool(l.text() == "true")) 
//line internal/lexer/lexer.go:1962
		}
	}

	_out: {}
	}

//line internal/lexer/lexer.go.rl:98

	if l.top > 0 {
		return nil, fmt.Errorf("stack parse error: %s", l.data)
	}

	if l.cs < 61 {
		return nil, fmt.Errorf("token parse error: %s", l.data)
	}

	return l.nodes.Top(), nil
}
