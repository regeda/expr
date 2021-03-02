package lexer

import (
	"fmt"

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
	action vm_int { l.nodes.Nest(value.Atoi(l.text())) }
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
		return nil, fmt.Errorf("stack parse error: %s", l.data)
	}

	if l.cs < %%{ write first_final; }%% {
		return nil, fmt.Errorf("token parse error: %s", l.data)
	}

	return l.nodes.Top(), nil
}
