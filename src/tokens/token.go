package tokens

type Token struct {
	Value, Type string
	Line, Col   int
}
