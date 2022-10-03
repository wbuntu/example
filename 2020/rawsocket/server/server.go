package main

import (
	"fmt"
	"net"
)

func main() {
	protocol := "icmp"
	netaddr, _ := net.ResolveIPAddr("ip4", "0.0.0.0")
	conn, err := net.ListenIP("ip4:"+protocol, netaddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		b := make([]byte, 128)
		rlen, addr, err := conn.ReadFrom(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Addr: %s %s Length: %d Data: %0x\n", addr.Network(), addr.String(), rlen, b[:rlen])
	}
}
