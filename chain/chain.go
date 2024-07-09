package chain

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"goblk/fs"
	"os"
)


type Chainer struct {
    chain []Block
}

type Block struct {
    Checksum string // checksum of all blocks and this one
    Previous string // prev block
    From string // does it matter for chain? not at its core
    To string
    Amount int
}

func NewChain(previous string, from string, to string, amount int) (*Chainer){
    c := Block{ Previous: previous, From: from, To: to, Amount: amount}
    c.Checksum = calcHexBlock(c)
    chainer := Chainer{chain: []Block{c}}
    return &chainer 
}


// json chain? - at the momemnt
func LoadChain(path string) (*Chainer, error){
    chain_data,err := os.ReadFile(path)

    if err != nil {
        return nil, err
    }

    var chain Chainer
    err = json.Unmarshal(chain_data, &chain)

    if err != nil {
        return nil, errors.New("Error parsing blockchain data: " + err.Error())
    }

    return &chain, nil
}   

// get all block before this one 
// filter out fields except Checksum
func getTextAllBlocks(chain Chainer, before *Block = nil) string{
    var text_block = ""
    var err error
    for _, block :=  range chain.chain {

        // does not return current block too
        if block.Checksum == before.Checksum {
            break
        }

        text_block, err =  getNakedText(block)
        
        if err != nil {
            // lets panic some
            panic("Fatal error: " + err.Error())
        }
    }
    return text_block
}

func (c Chainer) Hey(){
}


// this shall tigger mining? something like that.
func (chain Chainer) Send(from string, to string, amount int) (*Block, error){
    // where from, where to.
    // from to addresses?
    // wallet?
    // From has to have enough Amount to send to To

    walletBalance,_ := chain.getAmount(from)

    if walletBalance - amount > 0 {
        return nil, errors.New("insufficient balance")
    }

    var blk = Block{From: from, To: to, Amount: amount}

    chain.chain = append(chain.chain, blk)
    return &blk, nil
}

func (c Chainer) getBlock(address string) (*Block, error){
    for _, item := range c.chain {
        if item.Checksum == address {
            return &item, nil
        }
    }
    return nil, errors.New("no address " + address + " found")
}

func (c Chainer) getAmount(address string) (int, error){
    blk, err := c.getBlock(address)
    if err != nil {
        return 0, err
    }

    return blk.Amount,nil
}

func calcHexBlock(block Block) string {
    acopy := block
    acopy.Checksum = ""

    all_text := getTextAllBlocks(*block.prev)
    all_text_string := string(all_text)

    json_block, _ := json.Marshal(acopy)
    json_block_string := string(json_block)
    reply := all_text_string + json_block_string

    hash := md5.Sum([]byte(reply))
    text := hex.EncodeToString(hash[:])

    return  text
}

func getNakedText(block Block) (string, error){
    json_block, err := json.Marshal(block)

    if err != nil {
        return "", err
    }
    return string(json_block), nil
}

// returns pointer to object that failed checksum check.
// todo: cumilative hash hex sums
// how to telescope in
func ValidateChain(chain *Chainer) (string, *Chainer) {
    next:= chain
    var prev *Chainer = nil
    for true {
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

// todo: write in protocol buffers?
func (c Chainer) WriteChain(path string) int {

    // write chain
    a := fs.Fs{}
    file := a.Open(path)
    file.Write([]byte(getNakedText(c)))
    file.Close()

    return 0
}
