package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello/", handleHelloParameterized)

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

func handleHelloParameterized(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "not implemented", http.StatusInternalServerError)

}
