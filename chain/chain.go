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
    Checksum string // checksum of all blocks and this one
    Previous string // prev block
    From string // does it matter for chain? not at its core
    To string
    Amount int
    head *Chainer  `json:"-"`
    tail *Chainer  `json:"-"`
    prev *Chainer  `json:"-"`
    next *Chainer  `json:"-"`
}

func NewChainer(previous string, from string, to string, amount int) Chainer{
    c := Chainer{ Previous: previous, From: from, To: to, Amount: amount}
    c.Checksum = calcHexBlock(c)
    return c
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
func getTextAllBlocks(block Chainer) string{
    text_block := ""
    cursor := &block;
    for true {
        text_block += getNakedText(*cursor)
        cursor = cursor.next
        
        if cursor == nil {
            break
        }
    }

    return text_block
}

func (c Chainer) Hey(){
}


// this shall tigger mining? something like that.
func (c Chainer) Send(from string, to string, amount int) (error){
    // where from, where to.
    // from to addresses?
    // wallet?
    // From has to have enough Amount to send to To

    walletBalance,_ := c.getAmount(from)

    if walletBalance - amount > 0 {
        return errors.New("insufficient balance")
    }

    var  blk Chainer = Chainer{From: from,To: to, Amount: amount}

    // TODO prepend in datastructs
    c.tail.tail = &blk
    end, err := c.getBlock(to)

    if err != nil {
        return err
    }

    if end.tail != nil {
        return errors.New("end of the blockchain is not")
    }

    end.tail = &blk

    return nil
}

func (c Chainer) getBlock(address string) (*Chainer, error){

    var next *Chainer = &c;
    for {
        if next.To == address {
            return next, nil
        }
        if c.next == nil {
            break
        }
        next = c.next

    }
    return nil , errors.New("not found")
}

func (c Chainer) getAmount(address string) (int, error){
    blk, err := c.getBlock(address)
    if err != nil {
        return 0, err
    }

    return blk.Amount,nil
}

func calcHexBlock(block Chainer) string {
    acopy := block
    acopy.Checksum = ""
    acopy.head = nil
    acopy.tail = nil
    acopy.head = nil

    all_text := ""
    if block.prev != nil {
        all_text = getTextAllBlocks(*block.prev)
    }
    all_text_string := string(all_text)

    json_block, _ := json.Marshal(acopy)
    json_block_string := string(json_block)
    reply := all_text_string + json_block_string

    hash := md5.Sum([]byte(reply))
    text := hex.EncodeToString(hash[:])

    return  text
}

func getNakedText(block Chainer) string{
        acopy := block
        // acopy.Checksum = ""
        acopy.head = nil
        acopy.tail = nil
        acopy.next = nil
        acopy.prev = nil

    json_block, _ := json.Marshal(acopy)

    return string(json_block)
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
