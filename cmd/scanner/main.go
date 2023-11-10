package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

// scanme.nmap.org
// sequential ersion

var (
	host = "localhost"
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to scan (scanme.nmap.org)")
}

func main() {
	flag.Parse()

	validPorts := []string{}
	dialer := net.Dialer{Timeout: time.Second * 1}

	for i := 1; i < 1024; i++ {
		target := fmt.Sprint(host, ":", i)
		log.Println(target)
		conn, err := dialer.Dial("tcp", target)
		if err == nil {
			fmt.Println("Connection successful")
			validPorts = append(validPorts, target)
			conn.Close()
		} else {
			fmt.Println(err)
		}
	}

	log.Println("============== valid ==============")
	for _, item := range validPorts {
		log.Println(item)
	}
}
