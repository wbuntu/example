package main

func main() {
	rxChan := make(chan int)
	endChan := make(chan int)
	go rxFunc(rxChan, endChan)
	for i := 0; i < 10; i++ {
		rxChan <- i
	}
	close(rxChan)
	<-endChan
}

func rxFunc(rxChan chan int, endChan chan int) {
	for range rxChan {
	}
	endChan <- 123456789
}
