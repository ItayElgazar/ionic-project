package Routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func GetRoutes() *httprouter.Router {
	router := httprouter.New()

	GetClientsRouting(router)

	// ==================================== //
	// =========== Static files =========== //
	// ==================================== //

	pwd, error := os.Getwd()

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	router.ServeFiles("/static/*filepath", http.Dir(pwd))

	return router
}
