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
