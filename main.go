package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type","text/plain: charset=utf-8")

	fmt.Fprintln(w, "Hello from Go!")
}

func main(){
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)

	server := &http.Server {
		Addr: "127.0.0.1:8080",
		Handler: mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	log.Println("Server listenig on http://127.0.0.1:8080")

	if err := server.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}