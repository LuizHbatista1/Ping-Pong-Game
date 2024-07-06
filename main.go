package main

import (
	"fmt"
	"time"
)

func pinger(ping <-chan string, pong chan<- string) {
	for m := range ping {
		printAndDelay(m)
		pong <- "pong"
	}
}

func ponger(ping chan<- string, pong <-chan string) {
	for m := range pong {
		printAndDelay(m)
		ping <- "ping"
	}
}

func printAndDelay(msg string) {
	fmt.Println(msg)
	time.Sleep(time.Second)
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go pinger(ping, pong)

	go ponger(ping, pong)

	ping <- "ping"

	for {
	}
}
