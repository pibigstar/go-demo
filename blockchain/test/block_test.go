package test

import (
	. "go-demo/blockchain/core"
	"testing"
  "strconv"
)

func Test_Block(t *testing.T) {

	t.Log(GenerateGenesisBlock())
}

func Test_BlockChain(t *testing.T) {

	chain := NewBlockChain()
	for i := 0; i < 10; i++ {
		chain.SendData("block:"+strconv.Itoa(i))
	}

	for _, value := range chain.Blocks {
		t.Logf("%+v", value)
	}

}
