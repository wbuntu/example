package main

import "fmt"

func forKRangeMap(m map[string]uint64) {
	for k := range m {
		fmt.Printf("range: Key: %s %v\n", k, &k)
	}
}

func forKVRangeMap(m map[string]uint64) {
	for k, v := range m {
		fmt.Printf("range: Key: %s Value: %d Addr: %v %v\n", k, v, &k, &v)
	}
}
