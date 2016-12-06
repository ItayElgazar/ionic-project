package Routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/bnsd55/ionic-project/BL"
)

func GetClientsRouting(router *httprouter.Router) {

	// Create new client
	router.POST("/clients", BL.CreateClient)
}
