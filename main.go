package main

import (
	"github.com/eyedeekay/goSam"
	"github.com/getlantern/go-socks5"
	"log"
)

func main() {
	// Create a SOCKS5 server
	sam, err := goSam.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	log.Println("Client Created")

	// create a transport that uses SAM to dial TCP Connections
	conf := &socks5.Config{
		Dial:     sam.DialContext,
		Resolver: sam,
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "127.0.0.1:8000"); err != nil {
		panic(err)
	}
}
