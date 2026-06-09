package main

import (
    "html/template"
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

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, info)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
