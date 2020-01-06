package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Message struct {
	BPM int
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
func handleWriteBlockchain(w http.ResponseWriter, r *http.Request) {
	var m Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	pb := Block{}
	if len(Blockchain) != 0 {
		pb = Blockchain[len(Blockchain)-1]
	}
	newBlock := generateBlock(m.BPM, pb)
	Blockchain = append(Blockchain, newBlock)
	if isBlockValid(newBlock, pb) {
		bytes, err := json.MarshalIndent(Blockchain, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, string(bytes))
	}
	return
}
