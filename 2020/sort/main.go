package main

import (
	"fmt"
	"sort"
)

// IntSlice obj
type IntSlice []int

// Len method
func (s IntSlice) Len() int {
	return len(s)
}

// Less method
func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap method
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := IntSlice{3843, 6695, 44722, 31230, 2483, 37300, 19869, 13995, 20156, 49553, 58055, 32813, 48747, 63995, 7986, 8513, 39982, 25459, 5259, 13594}
	fmt.Println("orignal: ", s)
	var s1, s2, s3, s4, s5, s6 IntSlice
	s1 = append(s1, s[:]...)
	s2 = append(s2, s[:]...)
	s3 = append(s3, s[:]...)
	s4 = append(s4, s[:]...)
	s5 = append(s5, s[:]...)
	s6 = append(s6, s[:]...)
	sort.Sort(s1)
	fmt.Println("quickSort: ", s1)
	sort.Stable(s2)
	fmt.Println("stableSort: ", s2)
	bubbleSort(s3)
	fmt.Println("bubbleSort: ", s3)
	selectionSort(s4)
	fmt.Println("selectionSort: ", s4)
	insertionSort(s5)
	fmt.Println("insertionSort: ", s5)
	insertionSort(s6)
	fmt.Println("shellSort: ", s6)
}

func bubbleSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func selectionSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

func insertionSort(data sort.Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func shellSort(data sort.Interface) {
	n := data.Len()
	gap := n / 2
	for i := gap; i < n; i++ {
		if data.Less(i, i-gap) {
			data.Swap(i, i-gap)
		}
	}
}
