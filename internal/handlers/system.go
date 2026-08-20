package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/maglovskiNenad/pingo/internal/metrics"
)

type SystemData struct {
	CPU float64
}

func SystemHandler(w http.ResponseWriter, r *http.Request) {
	// Read the current CPU usage from the Linux system.
	cpu, err := metrics.CPUUsage()
	if err != nil {
		http.Error(w, "Failed to read CPU usage", http.StatusInternalServerError)
		return
	}

	// Load the system HTML template.
	tmpl, err := template.ParseFiles("web/templates/system.html")
	if err != nil {
		http.Error(w, "Failed to load system template", http.StatusInternalServerError)
		return
	}
	

	// Prepare the data for the HTML template.
	data := SystemData{
		CPU: cpu,
	}

	// Render the template and send it to the browser.
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Template error:", err)
	}
}

func CPUAPIHandler(w http.ResponseWriter, r *http.Request) {
	//Read the current CPU usafe
	cpu, err := metrics.CPUUsage()
	if err != nil {
		http.Error(w, "Failed to read CPU usage", http.StatusInternalServerError)
		return
	}

	// Tell the browser that the response contains JSON.
	w.Header().Set("Content-Type", "application/json")

	//Tell the browser that the respones contains JSON
	err = json.NewEncoder(w).Encode(map[string]float64{
		"cpu": cpu,
	})

	if err != nil {
		fmt.Println("JSON error:", err)
	}
}

func MemoryAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Read current RAM information.
	memory, err := metrics.ReadMemory()
	if err != nil {
		http.Error(w, "Failed to read memory usage", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Convert KB to GB.
	totalGB := float64(memory.TotalKB) / 1024 / 1024
	usedGB := float64(memory.UsedKB) / 1024 / 1024

	err = json.NewEncoder(w).Encode(map[string]float64{
		"usage": memory.Usage,
		"used":  usedGB,
		"total": totalGB,
	})

	if err != nil {
		fmt.Println("JSON error:", err)
	}
}