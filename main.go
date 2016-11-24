package main

import (
	"log"
	"net/http"
	"os"
	"io"
	"gopkg.in/pg.v5"
	"fmt"
)

type User struct {
	Id     int64
	Name   string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Name)
}

func index (res http.ResponseWriter, req *http.Request) {
	db := pg.Connect(&pg.Options{
		User:     "bzvpyynezpuruw",
		Database: "derkj04fmn4ff8",
		Password: "xpyHRQ1ScLn0ZCGqsJLxOficyO",
		Addr:     "54.75.230.123:5432",
	})

	user1 := &User{
		Name:   "admin",
	}

	err := db.Insert(user1)
	if err != nil {
		panic(err)
	}

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
