package Routes

import (
	"github.com/bnsd55/ionic-project/BL"
	"github.com/julienschmidt/httprouter"
)

func GetClientsRouting(router *httprouter.Router) {

	// Create new client
	router.POST("/clients", BL.CreateClient)
}
