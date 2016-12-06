package BL

import (
	"encoding/json"
	"fmt"
	"github.com/bnsd55/ionic-project/DAL"
	"github.com/bnsd55/ionic-project/Models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func CreateClient(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
