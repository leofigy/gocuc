package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	a "github.com/seancfoley/ipaddress-go/ipaddr"
)

func main() {

	// nmap -sn 172.27.6.0/24
	// nmap -sn  192.168.2.0/24

	var addrs []net.IP
	block := a.NewIPAddressString("192.168.2.0/24").GetAddress()
	for i := block.Iterator(); i.HasNext(); {
		addrs = append(addrs, i.Next().GetNetIP())
	}

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	for _, value := range addrs {
		target := fmt.Sprintf("http://%s:%d", value, 8080)
		fmt.Println(target)
		resp, err := client.Get(target)
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Println(resp)

		data, err := io.ReadAll(resp.Request.Response.Body)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(data)
	}
}
