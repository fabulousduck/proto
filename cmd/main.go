package main

import (
	"github.com/fabulousduck/proto"
)

func main() {
	p := proto.NewProto()
	p.RunFile("examples/test.po")
}
