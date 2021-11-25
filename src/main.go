package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello")
	fileServer := http.FileServer(http.Dir("dist"))
	http.Handle("/", fileServer)

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}

}
