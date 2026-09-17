package main

import (
	"bytes"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello", handleHelloParameterized)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Welcome to he Homepage!\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}

}

func handleGoodbye(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Goodbye!\n"))

	if err != nil {
		slog.Error("error waiting response", "err", err)
		return
	}
}

func handleHelloParameterized(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	if user == "" {
		http.Error(w, "Missing user parameter", http.StatusBadRequest)
		return
	}

	var output bytes.Buffer

	output.WriteString("Hello, ")
	output.WriteString(user)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())

	if err != nil {
		slog.Error("error writing response body", "err", err)
	}
}
