package main

import "fmt"

func forRangeString(data string) {
	fmt.Println("for i,v range { body }")
	for i, v := range data {
		fmt.Printf("Index: %d :Value: %v :Addr: %v %v\n", i, v, &i, &v)
	}
}
