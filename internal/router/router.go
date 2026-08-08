package router

import "net/http"

// RegisterRoutes defines all HTTP routes used by the Pingo web application.
func RegisterRoutes() {

	// Serve the main dashboard page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

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