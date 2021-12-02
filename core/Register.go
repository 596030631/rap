package core

import "net"

type Point struct {
	ip net.IPAddr
	port int
	name string
}

var sk = make(map[string] Point)

func Register()  {

}
