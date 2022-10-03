package main

import (
	"errors"
	"fmt"
)

func main() {
	arrayData := [][]int{
		[]int{4, 5, 6, 7, 5, 6, 8}, []int{5, 6},
		[]int{4, 5, 7, 5, 8}, []int{5, 6},
		[]int{4, 5, 6, 7, 5, 6, 8}, []int{6},
		[]int{4, 5, 6, 7, 5, 6, 8}, []int{4, 5, 6},
	}
	for i := 0; i < len(arrayData); i += 2 {
		fmt.Printf("Data: %v %v  Result: %d\n", arrayData[i], arrayData[i+1], subArray(arrayData[i], arrayData[i+1]))
	}
	fmt.Println("******************************************")
	strData := []string{
		"123456789123456789123456789", "923456789123456789123456789",
		"1", "2",
		"A", "123456789123456789123456789",
		"22", "99",
	}
	for i := 0; i < len(strData); i += 2 {
		if data, err := addStr(strData[i], strData[i+1]); err != nil {
			fmt.Printf("Data: %s %s  Result: %s\n", strData[i], strData[i+1], err)
		} else {
			fmt.Printf("Data: %s %s  Result: %s\n", strData[i], strData[i+1], data)
		}
	}
}

func subArray(parent, child []int) int {
	index := -1
LOOP:
	for i := range parent {
		if parent[i] == child[0] {
			for j := range child {
				if child[j] != parent[i+j] {
					continue LOOP
				}
			}
			index = i
		}
	}
	return index
}

func addStr(a, b string) (string, error) {
	resultLen := 0
	aLen := len(a)
	bLen := len(b)
	if aLen > bLen {
		resultLen = aLen + 1
	} else {
		resultLen = bLen + 1
	}
	aContainer, bContainer, resultContainer := make([]int, resultLen), make([]int, resultLen), make([]int, resultLen)

	for i := 0; i < aLen; i++ {
		v := a[aLen-i-1]
		if v < '0' || v > '9' {
			return "", errors.New("Error")
		}
		aContainer[i] = int(v - '0')
	}

	for i := 0; i < bLen; i++ {
		v := b[bLen-i-1]
		if v < '0' || v > '9' {
			return "", errors.New("Error")
		}
		bContainer[i] = int(v - '0')
	}

	for i := 0; i < resultLen; i++ {
		resultContainer[i] += aContainer[i] + bContainer[i]
		if resultContainer[i] > 9 {
			resultContainer[i] = resultContainer[i] % 10
			resultContainer[i+1] = 1
		}
	}

	var resultStr string
	if resultContainer[resultLen-1] > 0 {
		resultStr += fmt.Sprintf("%d", resultContainer[resultLen-1])
	}
	for i := resultLen - 2; i >= 0; i-- {
		resultStr += fmt.Sprintf("%d", resultContainer[i])
	}
	return resultStr, nil
}
