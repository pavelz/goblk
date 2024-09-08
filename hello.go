package main

import (
	"encoding/json"
	"fmt"
	"goblk/chain"
	"goblk/proto/gopkg/block"

	"google.golang.org/protobuf/internal/impl"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	a := chain.NewChain()
  b := block.Block{
    Entries: []*block.Entry{{Previous: "a", From: "b", To:"c", Amount: 100}},
  }

  a.AddBlock(b)
}
