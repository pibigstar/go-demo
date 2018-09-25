package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

/**
区块
*/

type Block struct {
	Index         int64  `json:"index"`         // 区块编号
	Data          string `json:"data"`          // 区块保存的数据
	Hash          string `json:"hash"`          // 当前区块的Hash值
	Timestamp     int64  `json:"timestamp"`     // 时间戳
	prevBlockHash string `json:"prevBlockHash"` // 上一个区块的Hash值
}

/**
生成新的区块
*/
func GenerateNewBlock(prevBlock *Block, data string) *Block {
	newBlock := new(Block)

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.prevBlockHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

/**
计算 Hash 值
*/
func calculateHash(block *Block) string {
	blockHash := string(block.Index) + string(block.Timestamp) + block.Hash + block.Data + block.prevBlockHash
	blockBytes := sha256.Sum256([]byte(blockHash))
	return hex.EncodeToString(blockBytes[:])
}

/**
生成创始区块
*/
func GenerateGenesisBlock() *Block {
	block := new(Block)
	block.Index = -1
	block.Timestamp = time.Now().Unix()
	block.Hash = ""
	return GenerateNewBlock(block, "Genesis Block")
}
