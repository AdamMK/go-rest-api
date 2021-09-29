package main

import (
"fmt"
"log"
"net/http"
)

const ServerAddr = "localhost:8080"

func main() {
serv := NewServer()
fmt.Println("Server is running on", ServerAddr)
log.Fatal(http.ListenAndServe(ServerAddr, serv))
}