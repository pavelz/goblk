package block_store

import (
	"encoding/binary"
	"io"
	"os"
)
const INDEX_SIZE uint = 0xffff

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

// get binary block via index at not binary offset
// oh no need sizes of the blocks too?
// should blocks have size at the beginning? this is two reads. no need.
// index - offset
// index + 1  = blocks size
// keep semantis in one place

func (b * BlockStore) BlockAt(at uint64) ([]byte){
	offset, size := b.index[at], b.index[at + 1]
	block := make([]byte, size, size)

	b.file.Seek(int64(offset),io.SeekStart)
	b.file.Read(block)
	return block
}
