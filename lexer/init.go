package lexer

var Default = Lexer{}

func Parse(input []byte) ([]Node, error) {
	return Default.Parse(input)
}
