package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/message", message)
	http.HandleFunc("/guzzle", urlGuzzler)
	fmt.Println("listening...")

	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. This is our first Go web app on Heroku!")
}

func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a secret message 🙂")
}

func urlGuzzler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		url := r.FormValue("url")

		if isValidUrl(url) {
			fmt.Fprint(w, url)
		} else {
			http.Error(w, "Invalid URL", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
