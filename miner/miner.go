package miner

type Miner interface {
	Mine(mine_path string) (error)
	LastTransaction(mine_path string) (int)
}

type Mine struct {
}

func (m Mine) Mine(mine_path string) (error) {
	return nil
}
