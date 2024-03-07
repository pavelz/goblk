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
    c.Checksum = calcHexBlock(c)
    return c
}

func calcHexBlock(block Chainer) string {
    json_block, _ := json.Marshal(block)
    hash := md5.Sum([]byte(json_block))
    text := hex.EncodeToString(hash[:])
    return  text
}


// returns pointer to object that failed checksum check.
// todo: cumilative hash hex sums
func ValidateChain(chain *Chainer) (string, *Chainer) {
    next := chain
    for true{
        acopy := *next
        acopy.Checksum = ""
        acopy.head = nil
        acopy.tail = nil
        acopy.head = nil
        hex_check := calcHexBlock(acopy)
        if hex_check != next.Checksum {
            return hex_check, next
        }
        next = next.next
        break
    }
    return "", nil
}

func writeChain(path string) int {
    return 0
}
