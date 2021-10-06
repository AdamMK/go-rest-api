package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const ServerAddr = ":8282"

func main() {

	serv := NewServer()

	//create a channel to catch exit signal
	quitSig := make(chan os.Signal, 1)

	//relay signals to quitSig channel - mainly for os.Interrupt
	signal.Notify(quitSig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//needed to create another go routine to be able to start a server
	go func() {
		if err := http.ListenAndServe(ServerAddr, serv); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()
	log.Println("Server is running on", ServerAddr)

	//receive output from quitSig
	<-quitSig
	log.Print("Server Stopped")
}