package main

import (
	"crypto/md5"
	"goblk/chain"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := md5.Sum([]byte("hex me"))
	text := hex.EncodeToString(hash[:])
	a := chain.NewChainer("123", "321", "312", 0)
	a.WriteChain("chain.json")
	
	fmt.Printf("Hello, World! %s %s\n", text, a.Checksum)
}
