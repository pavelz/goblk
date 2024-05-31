package chain


import (
  "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Printf

func TestCalcHexBlock(t *testing.T){
  block := Chainer{Previous: "two", From: "me", To: "who"}

  assert.Equal(t, calcHexBlock(block), "dc600b3b9a29452c20e85e8d0e848ea9")
}

func TestGetAllTextBlocks(t *testing.T){
  block1 := Chainer{Previous: "two", From: "me", To: "who"}
  block2 := Chainer{Previous: "two", From: "me", To: "who",next: &block1}
  block3 := Chainer{Previous: "two", From: "me", To: "who", next: &block2}

  //Printf("\nblocks: %d\n", len(getTextAllBlocks(block3)))
  assert.Equal(t, 198, len(getTextAllBlocks(block3)))
}


func TestGetNakedText(t *testing.T){
  block1 := Chainer{Previous: "two", From: "me", To: "who"}
  //fmt.Printf("%d\n", len(getNakedText(block1)))
  assert.Equal(t, 66, len(getNakedText(block1)))
}
//func TestSomething(t *testing.T) {

  //// assert equality
  //assert.Equal(t, 123, 123, "they should be equal")

  //// assert inequality
  //assert.NotEqual(t, 123, 456, "they should not be equal")

  //// assert for nil (good for errors)
  //assert.Nil(t, object)

  //// assert for not nil (good when you expect something)
 //if assert.NotNil(t, object) {

    //// now we know that object isn't nil, we are safe to make
    //// further assertions without causing any errors
    //assert.Equal(t, "Something", object.Value)
  //}
//}
