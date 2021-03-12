package main

import (
	"github.com/eyedeekay/sam3/helper"
	"github.com/eyedeekay/sam3/i2pkeys"
	"github.com/txthinking/socks5"
	"log"
)

func main() {
	// Create a SOCKS5 server
	addr := "127.0.0.1:8888"
	ip := "127.0.0.1"
	username := ""
	password := ""
	tcpTimeout := 60000
	udpTimeout := 60000
	i2pkeys.FakePort = true

	primary, err := sam.I2PPrimarySession("sam-socks", "127.0.0.1:7656", "socks5")
	if err != nil {
		panic(err)
	}

	socks5.Dial = primary
	socks5.Resolver = primary

	server, err := socks5.NewClassicServer(addr, ip, username, password, tcpTimeout, udpTimeout)
	if err != nil {
		panic(err)
	}
	log.Println("Client Created")

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe(nil); err != nil {
		panic(err)
	}
}
