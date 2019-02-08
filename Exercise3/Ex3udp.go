package main

import (
	"fmt"
	"net"
	"time"
)



func receive() {
    ServerAddr,err := net.ResolveUDPAddr("udp4",":20009")
    if err != nil {
		fmt.Println(err)
	}

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp4", ServerAddr)
    if err != nil {
		fmt.Println(err)
	}

    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}



func send() {
	time.Sleep(time.Second*1)

	ServerAddr, err := net.ResolveUDPAddr("udp4","10.100.23.242:20009")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp4", nil, ServerAddr)
	if err != nil {
		fmt.Println(err)
	}


	conn.Write([]byte("This is a UDP packet"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	go send()
	go receive()
	time.Sleep(time.Second*5)

}


