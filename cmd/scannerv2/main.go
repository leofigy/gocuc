package main

/*
Note : this program is not okay :) we need to fix it in the workshop :)
./scannerv2 > out 2>&1
*/

import (
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

var (
	host = "localhost"
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host to scan (scanme.nmap.org)")
}

func main() {
	flag.Parse()
	log.Println("V2 version")

	validPorts := []string{}

	dialer := net.Dialer{Timeout: time.Second * 1}
	waiter := sync.WaitGroup{}

	//in := make(chan string)
	// 22 - 80
	for i := 1; i < 40000; i++ {
		// add an element to wait
		waiter.Add(1)
		target := fmt.Sprint(host, ":", i)

		go func() {
			defer waiter.Done()
			conn, err := dialer.Dial("tcp", target)
			if err == nil {
				fmt.Println(runtime.NumGoroutine(), "Connection successful")
				validPorts = append(validPorts, target)
				//in <- target
				conn.Close()
			} else {
				log.Println(runtime.NumGoroutine(), err)
				validPorts = append(validPorts, target)
			}
		}()

	}

	waiter.Wait()

	/*

		done := make(chan bool)
		go func() {
			// this is the cleaner
			waiter.Wait()
			done <- true
		}()
		   Analyzer:
		   	for {
		   		select {
		   		case validPort := <-in:
		   			validPorts = append(validPorts, validPort)
		   		case <-time.After(time.Second):
		   			log.Println("********* threads still working")
		   		case <-done:
		   			break Analyzer

		   		}

		   	}
	*/
	log.Println("VALID ------ PORTS ARE :", len(validPorts))
	for _, val := range validPorts {
		log.Println(val)
	}
}
