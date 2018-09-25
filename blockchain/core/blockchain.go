package core

import "github.com/smallnest/rpcx/log"

type BlockChain struct {
	Blocks []*Block
}

/**
创建一个新的区块链
 */
func NewBlockChain() *BlockChain {
	block := GenerateGenesisBlock()
	bc := new(BlockChain)
	bc.AppendBlock(block)
	return bc
}

/**
根据data生成一个新的区块并加入到区块链中
 */
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	nextBlock := GenerateNewBlock(preBlock, data)
	bc.AppendBlock(nextBlock)
}

/**
将区块加入到区块链中
 */
func (bc *BlockChain) AppendBlock(block *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, block)
	} else {
		if bc.isValid(block) {
			bc.Blocks = append(bc.Blocks, block)
		} else {
			log.Errorf("%+v is invalid", block)
		}
	}
}

/**
验证区块是否有效
 */
func (bc *BlockChain) isValid(block *Block) bool {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	if prevBlock.Index != block.Index-1 {
		return false
	}

	if prevBlock.Hash != block.PrevBlockHash {
		return false
	}

	return true
}
