package lexer

import (
	"fmt"
	"strconv"

	"github.com/regeda/expr/internal/ast"
	"github.com/regeda/expr/internal/ast/value"
	"github.com/regeda/expr/internal/ast/stack"
)

%%{
	machine lexer;
	write data;
}%%

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

%%{
	access l.;

	variable p l.p;
	variable pe l.pe;
	variable eof l.eof;

	action mark { l.pb = l.p }

	action vm_go_up { l.nodes.Pop() }

	action vm_call { l.nodes.Push(l.nodes.Nest(value.Call(l.text()))) }
	action vm_arr { l.nodes.Push(l.nodes.Nest(value.Arr())) }
	action vm_str { l.nodes.Nest(value.Str(l.text())) }
	action vm_int {
		n64, perr = strconv.ParseInt(l.text(), 10, 64)
		if perr != nil {
			fbreak;
		}
		l.nodes.Nest(value.Int(n64))
	}
	action vm_bool { l.nodes.Nest(value.Bool(l.text() == "true")) }

	not_dquote = [^"\\];
	esc_smth = /\\./;

	str_body = not_dquote | esc_smth;
	func_name = [A-Za-z_][A-Za-z_0-9]*;

	num = ( ('+'|'-')? digit+ ) >mark %vm_int;
	str = '"' str_body* >mark %vm_str '"';
	bool = ('true' | 'false') >mark %vm_bool;
	func = (func_name & !bool) >mark %vm_call;

	scalar = space* (str | num | bool);

	comma = space* ',';

	# round brackets
	orbr = space* '(';
	crbr = space* ')';

	# square brackets
	osbr = space* '[';
	csbr = space* ']';

	array = osbr @vm_arr ( scalar (comma scalar)* )? csbr @vm_go_up;
	opt = scalar | array;
	invoke = space* func orbr @{ fcall invoke_rest; };
	opts = opt | invoke;
	invoke_rest := ( opts (comma opts)* )? crbr @vm_go_up @{ fret; };

	main := scalar | array | invoke;

	write init;
	write exec;
}%%
	if l.top > 0 {
		return nil, fmt.Errorf("stack parsing error at %d", l.pb)
	}

	if l.cs < %%{ write first_final; }%% {
		if perr != nil {
			return nil, fmt.Errorf("parse error at %d: %w", l.pb, perr)
		}
		return nil, fmt.Errorf("token parsing error at %d", l.pb)
	}

	return l.nodes.Top(), nil
}
