package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//exit process immediately upon sigterm
	handleSigTerms()

}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
