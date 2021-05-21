package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("server1, timestamp: %d", time.Now().Unix())
		w.Write([]byte(message))
	})
	go http.ListenAndServe(":3000", mux1)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("server2, timestamp: %d", time.Now().Unix())
		w.Write([]byte(message))
	})
	go http.ListenAndServe(":3001", mux2)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	fmt.Printf("server exit now. %v", <-c)
}
