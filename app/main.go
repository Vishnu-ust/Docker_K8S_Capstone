package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type AppInfo struct {
	AppName     string
	Version     string
	Hostname    string
	CurrentTime string
	Environment string
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	info := AppInfo{
		AppName:     "ShadowXLab Cyber Learning Portal",
		Version:     "1.0.0",
		Hostname:    hostname,
		CurrentTime: time.Now().Format("2006-01-02 15:04:05"),
		Environment: os.Getenv("APP_ENV"), // set APP_ENV=production or development
	}

	// parse and execute template
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, info); err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Log the request and what was served so kubectl logs shows activity
	log.Printf("%s - %s %s - served version=%s env=%s host=%s", r.RemoteAddr, r.Method, r.URL.Path, info.Version, info.Environment, info.Hostname)
}

func main() {
	// include timestamps with microseconds for clearer logs
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	http.HandleFunc("/", handler)
	addr := ":8080"
	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
