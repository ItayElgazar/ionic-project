package DAL

import (
	"time"
	"log"
	"github.com/bnsd55/ionic-project/Models"
)

func CreateClient(client Models.Client) int64 {
	query := "INSERT INTO clients (uuid, name, phone_number, created_at, activated, verification_code) VALUES ($1, $2, $3, $4, $5, $6)"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	defer stmt.Close()
	res, err := stmt.Exec(createUUID(), client.Name, client.PhoneNumber, time.Now(), false, generateVerificationCode())

	 if err != nil {
		 log.Panic(err)
		 return 0
	 }

	affectedRows, err := res.RowsAffected()

	return affectedRows
}
