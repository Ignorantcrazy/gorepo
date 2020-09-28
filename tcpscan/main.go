package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

const taskcount = 1024

func worker(ports chan int, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("39.98.144.216:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			errp := p - taskcount
			result <- errp
			continue
		}
		conn.Close()
		result <- p
	}
}

func main() {
	start := time.Now()
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var closeports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < taskcount; i++ {
			ports <- i
		}
	}()

	for i := 1; i < taskcount; i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
		} else {
			closeports = append(closeports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)
	sort.Ints(closeports)

	for _, port := range closeports {
		fmt.Printf("%d close\n", port+taskcount)
	}

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)
}
