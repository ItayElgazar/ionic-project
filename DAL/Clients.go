package DAL

import (
	"github.com/bnsd55/ionic-project/Models"
	"log"
	"time"
)

// An alias to Client struct
type Client Models.Client

func CreateClient(client Models.Client) bool {
	query := "INSERT INTO clients (uuid, name, phone_number, created_at, activated, verification_code) VALUES ($1, $2, $3, $4, $5, $6)"
	stmt, err := DB.Prepare(query)
	if err != nil {
		//log.Fatal(err)
		return false
	}

	defer stmt.Close()
	res, err := stmt.Exec(createUUID(), client.Name, client.PhoneNumber, time.Now(), false, generateVerificationCode())

	if err != nil {
		//log.Panic(err)
		return false
	}

	affectedRows, err := res.RowsAffected()

	return (affectedRows != 0)
}

func (client *Client) Verify() bool {
	query := "SELECT verification_code FROM clients WHERE phone_number = $1"
	row := DB.QueryRow(query, client.PhoneNumber)

	var verification_code string
	row.Scan(&verification_code)

	// Change client status to active
	if verification_code == client.VerificationCode {
		query := "UPDATE clients SET activated = $1 WHERE phone_number = $2"
		stmt, err := DB.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return false
		}

		defer stmt.Close()
		res, err := stmt.Exec(true, client.PhoneNumber)

		if err != nil {
			log.Panic(err)
			return false
		}

		affectedRows, err := res.RowsAffected()

		return (affectedRows != 0)
	} else {
		return false
	}
}

func (client *Client) UpdateClientEmail() bool {
	query := "UPDATE clients SET email = $1 WHERE phone_number = $2"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer stmt.Close()
	res, err := stmt.Exec(client.Email, client.PhoneNumber)

	if err != nil {
		log.Panic(err)
		return false
	}

	affectedRows, err := res.RowsAffected()

	return (affectedRows != 0)
}
