package main

import (
	"fmt"
	"github.com/rap/core"
	"net"
	"time"
)

func main() {
	laddr, e1 := net.ResolveUDPAddr("udp", ":0")
	laddr2, e1 := net.ResolveUDPAddr("udp", ":0")
	if e1 != nil {
		return
	}
	raddr, e2 := net.ResolveUDPAddr("udp", ":8081")
	if e2 != nil {
		return
	}
	udp, e3 := net.DialUDP("udp", laddr, raddr)
	udp2, e3 := net.DialUDP("udp", laddr2, raddr)
	if e3 != nil {
		return
	}

	var msg = "测试消息"
	b := core.Encoder(msg)
	stx, etx := core.Pick()
	stx[1] = byte(len(b))
	var e error

	go func() {
		var msg2 = "测试消息2"
		b2 := core.Encoder(msg2)

		for  {
			_, e = udp2.Write(stx)
			if e == nil {
				if e == nil {
					_, e := udp2.Write(b2)
					if e == nil {
						_, _ = udp2.Write(etx)
					}
				}
			}
			fmt.Printf("%s\n", string(b))
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		_, e = udp.Write(stx)
		if e == nil {
			if e == nil {
				_, e := udp.Write(b)
				if e == nil {
					_, _ = udp.Write(etx)
				}
			}
		}
		fmt.Printf("%s\n", string(b))
		time.Sleep(1 * time.Second)
	}
}
