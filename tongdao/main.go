package main

import "runtime"

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		//不关闭通道会出现死锁
		close(ci)

		//c起到同步的作用
		c <- struct{}{}
	}()

	println("number goroutine=", runtime.NumGoroutine())

	//读通道c，通过通道进行同步等待
	<-c

	println("number goroutine=", runtime.NumGoroutine())

	for v := range ci {
		println(v)
	}

	println("number goroutine=", runtime.NumGoroutine())
}
