package main // Defines an executable Go program.

import (
	"fmt"
	"net/http" //Go standard library package for HTTP servers and clients.
)

func main() {

	// Register a handler for the root URL "/".
	// When the browser opens localhost:8080/, this function handles the request.
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		http.ServeFile(w, r, "web/templates/index.html")
	})

	// Print the server address in the terminal.
	fmt.Println("Pingo running on http://localhost:8080")

	// Start the HTTP server and listen on port 8080.
	// nil means Go uses the default HTTP router (DefaultServeMux).
	err := http.ListenAndServe(":8080", nil)

	// Print an error if the server fails to start or stops unexpectedly.
	if err != nil{
		fmt.Println("Server error:", err)
	}

}