package main // Defines an executable Go program.

import (
	"fmt"
	"net/http" //Go standard library package for HTTP servers and clients.

	"github.com/maglovskiNenad/pingo/internal/metrics"
	"github.com/maglovskiNenad/pingo/internal/router"
)

func main() {



	cpu, err := metrics.CPUUsage()

	if err != nil {
		fmt.Println("CPU error:", err)
	} else {
		fmt.Printf("CPU Usage: %.1f%%\n", cpu)
	}

	// RegisterRoutes defines all HTTP routes used by the Pingo web application.
	router.RegisterRoutes()

	// Print the server address in the terminal.
	fmt.Println("Pingo running on http://localhost:8080")

	// Start the HTTP server and listen on port 8080.
	// nil means Go uses the default HTTP router (DefaultServeMux).
	err = http.ListenAndServe(":8080", nil)

	
	// Print an error if the server fails to start or stops unexpectedly.
	if err != nil{
		fmt.Println("Server error:", err)
	}

}