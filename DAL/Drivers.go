package DAL

import (
	"log"
)

type Driver struct {
	Id int
	Username string
}

// This function gets a slice of *Driver from Database and returns it as JSON
func (db *DB) GetAllDrivers() ([]*Driver, error) {
	db, err := GetPgConnection()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	drivers := make([]*Driver, 0)

	for rows.Next() {
		driver := new(Driver) // Allocate memory for Driver and returns pointer to it
		err = rows.Scan(&driver.Id, &driver.Username)

		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil
}

// This function gets a pointer to driver struct which has the driver id information and returns it as JSON
func (db *DB) GetDriverById(id int) (*Driver, error) {
	db, err := GetPgConnection()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	driver := new(Driver)

	for row.Next() {
		err = row.Scan(&driver.Id, &driver.Username)

		if err != nil {
			return nil, err
		}
	}

	return driver, nil
}