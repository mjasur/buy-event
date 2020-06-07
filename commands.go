package main

import (
	"buy-event/models"
	"buy-event/utils"
	"errors"
	"fmt"
)

func AddNewClient() (err error) {
	var id int
	var email string
	var phone string

	fmt.Println("Enter client id(int): ")
	_, err = fmt.Scanln(&id)
	if err != nil {
		return
	}

	fmt.Println("Enter client email(string): ")
	_, err = fmt.Scanln(&email)
	if err != nil {
		return
	}

	fmt.Println("Enter client phone number(string): ")
	_, err = fmt.Scanln(&phone)
	if err != nil {
		return
	}

	client := models.Client{id, phone, email}
	err = utils.AddToFile(client)
	return
}

func AddNewPurchase() (err error) {
	var id int
	var client_id int
	var product string
	var price int

	fmt.Println("Enter purchase id(int): ")
	_, err = fmt.Scanln(&id)
	if err != nil {
		return
	}

	fmt.Println("Enter client id(int): ")
	_, err = fmt.Scanln(&client_id)
	if err != nil {
		return
	}
	fmt.Println("Enter product(string): ")
	_, err = fmt.Scanln(&product)
	if err != nil {
		return
	}
	fmt.Println("Enter price(int): ")
	_, err = fmt.Scanln(&price)
	if err != nil {
		return
	}

	purchase := models.Purchase{id, client_id,product, price}

	err = utils.AddToFile(purchase)
	return
}

func Notify(id int, option string) (err error){
	switch option {
	case mail:
		var client_id int
		data := make(map[string]interface{})
		purchases, err := utils.GetPurchaseList()
		if err != nil {
			return err
		}
		found := false
		for _, v := range purchases{
			if v.Id == id {
				found = true
				client_id = v.ClientId
				data["price"] = v.Price
				data["product"] = v.Product
			}
		}
		if !found {
			return errors.New("purchase not found!")
		}
		clients, err := utils.GetClientList()
		if err != nil {
			return err
		}
		for _, v := range clients {
			if v.Id == client_id {
				data["email"] = v.Email
			}
		}
		err = utils.SendByEmail(data)
		if err != nil {
			return err
		}
	case sms:
		//TODO implement sms notification
	}
	return
}

func ListClients() (err error) {
	clients, err := utils.GetClientList()
	if err != nil {
		return
	}
	for _, v := range clients {
		fmt.Printf("%+v\n", v)
	}
	return
}

func ListPurchases() (err error) {
	purchases, err := utils.GetPurchaseList()
	if err != nil {
		return
	}
	for _, v := range purchases {
		fmt.Printf("%+v\n", v)
	}
	return
}