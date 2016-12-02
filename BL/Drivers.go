package BL

import (
	"encoding/json"
	"github.com/bnsd55/ionic-project/DAL"
	"log"
)

// Returns all drivers as JSON
func GetAllDrivers() string {
	allDrivers, err := DAL.Users()

	if err != nil {
		log.Panic(err)
	}

	returnJSON, _ := json.Marshal(allDrivers)
	return string(returnJSON)
}

// Returns driver by ID as JSON
func GetDriverByID(uuid string) string {
	allDrivers, err := DAL.UserByUUID(uuid)

	if err != nil {
		log.Panic(err)
	}

	returnJSON, _ := json.Marshal(allDrivers)
	return string(returnJSON)
}
