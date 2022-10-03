package main

import "fmt"

func main() {
	var data []uint64
	for i := uint64(1); i <= uint64(100); i++ {
		data = append(data, i)
	}
	// slice
	forRange(data)
	forIRange(data)
	forIVRnage(data)
	forIVRangeAsyncCapture(data)
	forIVRangeAsyncCopy(data)
	// string
	var s string
	for i := range data {
		s += fmt.Sprintf("%d", data[i]%10)
	}
	forRangeString(s)
	// map
	m := make(map[string]uint64, len(data))
	for i := range data {
		key := fmt.Sprintf("%d", i)
		m[key] = data[i]
	}
	forKRangeMap(m)
	forKVRangeMap(m)
	// channel
	forRangeChannel(data)
}
