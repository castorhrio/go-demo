package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func worker(ports, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}

		conn.Close()
		result <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go worker(ports, results, &wg)
	}

	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}

		close(ports)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("port %d open\n", port)
	}
}
