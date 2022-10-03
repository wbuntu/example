package main

import (
	"fmt"
	"sync"
)

func forRange(data []uint64) {
	fmt.Println("for range { body }")
	for range data {
		fmt.Printf("range\n")
	}
}

func forIRange(data []uint64) {
	fmt.Println("for i range { body }")
	for i := range data {
        v := data[i]
		fmt.Printf("Index: %d :Value: %d :Addr: %v %v\n", i, data[i], &v, &data[i])
	}
}

func forIVRnage(data []uint64) {
	fmt.Println("for i,v range { body }")
	for i, v := range data {
		fmt.Printf("Index: %d :Value: %d :Addr: %v %v\n", i, v, &i, &v)
	}
}

func forIVRangeAsyncCapture(data []uint64) {
	fmt.Println("for i,v range { async capture body }")
	var wg sync.WaitGroup
	for i, v := range data {
		wg.Add(1)
		go func() {
			fmt.Printf("Capture Index: %d :Value: %d :Addr: %v %v\n", i, v, &i, &v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func forIVRangeAsyncCopy(data []uint64) {
	fmt.Println("for i,v range { async copy body }")
	var wg sync.WaitGroup
	for i, v := range data {
		wg.Add(1)
		go func(i int, v uint64) {
			fmt.Printf("Copy Index: %d :Value: %d :Addr: %v %v\n", i, v, &i, &v)
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}
