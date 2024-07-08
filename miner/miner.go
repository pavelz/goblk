package miner
import (
	"goblk/chain"
)

type Miner interface {
	Mine(mine_path string) (error)
	LastTransaction(mine_path string) (int)
}

type SolutionCheck interface{
	Check(problem interface{}) (bool)
}

type HashCheck struct{
}

func (h HashCheck) Check(problem interface{}){
}

type Mine struct {
}

func (m Mine) Mine(mine_path string) (error) {
	return nil
}


