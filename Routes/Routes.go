package Routes

import (
	"github.com/julienschmidt/httprouter"
	"os"
	"fmt"
	"net/http"
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

	fmt.Println(pwd)

	router.ServeFiles("/static/*filepath", http.Dir(pwd))

	return router
}

