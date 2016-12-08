package Models

import "time"

type Client struct {
	Id               int       `json:"id"`
	Uuid             string    `json:"uuid"`
	Name             string    `json:"name"`
	PhoneNumber      string    `json:"phone_number"`
	CreatedAt        time.Time `json:"created_at"`
	Activated        bool      `json:"activated"`
	VerificationCode string    `json:"verification_code"`
	Email            string    `json:"email"`
}
