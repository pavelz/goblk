package block_store

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
  "unsafe"
	"slices"

	"encoding/gob"
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

// Read index into struct
func ReadIndex(path string) (BlockIndex) {

  file, err := os.Open(path)

  if err != nil {
    os.Exit(-1)
  }

  var index_size_byte [8]byte
  file.Seek(8, 2)
  file.Read(index_size_byte[:])

  reader := bytes.NewBuffer(index_size_byte[:])

  var index_size int64
  gob.NewDecoder(reader).Decode(&index_size)
  blob := make([]byte, index_size)
  file.Read(blob)

  file.Seek(8 + index_size,2)
  index_data := make([]byte, index_size)

  file.Read(index_data)

  if err != nil {
    println(err)
    os.Exit(-1)
  }

  index_blocks := index_size / int64(unsafe.Sizeof(BlockIndexEntry{}))

  blocks := make([]BlockIndexEntry, index_blocks)
  err = gob.NewDecoder(bytes.NewReader(index_data)).Decode(&blocks)
  return BlockIndex{index: blocks}
}

// Write index to the end of the file
func (b *BlockIndex) WriteIndex(path string){
  file, err := os.Open(path)
  if err != nil {
    println(err)
    os.Exit(-1)
  }

  file.Seek(0, 2)

  buf := new(bytes.Buffer)
  gob.NewEncoder(buf).Encode(*b)

  file.Write(buf.Bytes())
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
