package main

import (
	"goblk/chain"
	"fmt"
	"encoding/json"
  "goblk/proto/gopkg/block"
)

func main() {
	a := chain.NewChain()
  c := block.Entry{Previous: "a", From: "b", To:"c", Amount: 100}
}
