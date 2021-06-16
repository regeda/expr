package tokenz

%%{
  machine tokenz;
  write data;
}%%

type Tokenz struct {
  data []byte
  cs int
  p, pe, eof int
  ts, te, act int
  tokens []Token
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

%%{
  access t.;

  variable p t.p;
  variable pe t.pe;
  variable eof t.eof;

  not_dquote = [^"\\];
  esc_smth = /\\./;
  str_body = not_dquote | esc_smth;

  alpha_u = alpha | '_';
  alnum_u = alnum | '_';

  main := |*
    ('+'|'-')? digit+ {
      t.addTxt(TkInt)
    };

    '"' str_body* '"' {
      t.addTxt(TkStr)
    };

    'true' {
      t.addTk(TkTrue)
    };

    'false' {
      t.addTk(TkFalse)
    };

    alpha_u alnum_u* {
      t.addTxt(TkIdent)
    };

    [\[\]\(\)\,] {
      t.addTxt(TkPunct)
    };

    space;
  *|;

  write init;
  write exec;
}%%

  if t.cs < %%{ write first_final; }%% {
    return nil, fmt.Errorf("token parsing error at %d", t.ts)
  }

  return t.tokens, nil
}
