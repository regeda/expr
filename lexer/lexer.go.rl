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
  ast             ast
}

func (l *Lexer) text() string {
  return string(l.data[l.pb:l.p])
}

func (l *Lexer) Parse(input []byte) ([]Node, error) {
  l.ast.reset()

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
  action opts_first_param { l.ast.incCaps() }
  action rotate_comma { l.ast.rotateComma() }
  action pop_caps { l.ast.popCaps() }
  action push_invoke { l.ast.pushInvoke() }
  action push_vector { l.ast.pushVector() }
  action open_pths { l.ast.openPths() }
  action close_pths { l.ast.closePths() }
  action push_ident { l.ast.pushIdent(l.text()) }
  action push_str {
    str, perr = strconv.Unquote(l.text())
    if perr != nil {
      perr = errors.Wrapf(perr, "strconv.Unquote %s", l.text())
      fbreak;
    }
    l.ast.pushStr(str)
  }
  action push_int {
    n64, perr = strconv.ParseInt(l.text(), 10, 64)
    if perr != nil {
      fbreak;
    }
    l.ast.pushInt(n64)
  }
  action push_true { l.ast.pushTrue() }
  action push_false { l.ast.pushFalse() }
  action push_math_op { l.ast.pushMathOp(l.data[l.pb]) }

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

  l.ast.rotateBreakOn()

  return l.ast.rpn, nil
}
