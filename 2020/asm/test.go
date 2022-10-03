package main

func test1() {
	var x [1000]byte
	x[1] = 1
}

func test2() {
	var x [17000]byte
	x[1] = 1
}

func main() {
	test1()
	test2()
}
