package main

import (
	"gofmt/chain"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := md5.Sum([]byte("hex me"))
	text := hex.EncodeToString(hash[:])
	a := chain.NewChainer("123", "321", "312", 0)
	a.Hey()
	
	fmt.Printf("Hello, World! %s %s\n", text, a.Checksum)
}
