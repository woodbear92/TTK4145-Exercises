// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	//"time"
)


func server(increment chan int, decrement chan int, get_value chan int) {
	var i = 0
	for {
		select {
			case <- increment:
				i += 1
			case <- decrement:
				i-= 1
			case get_value <- i:


	}
}
}

func incrementing(increment chan int, quit_inc chan int) {
	for j := 0; j < 999999; j++ {
		increment <- 0
	}
	quit_inc <- 1
}

func decrementing(decrement chan int, quit_dec chan int) {
	for j := 0; j < 1000000; j++ {
		decrement <- 0
	}
	quit_dec <- 1

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	increment := make (chan int)
	decrement := make (chan int)
	get_value := make (chan int)
	quit := make (chan int)

	go incrementing(increment, quit)
	go decrementing(decrement, quit)
	go server(increment, decrement, get_value)

	<- quit
	<- quit

	Println("The magic number is:", <-get_value)

}
