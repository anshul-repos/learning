package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

const (
	A = -1
	B = -2
)

func goroutine1() error {

	C := A + B
	if C > 0 {
		errMsg := "Goroutine1 Error"
		return fmt.Errorf("%s", errMsg)
	}

	return nil
}

func goroutine2() error {
	C := A - B
	if C > 0 {
		errMsg := "Goroutine2 Error"
		return fmt.Errorf("%s", errMsg)
	}

	return nil
}

func main() {

	var g errgroup.Group

	g.Go(func() error {
		return goroutine1()
	})

	g.Go(func() error {
		return goroutine2()
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Error Captured:%v", err)
		return
	}

	fmt.Println("All goroutines finished successfully")

}
