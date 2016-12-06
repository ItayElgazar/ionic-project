package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/bnsd55/ionic-project/BL"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := httprouter.New()

	// Client Routes
	router.POST("/clients", BL.CreateClient)

	// Static files
	pwd, error := os.Getwd()

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	fmt.Println(pwd)

	router.ServeFiles("/static/*filepath", http.Dir(pwd))

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatalln("Server error: ", err)
	}
}
