package utils

import (
	"bufio"
	"buy-event/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

func AddToFile(in interface{}) (err error) {
	dataJson, err := json.Marshal(in)
	if err != nil {
		return
	}

	switch in.(type){
	case models.Client:
		file, err := os.OpenFile("files/clients.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend)
		defer file.Close()
		if err != nil {
			return err
		}
		if _, err = file.Write(dataJson); err != nil {
			return err
		}

	case models.Purchase:
		err = ioutil.WriteFile("files/purchases.txt", dataJson, 0644)
	}

	return
}

func GetClientList() (clients []models.Client, err error) {
	file, err := os.Open("files/clients.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var client models.Client
		line := scanner.Bytes()
		if err = json.Unmarshal(line, &client); err != nil {
			return
		}
		clients = append(clients, client)
	}
	return
}

func GetPurchaseList() (purchases []models.Purchase, err error) {
	file, err := os.Open("files/purchases.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var purchase models.Purchase
		line := scanner.Bytes()
		if err = json.Unmarshal(line, &purchase); err != nil {
			return
		}
		purchases = append(purchases, purchase)
	}
	return
}