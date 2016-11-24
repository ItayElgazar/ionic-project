package main

import (
	"log"
	"net/http"
	"os"
	"io"
)

func index (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func name (res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ben")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/ben", name)

	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatalln("Server error: ", err)
	}
}
