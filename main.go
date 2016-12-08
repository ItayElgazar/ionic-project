package main

import (
	"github.com/bnsd55/ionic-project/Routes"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	err := http.ListenAndServe(":"+port, Routes.GetRoutes())

	if err != nil {
		log.Fatalln("Server error: ", err)
	}
}
