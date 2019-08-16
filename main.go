package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/korovkin/gotils"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("=> request", r.RequestURI)
	log.Println("=> response", r.RequestURI)
	http.Error(w, "ok", http.StatusOK)
}

func main() {
	port := 9001
	address := fmt.Sprintf("0.0.0.0:%d", port)

	log.Println("hello")
	log.Println(fmt.Sprintf(" => Running on http://%s", address))

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(address, nil)
	gotils.CheckFatal(err)
}
