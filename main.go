package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/bnsd55/ionic-project/BL"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

/*
func index(res http.ResponseWriter, req *http.Request) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=require",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	db, err := sql.Open("postgres", dbInfo)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, errr := db.Query("SELECT * FROM users")
	defer rows.Close()

	if errr != nil {
		log.Fatal(errr)
	}

	var results = make([]string, 2)

	for rows.Next() {
		var id int
		var username string
		err = rows.Scan(&id, &username)
		results = append(results, fmt.Sprintf("id: %v, username: %v", id, username))
	}

	slcB, _ := json.Marshal(results)
	io.WriteString(res, string(slcB))
}

func name(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ben")
}
*/

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Drivers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, BL.GetAllDrivers())
}

func DriverById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		fmt.Fprint(w, "An error occuer when trying to convert string to int")
		return
	}

	fmt.Fprint(w, BL.GetDriverByID(id))
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/drivers", Drivers)
	router.GET("/driver/:id", DriverById)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatalln("Server error: ", err)
	}
}
