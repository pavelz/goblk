package chain


import (
  "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Printf

func TestCalcHexBlock(t *testing.T){

  block := Block{Previous: "two", From: "me", To: "who"}

  assert.Equal(t, calcHexBlock(block), "dc600b3b9a29452c20e85e8d0e848ea9")
}

func TestGetAllTextBlocks(t *testing.T){
  // rewrite

  block1 := Block{Previous: "two", From: "me", To: "who"}
  block2 := Block{Previous: "two", From: "me", To: "who"}
  block3 := Block{Previous: "two", From: "me", To: "who"}

  assert.Equal(t, 198, len(getTextAllBlocks(block3)))
}

func TestGetNakedText(t *testing.T){

  block1 := Block{Previous: "two", From: "me", To: "who"}

  assert.Equal(t, 66, len(getNakedText(block1)))
}

func TestValidateChain(t *testing.T){
  t.Skip("to implement ValidateChain")
}

// thinking backwards what do i need to check for what is to be done writing stuff.
// mock fs methods and have them be called when running the code
func TestWriteChain(t *testing.T){
  t.Skip("to implement WriteChain")
}
