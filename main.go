package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"github.com/bnsd55/ionic-project/Routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	err := http.ListenAndServe(":" + port, Routes.GetRoutes())
	if err != nil {
		log.Fatalln("Server error: ", err)
	}
}
