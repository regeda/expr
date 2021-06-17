//line tokenz/tokenz.go.rl:1
package tokenz

import "fmt"

//line tokenz/tokenz.go:7
const tokenz_start int = 4
const tokenz_first_final int = 4
const tokenz_error int = 0

const tokenz_en_main int = 4

//line tokenz/tokenz.go.rl:6

type Tokenz struct {
	data        []byte
	cs          int
	p, pe, eof  int
	ts, te, act int
	tokens      []Token
}

func (t *Tokenz) text() []byte {
	return t.data[t.ts:t.te]
}

func (t *Tokenz) addTk(tk Tk) {
	t.add(tk, nil)
}

func (t *Tokenz) addTxt(tk Tk) {
	t.add(tk, t.text())
}

func (t *Tokenz) add(tk Tk, d []byte) {
	t.tokens = append(t.tokens, Token{tk, d})
}

func (t *Tokenz) Parse(input []byte) ([]Token, error) {
	t.data = input
	t.p = 0
	t.pe = len(input)
	t.eof = len(input)
	t.tokens = t.tokens[:0]

//line tokenz/tokenz.go:50
	{
		t.cs = tokenz_start
		t.ts = 0
		t.te = 0
		t.act = 0
	}

//line tokenz/tokenz.go:58
	{
		if (t.p) == (t.pe) {
			goto _test_eof
		}
		switch t.cs {
		case 4:
			goto st_case_4
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
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
		}
		goto st_out
	tr1:
//line tokenz/tokenz.go.rl:58
		t.te = (t.p) + 1
		{
			t.addTxt(TkStr)
		}
		goto st4
	tr5:
//line tokenz/tokenz.go.rl:78
		t.te = (t.p) + 1

		goto st4
	tr6:
//line tokenz/tokenz.go.rl:74
		t.te = (t.p) + 1
		{
			t.addTxt(TkPunct)
		}
		goto st4
	tr11:
//line tokenz/tokenz.go.rl:54
		t.te = (t.p)
		(t.p)--
		{
			t.addTxt(TkInt)
		}
		goto st4
	tr12:
//line NONE:1
		switch t.act {
		case 3:
			{
				(t.p) = (t.te) - 1

				t.addTk(TkTrue)
			}
		case 4:
			{
				(t.p) = (t.te) - 1

				t.addTk(TkFalse)
			}
		case 5:
			{
				(t.p) = (t.te) - 1

				t.addTxt(TkIdent)
			}
		}

		goto st4
	tr13:
//line tokenz/tokenz.go.rl:70
		t.te = (t.p)
		(t.p)--
		{
			t.addTxt(TkIdent)
		}
		goto st4
	st4:
//line NONE:1
		t.ts = 0

		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof4
		}
	st_case_4:
//line NONE:1
		t.ts = (t.p)

//line tokenz/tokenz.go:161
		switch t.data[(t.p)] {
		case 32:
			goto tr5
		case 34:
			goto st1
		case 44:
			goto tr6
		case 91:
			goto tr6
		case 93:
			goto tr6
		case 95:
			goto tr8
		case 102:
			goto st7
		case 116:
			goto st11
		}
		switch {
		case t.data[(t.p)] < 43:
			switch {
			case t.data[(t.p)] > 13:
				if 40 <= t.data[(t.p)] && t.data[(t.p)] <= 41 {
					goto tr6
				}
			case t.data[(t.p)] >= 9:
				goto tr5
			}
		case t.data[(t.p)] > 45:
			switch {
			case t.data[(t.p)] < 65:
				if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
					goto st5
				}
			case t.data[(t.p)] > 90:
				if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
					goto tr8
				}
			default:
				goto tr8
			}
		default:
			goto st3
		}
		goto st0
	st_case_0:
	st0:
		t.cs = 0
		goto _out
	st1:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch t.data[(t.p)] {
		case 34:
			goto tr1
		case 92:
			goto st2
		}
		goto st1
	st2:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof2
		}
	st_case_2:
		goto st1
	st3:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof5
		}
	st_case_5:
		if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
			goto st5
		}
		goto tr11
	tr8:
//line NONE:1
		t.te = (t.p) + 1

//line tokenz/tokenz.go.rl:70
		t.act = 5
		goto st6
	tr17:
//line NONE:1
		t.te = (t.p) + 1

//line tokenz/tokenz.go.rl:66
		t.act = 4
		goto st6
	tr20:
//line NONE:1
		t.te = (t.p) + 1

//line tokenz/tokenz.go.rl:62
		t.act = 3
		goto st6
	st6:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof6
		}
	st_case_6:
//line tokenz/tokenz.go:273
		if t.data[(t.p)] == 95 {
			goto tr8
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr12
	st7:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 97:
			goto st8
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 98 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st8:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 108:
			goto st9
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st9:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 115:
			goto st10
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st10:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 101:
			goto tr17
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st11:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof11
		}
	st_case_11:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 114:
			goto st12
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st12:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof12
		}
	st_case_12:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 117:
			goto st13
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st13:
		if (t.p)++; (t.p) == (t.pe) {
			goto _test_eof13
		}
	st_case_13:
		switch t.data[(t.p)] {
		case 95:
			goto tr8
		case 101:
			goto tr20
		}
		switch {
		case t.data[(t.p)] < 65:
			if 48 <= t.data[(t.p)] && t.data[(t.p)] <= 57 {
				goto tr8
			}
		case t.data[(t.p)] > 90:
			if 97 <= t.data[(t.p)] && t.data[(t.p)] <= 122 {
				goto tr8
			}
		default:
			goto tr8
		}
		goto tr13
	st_out:
	_test_eof4:
		t.cs = 4
		goto _test_eof
	_test_eof1:
		t.cs = 1
		goto _test_eof
	_test_eof2:
		t.cs = 2
		goto _test_eof
	_test_eof3:
		t.cs = 3
		goto _test_eof
	_test_eof5:
		t.cs = 5
		goto _test_eof
	_test_eof6:
		t.cs = 6
		goto _test_eof
	_test_eof7:
		t.cs = 7
		goto _test_eof
	_test_eof8:
		t.cs = 8
		goto _test_eof
	_test_eof9:
		t.cs = 9
		goto _test_eof
	_test_eof10:
		t.cs = 10
		goto _test_eof
	_test_eof11:
		t.cs = 11
		goto _test_eof
	_test_eof12:
		t.cs = 12
		goto _test_eof
	_test_eof13:
		t.cs = 13
		goto _test_eof

	_test_eof:
		{
		}
		if (t.p) == (t.eof) {
			switch t.cs {
			case 5:
				goto tr11
			case 6:
				goto tr12
			case 7:
				goto tr13
			case 8:
				goto tr13
			case 9:
				goto tr13
			case 10:
				goto tr13
			case 11:
				goto tr13
			case 12:
				goto tr13
			case 13:
				goto tr13
			}
		}

	_out:
		{
		}
	}

//line tokenz/tokenz.go.rl:83

	if t.cs < 4 {
		return nil, fmt.Errorf("token parsing error at %d", t.ts)
	}

	return t.tokens, nil
}
