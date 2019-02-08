package main
import (
    "net"
    "os"
    "time"
)
const (
    CONN_HOST = "10.100.23.253"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
    SERV_ADDR="10.100.23.242:33546"
)
func client() {
    TCPmessage := "Connect to:"+CONN_HOST+":"+CONN_PORT+"\000"

    servAddr := SERV_ADDR
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

    _, err = conn.Write([]byte(TCPmessage))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    println("write to server = ", TCPmessage)

    reply := make([]byte, 1024)

    _, err = conn.Read(reply)
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    println("reply from server=", string(reply))
    
    conn.Close()
}
func server(){
	println("server starting")
	ln, err := net.Listen("tcp",CONN_HOST+":"+CONN_PORT)
	if err != nil {
	println("Listen to server failed:", err.Error())
        os.Exit(1)
}
defer ln.Close()
// A common pattern is to start a loop to continously accept connections
for {
	println("Accepting connections")
    //accept connections using Listener.Accept()
	conn, err := ln.Accept()
	if err!= nil {
		println("Accept connection failed:", err.Error())
        os.Exit(1)
	}
    //It's common to handle accepted connection on different goroutines
    println("Handling connections")
	go handleConnection(conn)
}

}
func handleConnection(c net.Conn) {
println("Handling connection")
 buf := make([]byte, 1024)

	length, err := c.Read(buf)
	if err != nil {
		println("Read failed:", err.Error())
        os.Exit(1)
	}

	println(string(buf[0:length]))

	c.Write([]byte("Message received"))

	c.Close()
}
func main(){
	go server()
	go client()
	time.Sleep(time.Second*15)
}
