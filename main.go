package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	ws := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	http.HandleFunc("/ws", ws.wsHandler)

	fmt.Println("Started Server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
