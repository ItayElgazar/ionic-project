package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/bnsd55/ionic-project/BL"
	"github.com/julienschmidt/httprouter"
	"github.com/bnsd55/ionic-project/DAL"
	"encoding/json"
	"github.com/bnsd55/ionic-project/Models"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Drivers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, BL.GetAllDrivers())
}

func DriverByUuid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uuid := params.ByName("uuid")

	fmt.Fprint(w, BL.GetDriverByID(uuid))
}


func createClientRoute(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	var client Models.Client
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		log.Panic(err)
	}

	affectedRows := DAL.CreateClient(client)

	if affectedRows > 0 {
		fmt.Fprint(w, true)
		// TODO: send verification code to user's phone by calling some function that gets the code we've generated and send it to phone
	} else {
		fmt.Fprint(w, false)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := httprouter.New()

	// Driver Route
	router.GET("/", Index)
	router.GET("/drivers", Drivers)
	router.GET("/driver/:uuid", DriverByUuid)

	// Client Routes
	router.POST("/clients", createClientRoute)


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
