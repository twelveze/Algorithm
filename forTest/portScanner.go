package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", port)
		fmt.Printf("start check port: %s\n", address)
		conn, err := net.DialTimeout("tcp", address, time.Second*5)
		if err != nil {
			results <- 0
			continue
		}
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
		}
		results <- port
	}
}

func main() {
	ports := make(chan int, 30000)
	results := make(chan int, 1000)
	scanEndPort := 65535
	core := 2
	var openPorts []int
	begin := time.Now().Unix()
	// push port to channel
	go func() {
		for i := 1; i <= scanEndPort; i++ {
			ports <- i
		}
		close(ports)
	}()

	for i := 0; i < core; i++ {
		go worker(ports, results)
	}

	for i := 1; i <= scanEndPort; i++ {
		res := <-results
		if res != 0 {
			openPorts = append(openPorts, res)
		}
	}
	//close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("127.0.0.1's opened port is: %d\n", port)
	}
	end := time.Now().Unix()
	processTime := end - begin
	fmt.Printf("%d core process time is: %d", core, processTime)
}
