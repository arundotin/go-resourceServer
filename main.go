package main

import (
	"fmt"
	"log"
	"marunk20/rs/server"
	"net/http"
)


func main() {
	fmt.Println("Starting server..")
	router:= server.RegisterRoutesInServer()
	log.Fatal(http.ListenAndServe(":8080", router))
}
