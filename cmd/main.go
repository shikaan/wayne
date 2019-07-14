package main

import (
	"log"
	"net/http"

	"example.com/goshort/internal"
)

func main() {
	// http.HandleFunc("/", internal.SayHello)

	http.HandleFunc("/api/urls", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			internal.ListURLs(writer, request)
		} else if request.Method == http.MethodPost {
			internal.StoreURL(writer, request)
		} else {
			writer.WriteHeader(405)
			writer.Write([]byte("405 - Method not allowed"))
		}
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			internal.Redirect(writer, request)
		} else {
			writer.WriteHeader(405)
			writer.Write([]byte("405 - Method not allowed"))
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic("Cannot launch the server", err)
	}
}
