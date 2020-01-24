package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	PageTitle   string
	ClusterName string
	Hostname    string
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func addHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Origin-Cluster-Name", getEnv("ClusterName", "Unknown (env unset)"))
	(*w).Header().Set("Origin-Hostname", getHostname())
}

func main() {
	tmpl := template.Must(template.ParseFiles("src/html/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addHeader(&w)
		data := PageData{
			PageTitle:   getEnv("PageTitle", "Demo test page"),
			ClusterName: getEnv("ClusterName", "Unknown (env unset)"),
			Hostname:    getHostname()}
		tmpl.Execute(w, data)
	})

	log.Println("Server listening on port " + getEnv("ListenPort", "8080"))
	log.Fatal(http.ListenAndServe(":"+getEnv("ListenPort", "8080"), nil))
}
