package main

import (
	"crypto/md5"
	"goblk/chain"
	"encoding/hex"
	"fmt"
	"encoding/json"
)

func main() {
	hash := md5.Sum([]byte("hex me"))
	text := hex.EncodeToString(hash[:])
	a := chain.NewChain()
  b := chain.Block{Previous: "two", From: "me", To: "who"}
  a.AddBlock(&b)
	a.WriteChain("chain.json")

	// load chain
	chain, err := chain.LoadChain("chain.json")

	if err != nil {
		println("there was an error: " + err.Error())
	}
	println(json.Marshal(*chain))
	
	fmt.Printf("Hello, World! %s %s\n", text, a.GetChecksum())
}
