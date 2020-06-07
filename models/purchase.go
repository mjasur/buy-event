package models

type Purchase struct {
	Id       int    `json:"id"`
	ClientId int    `json:"client_id"`
	Product  string `json:"product"`
	Price    int    `json:"price"`
}
