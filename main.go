package main

import (
	"log"
	"net/http"
	"os"
	"io"
	"database/sql"
	"fmt"
	"encoding/json"
	_ "github.com/lib/pq"
)


func index (res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "postgres://bzvpyynezpuruw:xpyHRQ1ScLn0ZCGqsJLxOficyO@ec2-54-75-230-123.eu-west-1.compute.amazonaws.com/derkj04fmn4ff8?sslmode=require")

	if err != nil {
		log.Fatal(err)
	}

	rows, errr := db.Query("SELECT * FROM users")

	if errr != nil {
		log.Fatal(errr)
	}

	defer rows.Close()

	var results []string

	for rows.Next() {
		var id int
		var username string
		err = rows.Scan(&id, &username)
		results = append(results, fmt.Sprintf("id: %v, username: %v", id, username))
	}

	slcB, _ := json.Marshal(results)
	io.WriteString(res, string(slcB))
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
