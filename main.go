package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const ServerAddr = "localhost:8080"

func main() {

	serv := NewServer()

	//create a channel to catch exit signal
	quitSig := make(chan os.Signal, 1)

	//relay signals to quitSig channel - mainly for os.Interupt
	signal.Notify(quitSig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//needed to create another go routine to be able to start a server
	go func() {
		if err := http.ListenAndServe(ServerAddr, serv); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server is running on", ServerAddr)

	//receive output from quitSig
	<-quitSig
	log.Print("Server Stopped")
}