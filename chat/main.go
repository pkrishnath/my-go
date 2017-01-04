package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html> <head> <title>Chat</title> </head> <body> this is chat serevr</body> </html>`))
	})

	// start web server
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
