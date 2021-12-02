package main

import (
	"fmt"
	"github.com/rap/core"
	"log"
	"net"
)

func main() {
	laddr, e := net.ResolveUDPAddr("udp", ":8081")
	if e != nil {
		log.Fatal(e)
		return
	}

	listener, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return
	}

	defer func(listener *net.UDPConn) {
		err := listener.Close()
		if err != nil {
			log.Fatal(e)
		}
	}(listener)

	p := 0
	//end := 0
	//cache := make([]byte, 1024)
	b := make([]byte, 32)

	for {
		n, addr, e := listener.ReadFromUDP(b[p:])

		if e != nil || n <= 0 {
			continue
		}

		fmt.Printf("Read From %d %s\n", n, addr.String())

		p += n

		if b[0] == core.STX && p-4 >= int(b[1]&0xff) {
			p = 0
			fmt.Printf("Read From %d %s %s\n", n, addr.String(), core.Decoder(b[2:(2+b[1])]))
		}
		if e != nil {
			continue
		}

	}
}
