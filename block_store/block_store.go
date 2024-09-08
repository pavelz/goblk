package block_store

import (
	"encoding/binary"
	"io"
	"os"
  "slices"
  "bytes"
)

const INDEX_SIZE uint = 0xffff

type BlockHeader struct {
  Index_blocks int64 // index_blocks * blocks_size = start of the blocks stored
}

// what is index here?
// hash to block number?

type BlockIndexEntry struct{
  Identifier [128]byte
  BlockNumber uint64
}

type BlockIndex struct{
  index []BlockIndexEntry
}

// index code
func (b *BlockIndex) WriteIndex(){
}

func (b *BlockIndex) Lookup(id string) (*BlockIndexEntry){
  idx := slices.IndexFunc(b.index, func(n BlockIndexEntry) (bool){
    return bytes.Equal([]byte(id)[:128], n.Identifier[:])
  })

  if idx < 0 {
    return nil
  }

  return &b.index[idx]
}

type BlockIfc interface {
	BlockAt(at uint64)
	SaveAt(at uint64, block []byte)
}

type BlockStore struct {
	index []uint64
	file *os.File
}

func Open(path string) (*BlockStore, error) {
	var store = BlockStore{index: make([]uint64, INDEX_SIZE, INDEX_SIZE)}

  file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644) 

	if err != nil {
		panic("could not open storage file: " + err.Error())
	} else {
		store.file = file // ngl this is shit
		binary.Read(file, binary.LittleEndian, store.index)
	}

	return &store, nil
}

// I guess at is a primary id index stuff
// is there a cache?

func (b * BlockStore) BlockAt(at uint64) ([]byte){
	at_s := at * 2
	offset, size := b.index[at_s], b.index[at_s + 1]
	block := make([]byte, size, size)

	b.file.Seek(int64(offset),io.SeekStart) //x
	b.file.Read(block)
	return block
}

func (b * BlockStore) SaveAt(at uint64,at_block []byte){
	at_s := at * 2
	offset, size := b.index[at_s], b.index[at_s + 1]
	block := make([]byte, size, size)
	at_size := len(at_block)

	if uint64(at_size) > size {
		end_offset, err := b.file.Seek(0, io.SeekEnd)

		if err != nil {
			panic("cannot seek:" + err.Error())
		}

		b.file.Write(at_block)
		b.index[at_s] = uint64(end_offset) - uint64(at_size)
		b.index[at_s + 1] = uint64(at_size)

	} else { 
		// replace block
		b.file.Seek(int64(offset), io.SeekStart)
		b.file.Write(block)
	}
}
