package main

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"fmt"
)

func main() {
	for i := 1; i <= 1024; i++ {
		data := make([]byte, i)
		rand.Read(data)
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		zw.Write(data)
		zw.Close()
		fmt.Printf("%d %d\n", i, buf.Len())
	}
}
