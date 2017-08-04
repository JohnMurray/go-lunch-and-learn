package main

import (
	"fmt"
	"runtime"
	"time"
)

func responder(msg string, init bool, in chan string, out chan string) {
	i := 0
	if init {
		out <- msg
	}
	for {
		received_msg := <-in
		fmt.Println(received_msg)
		out <- msg

		// slow the loop and only run for 5 iterations
		i += 1
		if i == 5 {
			break
		}
		time.Sleep(100)
	}
}

func main() {
	pinger := make(chan string, 1)
	ponger := make(chan string, 1)

	go responder("Ping", true, pinger, ponger)
	go responder("Pong", false, ponger, pinger)

	runtime.Goexit()
}
