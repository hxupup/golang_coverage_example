package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var endRunning chan bool

func handleSignals(stopFunction func()) {
	var callback sync.Once

	// On ^C or SIGTERM, gracefully stop the sniffer
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		callback.Do(stopFunction)
	}()
}

func stop() {
	endRunning <- true
}

func main() {
	endRunning = make(chan bool, 1)
	handleSignals(stop)
	fmt.Printf("test!\n")
	<-endRunning
}
