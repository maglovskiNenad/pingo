package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/maglovskiNenad/pingo/internal/metrics"
)

type DashboardData struct {
	CPU float64
}

// RegisterRoutes defines all HTTP routes used by the Pingo web application.
func RegisterRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read the current CPU usage from the Linux system.
		cpu, err := metrics.CPUUsage()
		if err != nil {
			http.Error(w, "Failed to read CPU usage", http.StatusInternalServerError)
			return
		}
	
		// Load the dashboard HTML template.
		tmpl, err := template.ParseFiles("web/templates/index.html")
			if err != nil {
				http.Error(w, "Failed to load template", http.StatusInternalServerError)
				return
		}
	
		// Prepare the data that will be available inside the HTML template.
		data := DashboardData{
			CPU: cpu,
		}
	
		// Render the HTML template and send it to the browser.
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println("Template error:", err)
		}
	})
	
	// Serve the main dashboard page.
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "web/templates/index.html")
	//})

	// Serve the system information page.
	http.HandleFunc("/system", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/system.html")
	})

	// Serve the network information page.
	http.HandleFunc("/network", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/network.html")
	})

	// Serve the logs information page.
	http.HandleFunc("/logs", func(w http.ResponseWriter,r *http.Request){
		http.ServeFile(w, r, "web/templates/logs.html") 
	})

	// Serve the services information page.
	http.HandleFunc("/service", func(w http.ResponseWriter,r *http.Request){
		http.ServeFile(w, r, "web/templates/services.html") 
	})

	// Serve static files such as CSS and JavaScript.
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}