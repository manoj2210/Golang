package main

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var bcServer chan []Block
var mutex = &sync.Mutex{}

func main() {
	bcServer = make(chan []Block)
	genesisBlock := Block{0, time.Now(), 0, "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :8080")
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlefunc(conn)

	}
}
