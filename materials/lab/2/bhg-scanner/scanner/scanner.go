// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: Call PortScanner to get back (the number of open ports, the number of closed ports) at scanme.nmap.org.

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

type labeledPort struct {
	port int
	open bool
}

func worker(ports chan int, results chan labeledPort) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, time.Millisecond*500)
		if err != nil {
			results <- labeledPort{p, false}
			continue
		}
		conn.Close()
		results <- labeledPort{p, true}
	}
}

func PortScanner() (int, int) {
	var openports []int // notice the capitalization here. access limited!
	var closedports []int

	ports := make(chan int, 10)
	results := make(chan labeledPort)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		labeledPort := <-results
		if labeledPort.open {
			openports = append(openports, labeledPort.port)
		} else {
			closedports = append(closedports, labeledPort.port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)

	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d, open\n", port)
	}
	for _, port := range closedports {
		fmt.Printf("%d, closed\n", port)
	}

	return len(openports), len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed);
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
