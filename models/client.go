package models

type Client struct {
	Id          int    `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	Email       string `json:"email"`
}