package main

import (
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./static/index.html")

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w,r, "./templates/login.html")
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",homeHandler)
	mux.HandleFunc("/login", loginHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

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