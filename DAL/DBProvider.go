package DAL

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

// All DB functions
type DBProvider interface {
	Drivers
}

// All driver functions
type Drivers interface {
	GetAllDrivers() ([]*Driver, error)
	GetDriverById(id int) (*Driver, error)
}

// DB structg
type DB struct {
	*sql.DB
}

// Database connection details
const (
	DBUser     = "bzvpyynezpuruw"
	DBPassword = "xpyHRQ1ScLn0ZCGqsJLxOficyO"
	DBName     = "derkj04fmn4ff8"
	DBHost     = "ec2-54-75-230-123.eu-west-1.compute.amazonaws.com"
)

// Connection to PostgreSQL
func GetPgConnection() (*DB, error) {
	dbConnectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=require", DBUser, DBPassword, DBName, DBHost)
	db, err := sql.Open("postgres", dbConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}