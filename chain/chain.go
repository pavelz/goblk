package chain

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"goblk/fs"
	"os"

	"goblk/proto/gopkg/block"

	//"google.golang.org/protobuf/proto"
)


type Chainer struct {
    chain []Block
    lastBlock int
}

type Entry struct {
  From string
  To string
  Amount int64
}

type Block struct {
    Checksum string // checksum of all blocks and this one
    Previous string // prev block
    From string // does it matter for chain? not at its core
    To string
    Amount int
    Entries []Entry
}

// i have some reservations about this
func (b *Block) toProtobuf() (block.Block){

  blk := MapEntriesBlock(b.Entries, func(e *Entry)(*block.Entry){
    return &block.Entry{From: e.From, To: e.To, Amount: int64(e.Amount)} 
  })

  return block.Block{
    Entries: blk,
  }

}

func NewChain() (*Chainer){
    chainer := Chainer{chain: make([]Block, 1024) }
    return &chainer 
}

func (c *Chainer) GetChecksum() (string){
  return c.chain[c.lastBlock].Checksum
}

func (c *Chainer) AddBlock(b *Block) (error) {
  // find block
  if(len(c.chain) > c.lastBlock) {
    return errors.New("Out of Mememory")
  }

  // add block
  c.lastBlock ++
  c.chain[c.lastBlock] = *b

  return nil
}

func (b *Block) AddEntry(from string, to string, amount string){
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
func getTextAllBlocks(chain *Chainer) (string){
    return getTextAllBlocksBefore(chain, nil)
}

func getTextAllBlocksBefore(chain *Chainer, before *Block) (string){
    var text_block = ""
    var err error
    for _, block :=  range chain.chain {
        if before != nil && block.Checksum == before.Checksum {
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

func calcHexBlock(chain *Chainer, block *Block) string {
    acopy := block
    acopy.Checksum = ""

    all_text := getTextAllBlocksBefore(chain, block)

    hash := md5.Sum([]byte(all_text))
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
func ValidateChain(chain *Chainer) (*Block, error) {
    for _, block := range chain.chain {

        // calc hex for the block and all previous blocks
        text := getTextAllBlocksBefore(chain, &block)
        // exclude checksum from current calc. avoids off by one errors for first blocks.
        // when hash is calculated for current block originally it should not have a check sum, right?
        // history blocks can have their checksums in hash ingestion dataset.
        var save_hash = block.Checksum 
        block.Checksum = ""

        extra,_ := getNakedText(block)
        hash := md5.Sum([]byte(text + extra))
        hex := hex.EncodeToString(hash[:])

        if save_hash != hex {
            block.Checksum = save_hash
            return &block, errors.New("hash checksum mismatch")
        }
    }
    return nil, nil
}

func (c Chainer) WriteChainProtobuf(path string){
  for _, block :=  range(c.chain) {
    println("whoop")
    fmt.Printf("o: %s\n", block.From)
  }
}


// todo: write in protocol buffers?
func (c Chainer) WriteChain(path string) int {

    // write chain
    a := fs.Fs{}
    file := a.Open(path)
    file.Write([]byte(getTextAllBlocks(&c)))
    file.Close()


    return 0
}

func MapEntriesChain(data []*block.Entry, f func(*block.Entry) Entry) []Entry {
    var result []Entry
    for _, e := range data {
        result = append(result, f(e))
    }
    return result
}

func MapEntriesBlock(data []Entry, f func(*Entry) *block.Entry) []*block.Entry {
    var result []*block.Entry
    for _, e := range data {
        result = append(result, f(&e))
    }
    return result
}
