package chain
import (
    "encoding/json"
    "crypto/md5"
    "encoding/hex"
)

type Chainer struct {
    Checksum string //checksum of all blocks and this one
    Previous string //prev block
    From string // does it matter for chain? not at its core
    To string
    Amount int
    head *Chainer  `json:"-"`
    tail *Chainer  `json:"-"`
    next *Chainer  `json:"-"`
}

func NewChainer(previous string, from string, to string, amount int) Chainer{
    c := Chainer{ Previous: previous, From: from, To: to, Amount: amount}
    json_block, _ := json.Marshal(c)
    hash := md5.Sum([]byte(json_block))
    text := hex.EncodeToString(hash[:])
    c.Checksum = text
    return c
}

func calcHexBlock(block Chainer) string {
    json_block, _ := json.Marshal(block)
    hash := md5.Sum([]byte(json_block))
    text := hex.EncodeToString(hash[:])
    return  text
}

func ValidateChain(chain *Chainer) int{
    next := chain
    for true{
        calcHexBlock(*next)

        next = next.next
        break
    }
    return 0
}

func writeChain(path string) int {
    return 0
}
