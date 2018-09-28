package main

import (
	"encoding/json"
	"go-demo/blockchain/core"
	"io"
	"net/http"
)

type BlockChainResponse struct {
	BlockChain *core.BlockChain
	Total      int
}

var bcr *BlockChainResponse

func main() {
	blockchain := core.NewBlockChain()
	bcr = &BlockChainResponse{}
	bcr.BlockChain = blockchain
	bcr.Total = len(blockchain.Blocks)
	run()
}

func run() {
	http.HandleFunc("/block/get", GetBlockChain)
	http.HandleFunc("/block/write", WriteBlockChain)
	http.ListenAndServe(":9000", nil)
}

func WriteBlockChain(writer http.ResponseWriter, request *http.Request) {
	data := request.URL.Query().Get("data")
	bcr.BlockChain.SendData(data)
	bcr.Total = len(bcr.BlockChain.Blocks)
	GetBlockChain(writer, request)

}

func GetBlockChain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(bcr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
