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
	"strings"
)

func scanport(protocol, host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout(protocol, host + ":" + strconv.Itoa(port), timeout)
	if err != nil {
		return false	// Cannot connect
	}
	conn.Close()
	return true		// Connection successful
}

func main() {
	
	if len(os.Args) < 2 {
		fmt.Printf("Watchdoge is a network port sniffer. This is the usage:\n")
		fmt.Printf("Scan an IP for open ports: watchdoge [IP address] [timeout]\n")
		fmt.Printf("Scan the subnet of an IP:  watchdoge [IP address]/24 [port] [timeout]\n")
		fmt.Printf("[timeout] is optional and the default is 100 milliseconds\n")
		os.Exit(1)
	}

	ip := os.Args[1]	// Get the IP from the command line
	scansub := false;
	if strings.HasSuffix(ip, "/24")	{
		scansub = true
		ip = ip[0:strings.LastIndex(ip, "/24")]
		if net.ParseIP(ip).To4() == nil {
			fmt.Printf("Unexpected IPv4 address: %s\n", ip)
			os.Exit(4)
		}
		ip = ip[0:strings.LastIndex(ip, ".")] + "."
	}

	timeout := 100 * time.Millisecond

	if !scansub {
		// Check if IP is valid
		if ip != "localhost" && net.ParseIP(ip).To4() == nil {
			fmt.Printf("'%s' is not a valid IPv4 address\n", ip)
			os.Exit(3)
		}

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
			// Connect TCP
			if scanport("tcp", ip, port, timeout) == true {
				fmt.Printf("TCP port %d open\n", port)
			}
		}
	} else {
		// Get the timeout to use
		if len(os.Args) > 3 {
			var err error
			delay, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Println(err)
				os.Exit(6)
			} else {
				timeout = time.Duration(delay) * time.Millisecond
			}
		}

		// Get the port to be used
		if len(os.Args) > 2 {
			var err error
			port, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				os.Exit(5)
			}

			// Scan the IPs to see who responds
			for host := 1; host < 255; host++ {
				// Connect TCP
				if scanport("tcp", ip + strconv.Itoa(host), port, timeout) == true {
					fmt.Printf("Answer from %s\n", ip + strconv.Itoa(host))
				}
			}
		} else {
			fmt.Printf("No port specified for scan.\n")
			os.Exit(7)
		}
	}

	fmt.Println("done")
}
