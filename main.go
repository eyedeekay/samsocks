package main

import (
	"github.com/eyedeekay/sam3/helper"
	"github.com/eyedeekay/sam3/i2pkeys"
	"github.com/phayes/freeport"
	"github.com/txthinking/socks5"

	"flag"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func randomid() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	p := r.Intn(55534) + 10000
	return strconv.Itoa(p)
}

const version = "0.0.01"

func main() {
	// Create a SOCKS5 server
	addr := flag.String("socksaddr", "127.0.0.1", "Start the SOCKS5 proxy at this address(Almost always localhost. Can use a domain)")
	port := flag.Int("socksport", 7950, "Start the SOCKS5 proxy at this port.")
	name := flag.String("name", "sam-socks-"+randomid(), "Name to give the I2P client(The last part is always randomized)")
	username := flag.String("user", "", "Require a username to use the SOCKS5 Proxy.")
	password := flag.String("pass", "", "Require a password to use the SOCKS5 Proxy.")
	isolate := flag.Bool("isolate", true, "Enforce isolation. Works differently from torsocks.")
	version := flag.Bool("version", false, "Print the version and exit.")
	debug := flag.Bool("debug", false, "Show detailed information about the operation.")
	tcpTimeout := flag.Int("tcptimeout", 60000, "Set a default TCP Timeout(ms)")
	udpTimeout := flag.Int("udptimeout", 60000, "Set a default UDP Timeout(ms)")
	samaddress := flag.String("address", "127.0.0.1", "Specify I2P SAM address")
	samport := flag.Int("port", 7656, "Specify I2P SAM port")
	flag.Parse()
	ip := "127.0.0.1"
	//	shell := flag.Bool("shell", false, "spawn an I2P-only shell")
	if *version {
		log.Println("samsocks version:", version)
		os.Exit(0)
	}
	if *debug {
		log.Println("SAM client id:", *name)
	}

	var err error
	if *isolate {
		if b, _ := Check(*port); !b {
			*port, err = freeport.GetFreePort()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Port is occupied, the SOCKS5 proxy needs a new port to be isolated correctly")
			log.Println("SOCKS5 proxy will start on:", *addr, ":", *port)
		}
	}
	i2pkeys.FakePort = true

	primary, err := sam.I2PPrimarySession(*name, *samaddress+":"+strconv.Itoa(*samport), "")
	if err != nil {
		panic(err)
	}

	socks5.Dial = primary
	socks5.Resolver = primary

	server, err := socks5.NewClassicServer(*addr+":"+strconv.Itoa(*port), ip, *username, *password, *tcpTimeout, *udpTimeout)
	if err != nil {
		panic(err)
	}
	log.Println("Client Created SOCKS5 proxy at", *addr, ":", *port)
	// Create SOCKS5 proxy
	go func() {
		if err := server.ListenAndServe(nil); err != nil {
			panic(err)
		}
	}()

	for {

	}

}

// FROM: https://gist.github.com/montanaflynn/b59c058ce2adc18f31d6
func Check(port int) (status bool, err error) {

	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false, err
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true, nil

}
