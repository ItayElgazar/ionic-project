package DAL

import (
	"time"
	"log"
)

type Clients struct {
	Id int
	Uuid string
	Name string
	PhoneNumber string
	CreatedAt time.Time
}

func CreateClient(phoneNumber string) error {
	query := "INSERT INTO clients (uuid, phone_number, created_at) VALUES ($1, $2, $3)"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(createUUID(), phoneNumber, time.Now())

	 if err != nil {
		 log.Panic(err)
	 }

	affectedRows, err := res.RowsAffected()

	log.Println(affectedRows)

	return nil
}
