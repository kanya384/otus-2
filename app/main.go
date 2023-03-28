package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort = 8000
)

func main() {
	http.HandleFunc("/health/", successResponseHandler)
	http.HandleFunc("/", getHostNameHandler)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("app started on port %d\n", defaultPort)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	fmt.Println("gracefull shutdown...")
}

func successResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func getHostNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name, err := os.Hostname()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"hostname": nil})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"hostname": name})
}
