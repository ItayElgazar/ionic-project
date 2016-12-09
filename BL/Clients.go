package BL

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/bnsd55/ionic-project/DAL"
	"github.com/bnsd55/ionic-project/Models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Adds client to db
func CreateClient(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//defer r.Body.Close()
	var client Models.Client
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		log.Fatal(err)
	}

	if govalidator.IsNumeric(client.PhoneNumber) {
		doesUserCreated := DAL.CreateClient(client)

		if doesUserCreated {
			fmt.Fprint(w, true)
			// TODO: send verification code to user's phone by calling some function that gets the code we've generated and send it to phone
		} else {
			fmt.Fprint(w, false)
		}

	} else {
		fmt.Fprint(w, false)
	}
}

// Verify client
func VerifyClient(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	var client DAL.Client
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		log.Panic(err)
	}

	if govalidator.IsNumeric(client.PhoneNumber) && govalidator.IsNumeric(client.VerificationCode) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, client.Verify())
	} else {
		fmt.Fprint(w, false)
	}
}

// Update user's email
func UpdateClientEmail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	var client DAL.Client
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		log.Panic(err)
	}

	if govalidator.IsNumeric(client.PhoneNumber) && govalidator.IsEmail(client.Email) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, client.UpdateClientEmail())
	} else {
		fmt.Fprint(w, false)
	}

}
