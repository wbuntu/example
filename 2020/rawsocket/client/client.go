package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	raddr, err := net.ResolveIPAddr("ip", "67.230.183.176")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialIP("ip:icmp", nil, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// Make a new ICMP message
	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}
	ping, err := m.Marshal(nil)
	if err != nil {
		panic(err)
	}
	for {
		if _, err := conn.Write(ping); err != nil {
			panic(err)
		}
		b := make([]byte, 1024)
		rlen, addr, err := conn.ReadFrom(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Addr: %s Length: %d Data: %0x\n", addr, rlen, b[:rlen])
		time.Sleep(time.Second)
	}
}
