package main

import (
	"github.com/eyedeekay/portcheck"
	"github.com/eyedeekay/sam3/helper"
	"github.com/eyedeekay/sam3/i2pkeys"
	"github.com/phayes/freeport"
	"github.com/txthinking/socks5"

	"flag"
	"log"
	"strconv"
)

func main() {
	// Create a SOCKS5 server
	addr := flag.String("socksaddr", "127.0.0.1", "Start the SOCKS5 proxy at this address(Can use a domain)")
	port := flag.String("socksport", 7950, "Start the SOCKS5 proxy at this port.")
	ip := flag.String("ipaddr", "127.0.0.1", "Listen on this IP address")
	username := flag.String("user", "", "Require a username to use the SOCKS5 Proxy.")
	password := flag.String("pass", "", "Require a password to use the SOCKS5 Proxy.")
	isolate := flag.Bool("isolate", true, "Enforce isolation.")
	tcpTimeout := flag.Int("tcptimeout", 60000, "Set a default TCP Timeout(ms)")
	udpTimeout := flag.Int("udptimeout", 60000, "Set a default UDP Timeout(ms)")
	samaddress := flag.String("address", "127.0.0.1", "Specify I2P SAM address")
	samport := flag.Int("port", "7656", "Specify I2P SAM port")
	//	shell := flag.Bool("shell", false, "spawn an I2P-only shell")

	var err error
	if *isolate {
		if portcheck.SCR(addr, port) {
			port, err = freeport.GetFreePort()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	i2pkeys.FakePort = true

	primary, err := sam.I2PPrimarySession("sam-socks", *samaddress+":"+strconv.Itoa(*samport), "socks5")
	if err != nil {
		panic(err)
	}

	socks5.Dial = primary
	socks5.Resolver = primary

	server, err := socks5.NewClassicServer(*addr+":"+*port, *ip, *username, *password, *tcpTimeout, *udpTimeout)
	if err != nil {
		panic(err)
	}
	log.Println("Client Created SOCKS5 proxy at", *addr)
	// Create SOCKS5 proxy
	go func() {
		if err := server.ListenAndServe(nil); err != nil {
			panic(err)
		}
	}()

	for {

	}

}
