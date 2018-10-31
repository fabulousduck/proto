package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fabulousduck/proto"
)

func main() {
	p := proto.NewProto()
	p.RunFile("examples/example.po")
	spew.Dump(p.Tokens)
}
