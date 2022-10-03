package main

import (
	"fmt"
	"sync"
)

func forRangeChannel(data []uint64) {
	var wg sync.WaitGroup
	channel := make(chan uint64, 100)
	wg.Add(1)
	go func() {
		for v := range channel {
			fmt.Printf("Value: %d Addr: %v\n", v, &v)
		}
		wg.Done()
	}()
	for i := range data {
		fmt.Printf("Data: %d Addr: %v\n", data[i], &data[i])
		channel <- data[i]
	}
	close(channel)
	wg.Wait()
}
