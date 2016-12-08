package Routes

import (
	"github.com/bnsd55/ionic-project/BL"
	"github.com/julienschmidt/httprouter"
)

func GetClientsRouting(router *httprouter.Router) {

	// Create new client
	router.POST("/clients", BL.CreateClient)

	// Verify client
	router.POST("/verifyClient", BL.VerifyClient)

	// Update client's email
	router.PUT("/updateClientEmail", BL.UpdateClientEmail)
}
