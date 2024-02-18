package main

import (
	"fmt"
	"math/rand"
)

func generateIntA(done chan struct{}) chan int {
	ch := make(chan int, 10)

	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
	}()

	return ch
}

func generateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)

	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
	}()

	return ch
}

func GenerateNum(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})

	go func() {
	Lable:
		for {
			select {
			case ch <- <-generateIntA(send):
			case ch <- <-generateIntB(send):
			case <-done:
				send <- struct{}{}
				break Lable
			}
		}

		close(ch)
	}()

	return ch
}

func main() {
	done := make(chan struct{})

	ch := GenerateNum(done)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	done <- struct{}{}
	fmt.Println("stop generation")
}
