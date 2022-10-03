package main

/*
int calculate() {
	int mat[10000000][10000000];
	return mat[0][0];
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.calculate())
	fmt.Println("akari")
}
