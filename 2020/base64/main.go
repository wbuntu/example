package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
)

func main() {
	a := uint64(math.Pow(10, 8))
	var num uint64
	for i := uint64(8); i <= uint64(64); i += 8 {
		data := make([]byte, i/8)
		rand.Read(data)
		base64Str := base64.URLEncoding.EncodeToString(data)
		base64RawStr := base64.RawURLEncoding.EncodeToString(data)
		num = 1<<i - 1
		if num > a*a {
			fmt.Printf("Bit:%-2d Padding:%-2d NoPading:%-2d Size:%d亿亿 %s\n", i, len(base64Str), len(base64RawStr), num/(a*a), base64Str)
		} else if num > a {
			fmt.Printf("Bit:%-2d Padding:%-2d NoPading:%-2d Size:%d亿 %s\n", i, len(base64Str), len(base64RawStr), num/a, base64Str)
		} else {
			fmt.Printf("Bit:%-2d Padding:%-2d NoPading:%-2d Size:%d %s\n", i, len(base64Str), len(base64RawStr), num, base64Str)
		}
	}
}
