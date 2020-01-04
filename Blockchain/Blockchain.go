package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	BPM       int
	Hash      string
	Prehash   string
}

var Blockchain []Block

func calculateHash(b Block) string {
	h := sha256.New()
	constr := strconv.Itoa(b.Index) + b.Timestamp.String() + strconv.Itoa(b.BPM) + b.Prehash
	h.Write([]byte(constr))
	return hex.EncodeToString([]byte(h.Sum(nil)))
}

func generateBlock(b int, pb Block) Block {
	block := Block{
		pb.Index + 1, time.Now(), b, "", calculateHash(pb),
	}
	block.Hash = calculateHash(block)
	return block
}

func isBlockValid(cb Block, pb Block) bool {
	if cb.Prehash != calculateHash(pb) {
		return false
	}
	if calculateHash(cb) != cb.Hash {
		return false
	}
	if pb.Index+1 != cb.Index {
		return false
	}
	return true
}
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
