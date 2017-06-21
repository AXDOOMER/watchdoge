package main

import (
	"os"
	"fmt"
	"net"
	"time"
)

// nmap -p 36367 localhost
// nmap -p {00000..65535} localhost

func test(protocol, host string, port int) bool {
				// net.ListenTCP ?						// 100 * time.Millisecond
	conn, err := net.DialTimeout(protocol, host + ":" + port, 1 * time.Second)
	if err != nil {
		return false;	// Cannot connect
	}
	/*defer*/ conn.Close()
	return true;	// Connection successful
}

func listen(protocol, host string, port int) bool {
				// net.ListenTCP ?
	ln, err := net.Listen(protocol, host + ":" + port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
	}
	ln.Close()
}

func main() {
	
	ip := os.Args[1]	// Get the IP from the command line
	
	for port := 0; port <= 65535; port++ {
		
		//Connect TCP
		if test("tcp", ip, port) == true {
			
			fmt.Printf("tcp port %d open\n", port)
		}
		
		//Connect udp
		if test("udp", ip, port) == true {
			
			fmt.Println("udp port " + port + " open")
		}
	
		fmt.Println("done")

	}
}


