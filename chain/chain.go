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
        acopy := block
        acopy.Checksum = ""
        acopy.head = nil
        acopy.tail = nil
        acopy.head = nil

    json_block, _ := json.Marshal(acopy)
    hash := md5.Sum([]byte(json_block))
    text := hex.EncodeToString(hash[:])
    return  text
}

func getNakedText(block Chainer) string{
        acopy := block
        acopy.Checksum = ""
        acopy.head = nil
        acopy.tail = nil
        acopy.head = nil

    json_block, _ := json.Marshal(acopy)

    return string(json_block)
}

func getTextAllBlocks(block Chainer) string{
    text_block := ""
    cursor := &block;
    for true {
        text_block += getNakedText(*cursor)
    }
    return text_block
}

// returns pointer to object that failed checksum check.
// todo: cumilative hash hex sums
// how to telescope in
func ValidateChain(chain *Chainer) (string, *Chainer) {
    next:= chain
    var prev *Chainer = nil
    for true{
        hex_check := calcHexBlock(*next)
        if hex_check != next.Checksum {
            return hex_check, next
        }

        // calculate cumilitive checksum
        // shallow check - checksums of checksums
        if prev != nil {
            
        }

        next = next.next
        if next.next == nil {
            break
        }
    }
    return "", nil
}

func writeChain(path string) int {
    return 0
}
