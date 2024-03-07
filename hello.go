package main

import ("fmt"
    "crypto/md5"
    "encoding/hex"
)
func main() {
    hash := md5.Sum([]byte("hex me"))
    text := hex.EncodeToString(hash[:])

    fmt.Printf("Hello, World! %s\n", text)
}
