package proto

import (
	"io/ioutil"
	"os"

	"github.com/fabulousduck/proto/src/lexer"
	"github.com/fabulousduck/proto/src/tokens"
)

//Proto : Defines the global attributes of the interpreter
type Proto struct {
	Tokens   []*tokens.Token
	HadError bool
}

//NewProto : Creates a new proto instance
func NewProto() *Proto {
	return new(Proto)
}

//RunFile : Interprets a given file
func (proto *Proto) RunFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	proto.run(string(file), filename)
	if proto.HadError {
		os.Exit(65)
	}
}

func (proto *Proto) run(sourceCode string, filename string) {
	l := new(lexer.Lexer)
	l.Lex(sourceCode, filename)
	proto.Tokens = l.Tokens
	// p := NewParser(filename)
	// p.ast, _ = p.parse(l.tokens)
	// i := newInterpreter()
	// i.interpret(p.ast)

}
