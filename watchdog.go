// Copyright (c) 2017 Alexandre-Xavier Labonté-Lamoureux
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package main

import (
	"os"
	"strconv"
	"fmt"
	"net"
	"time"
)

func test(protocol, host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout(protocol, host + ":" + strconv.Itoa(port), timeout)
	if err != nil {
		return false	// Cannot connect
	}
	conn.Close()
	return true		// Connection successful
}

func main() {
	
	if len(os.Args) < 2 {
		fmt.Printf("Usage: watchdoge [IP address] [timeout]\n")
		fmt.Printf("[timeout] is optional and the default is 100 milliseconds\n")
		os.Exit(1)
	}

	ip := os.Args[1]	// Get the IP from the command line
	// Check if IP is valid
	if ip != "localhost" && net.ParseIP(ip).To4() == nil {
		fmt.Printf("'%s' is not a valid IPv4 address\n", ip)
		os.Exit(3)
	}

	timeout := 100 * time.Millisecond

	if len(os.Args) > 2 {
		var err error
		delay, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		} else {
			timeout = time.Duration(delay) * time.Millisecond
		}
	}
	
	for port := 0; port <= 65535; port++ {
		//Connect TCP
		if test("tcp", ip, port, timeout) == true {
			fmt.Printf("tcp port %d open\n", port)
		}
	}

	fmt.Println("done")
}
