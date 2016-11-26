package BL

import (
	"encoding/json"
	"github.com/bnsd55/ionic-project/DAL"
	"log"
)

// Returns all drivers as JSON
func GetAllDrivers() string {
	allDrivers, err := new(DAL.DB).GetAllDrivers()

	if err != nil {
		log.Panic(err)
	}

	returnJSON, _ := json.Marshal(allDrivers)
	return string(returnJSON)
}

// Returns driver by ID as JSON
func GetDriverByID(id int) string {
	allDrivers, err := new(DAL.DB).GetDriverById(id)

	if err != nil {
		log.Panic(err)
	}

	returnJSON, _ := json.Marshal(allDrivers)
	return string(returnJSON)
}
