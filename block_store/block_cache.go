package block_store


type BlockCached struct{
	cache [][]byte //list of byte blocks?
	blk *BlockStore
}

// index size what should it be
func (b * BlockCached) BlockAt(at uint64) ([]byte){
	if b.cache[at] != nil {
		return b.cache[at]
	} else {
		blk := b.blk.BlockAt(at)
		b.cache[at] = blk
		return blk
	}
}

func (b * BlockCached) SaveAt(at uint64, blk []byte) {
	b.cache[at] = blk
	b.blk.SaveAt(at, blk)
}
