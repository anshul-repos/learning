package main

import (
	"fmt"
	"sync"
)

// write two go routines to print eeven and odd numbers , also channels

func printEvenNums(numCh chan int, wg *sync.WaitGroup, signal chan struct{}) {
	defer wg.Done()

	for num := range numCh {
		<-signal
		if num%2 == 0 {
			fmt.Printf("\n Even : %+v", num)
		}
		signal <- struct{}{}
	}
}

func printOddNums(numCh chan int, wg *sync.WaitGroup, signal chan struct{}) {
	defer wg.Done()

	for num := range numCh {
		<-signal
		if num%2 != 0 {
			fmt.Printf("\n Odd : %+v", num)
		}
		signal <- struct{}{}
	}
}

func main() {

	n := 10
	numCh := make(chan int)
	signal := make(chan struct{}, 1)

	var wg sync.WaitGroup

	wg.Add(2)

	go printEvenNums(numCh, &wg, signal)
	go printOddNums(numCh, &wg, signal)

	signal <- struct{}{}

	for i := 1; i <= n; i++ {
		numCh <- i
	}

	close(numCh)
	wg.Wait()

}
