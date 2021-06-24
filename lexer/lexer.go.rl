package lexer

import (
  "github.com/pkg/errors"
)

%%{
  machine lexer;
  write data;
}%%

type Lexer struct {
  p, pe, pb, eof  int
  cs, top         int
  data            []byte
  stack           [1024]int
  rpn, ops, caps  nvec
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

%%{
  access l.;

  variable p l.p;
  variable pe l.pe;
  variable eof l.eof;

  action return { fret; }
  action mark { l.pb = l.p }
  action opts_first_param { l.caps.push(l.caps.pop().incCap()) }
  action rotate_comma { l.rotateComma() }
  action pop_caps { l.popCaps() }
  action push_invoke { l.pushInvoke() }
  action push_vector { l.pushVector() }
  action open_pths { l.openPths() }
  action close_pths { l.closePths() }
  action push_ident { l.pushIdent() }
  action push_str {
    str, perr = strconv.Unquote(l.text())
    if perr != nil {
      perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
      fbreak;
    }
    l.pushStr(str)
  }
  action push_int {
    n64, perr = strconv.ParseInt(l.text(), 10, 64)
    if perr != nil {
      fbreak;
    }
    l.pushInt(n64)
  }
  action push_true { l.pushTrue() }
  action push_false { l.pushFalse() }
  action push_math_op { l.pushMathOp() }

  not_dquote = [^"\\];
  esc_smth = /\\./;

  str_body = not_dquote | esc_smth;

  alpha_u = alpha | '_';
  alnum_u = alnum | '_';

  int = ([\+\-]? digit+) >mark %push_int;
  str = ('"' str_body* '"') >mark %push_str;
  true = 'true' %push_true;
  false = 'false' %push_false;
  bool = true | false;
  keywords = bool;
  ident = (alpha_u alnum_u* & !keywords) >mark %push_ident;

  ws = space*;

  comma = ',';

  # round brackets
  orbr = '(';
  crbr = ')';

  # square brackets
  osbr = '[';
  csbr = ']';

  vector = osbr @push_vector @{ fcall vector_rest; };
  invoke = ident ws orbr @push_invoke @{ fcall invoke_rest; };
  pths = orbr @open_pths @{ fcall pths_rest; };

  math_op = [\+\-\*\/] >mark %push_math_op;
  math_pred = int | invoke | pths;
  math = math_pred (ws math_op ws math_pred)*;

  opts = math | vector | str | bool;

  opts_delim = opts >opts_first_param (ws comma @rotate_comma ws opts)*;

  pths_rest := ws math ws crbr @close_pths @return;
  vector_rest := ws opts_delim? ws csbr @pop_caps @return;
  invoke_rest := ws opts_delim? ws crbr @pop_caps @return;

  main := ws opts ws;

  write init;
  write exec;
}%%
  if l.top > 0 {
    return nil, fmt.Errorf("stack parsing error at %d", l.pb)
  }

  if l.cs < %%{ write first_final; }%% {
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
