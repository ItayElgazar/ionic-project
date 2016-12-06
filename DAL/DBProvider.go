package DAL

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

// Database connection details
const (
	DBUser     = "bzvpyynezpuruw"
	DBPassword = "xpyHRQ1ScLn0ZCGqsJLxOficyO"
	DBName     = "derkj04fmn4ff8"
	DBHost     = "ec2-54-75-230-123.eu-west-1.compute.amazonaws.com"
)

func init() {
	var err error
	dbConnectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=require", DBUser, DBPassword, DBName, DBHost)
	DB, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Generate verification code to verify user
func generateVerificationCode () string {
	return "1234"
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
}