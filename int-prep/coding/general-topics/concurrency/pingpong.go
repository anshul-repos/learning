package main

import (
	"fmt"
	"sync"
)

// write a go code to print "ping" "pong" one at a time N times

func ping(pingCh, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= N; i++ {
		// step3: recieved singal from main, started execution of ping/pong
		<-pingCh
		fmt.Println("PING")
		// step4: send signal through channel pongCh that will be recieved in pong func
		pongCh <- struct{}{}
	}

}

func pong(pingCh, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= N; i++ {
		// step5: recieved singal from ping method, started execution of printing pong
		<-pongCh
		fmt.Println("PONG")
		if i < N { // avoid extra send at the end
			// step 6: send signal back through channel for ping method to execute again.
			pingCh <- struct{}{}
		}
	}
}

const N int = 5

// func main() {

// 	pingCh := make(chan struct{})
// 	pongCh := make(chan struct{})

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	//step1: started go routines ping and pong....
// 	go ping(pingCh, pongCh, &wg)
// 	go pong(pingCh, pongCh, &wg)

// 	// step2: send signal through channel pingCh...
// 	pingCh <- struct{}{}

// 	wg.Wait()
// }
