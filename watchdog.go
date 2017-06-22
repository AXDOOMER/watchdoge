package main

import (
	"os"
	"strconv"
	"fmt"
	"net"
	"time"
)

func test(protocol, host string, port int) bool {
	conn, err := net.DialTimeout(protocol, host + ":" + strconv.Itoa(port), 1 * time.Second)
	if err != nil {
		return false	// Cannot connect
	}
	/*defer*/ conn.Close()
	return true	// Connection successful
}

func listen(protocol, host string, port int) bool {
	ln, err := net.Listen(protocol, host + ":" + strconv.Itoa(port))
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s", port, err)
		return false
	}
	ln.Close()
	return true
}

func main() {
	
	if len(os.Args) < 2 {
		fmt.Printf("No IP address specified, please pass one as a paremeter on the CLI.\n")
		os.Exit(1)
	}

	ip := os.Args[1]	// Get the IP from the command line
	
	for port := 0; port <= 65535; port++ {
		
		//Connect TCP
		if test("tcp", ip, port) == true {
			
			fmt.Printf("tcp port %d open\n", port)
		}
		
		//Connect UDP
		/*if test("udp", ip, port) == true {
			
			fmt.Println("udp port " + strconv.Itoa(port) + " open")
		}*/
	}

	fmt.Println("done")
}


