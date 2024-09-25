package chain


import (
  "fmt"
	"testing"

	"github.com/stretchr/testify/assert"

)

var _ = fmt.Printf

func TestCalcHexBlock(t *testing.T){

  c := NewChain()
  block := Block{Previous: "two", From: "me", To: "who"}
  c.AddBlock(&block)
  
  assert.Equal(t, calcHexBlock(c, &block), "d41d8cd98f00b204e9800998ecf8427e")
}

func TestGetAllTextBlocks(t *testing.T){
  // rewrite
  c := NewChain()

  block1 := Block{Previous: "two", From: "me", To: "who"}
  c.AddBlock(&block1)

  block2 := Block{Previous: "two", From: "me", To: "who"}
  c.AddBlock(&block2)

  block3 := Block{Previous: "two", From: "me", To: "who"}
  c.AddBlock(&block3)

  assert.Equal(t, 58, len(getTextAllBlocks(c)))
}

func TestGetNakedText(t *testing.T){

  block1 := Block{Previous: "two", From: "me", To: "who"}
  text, err := getNakedText(block1)
  assert.Equal(t, nil, err)
  assert.Equal(t, 66, len(text))
}

func TestValidateChain(t *testing.T){
  t.Skip("to implement ValidateChain")
}

// thinking backwards what do i need to check for what is to be done writing stuff.
// mock fs methods and have them be called when running the code
func TestWriteChain(t *testing.T){
  t.Skip("to implement WriteChain")
}
