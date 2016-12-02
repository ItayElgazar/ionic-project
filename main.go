package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/bnsd55/ionic-project/BL"
	"github.com/julienschmidt/httprouter"
	"os/signal"
	"syscall"
	"github.com/bnsd55/ionic-project/DAL"
	"encoding/json"
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
	type UserPhone struct {
		PhoneNumber string `json:"phone_number"`
	}

	decoder := json.NewDecoder(r.Body)

	var phone UserPhone
	err := decoder.Decode(&phone)
	if err != nil {
		log.Panic(err)
	}

	DAL.CreateClient(phone.PhoneNumber);

	log.Println(phone.PhoneNumber)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/drivers", Drivers)
	router.GET("/driver/:uuid", DriverByUuid)
	router.POST("/client", createClientRoute)

	pwd, error := os.Getwd()
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println(pwd)

	router.ServeFiles("/static/*filepath", http.Dir(pwd))

	go func() {
		interruptChannel := make(chan os.Signal, 0)
		// look for system interruptions
		signal.Notify(interruptChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		// lock lower code until interruptChannel receives signal
		<-interruptChannel
		// Other cleanup tasks
		fmt.Println("Closing connection")
		fmt.Println("Saving session")

		// Accually Close DB session, maintain DATA integrity
		DAL.DB.Close()
		// Removes Temp, compiled JS files
		// os.RemoveAll("./public/assets/scripts/")
		// // Remove Temp, compiled stylesheets
		// os.RemoveAll("./public/assets/stylesheets/")
		// Explicitly call for system exit this is more graceful
		os.Exit(0)
	}()

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatalln("Server error: ", err)
	}
}
