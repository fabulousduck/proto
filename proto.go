package proto

import (
	"io/ioutil"
	"os"
)

//proto : Defines the global attributes of the interpreter
type proto struct {
	Tokens   []*token
	HadError bool
}

//Newproto : Creates a new proto instance
func Newproto() *proto {
	return new(proto)
}

//RunFile : Interprets a given file
func (proto *proto) RunFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	proto.run(string(file), filename)
	if proto.HadError {
		os.Exit(65)
	}
}

func (proto *proto) run(sourceCode string, filename string) {
	l := new(lexer)
	l.lex(sourceCode, filename)
	p := NewParser(filename)
	p.ast, _ = p.parse(l.tokens)
	i := newInterpreter()
	i.interpret(p.ast)

}
