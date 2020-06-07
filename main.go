package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var help bool

const (
	notifyClient = "notify"
	addItem      = "add"
	client       = "client"
	purchase     = "purchase"
	list         = "list"
	listShort    = "ls"
	mail         = "mail"
	sms          = "sms"
)

func init() {
	flag.BoolVar(&help, "h", false, "Help command")
	flag.BoolVar(&help, "help", false, "Help command")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No command entered. Try `buy-event -help` or `buy-event -h`")
		os.Exit(0)
	}
	flag.Parse()

	if help {
		fmt.Println("help called, not implemented for now")
		os.Exit(0)
		//TODO print help message
	}

	switch args[0] {
	case addItem:
		if len(args) < 2 {
			fmt.Println("Missing [option] - client/purchase")
			os.Exit(0)
		}
		switch args[1] {
		case client:
			err := AddNewClient()
			if err != nil {
				fmt.Println("Failed adding new client:", err)
			}
			endCommand(client)

		case purchase:
			err := AddNewPurchase()
			if err != nil {
				fmt.Printf("Failed adding new purchase: %s\n", err.Error())
			}
			endCommand(purchase)
		default:
			fmt.Println("Invalid option. Select one of (client, purchase)")
			break
		}

	case notifyClient:
		if len(args) < 2 {
			fmt.Println("Missing argument - [purchase id]")
			break
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid purchase id! Please enter an integer value")
			break
		}
		opt := mail // default value
		if len(args) >= 3 && args[2] == sms {
			opt = sms
		}
		err = Notify(id, opt)
		if err != nil {
			fmt.Printf("Failed sending notification: %s\n", err.Error())
		}
		fmt.Printf("Client has successfully been notified by %s\n", opt)

	case client:
		if len(args) < 2 {
			fmt.Println("Missing [command] - ls/list")
			break
		}
		if args[1] == listShort || args[1] == list {
			err := ListClients()
			if err != nil {
				fmt.Printf("Failed getting client list: %s\n", err.Error())
			}
		}

	case purchase:
		if len(args) < 2 {
			fmt.Println("Missing [command] - ls/list")
			break
		}
		if args[1] == listShort || args[1] == list {
			err := ListPurchases()
			if err != nil {
				fmt.Printf("Failed getting purchase list: %s\n", err.Error())
			}
		}

	default:
		fmt.Println("Unknown command. Try `buy-event -help` or `buy-event -h`")
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func endCommand(c string) {
	switch c {
	case client:
		fmt.Println("Client has successfully been added")
	case purchase:
		fmt.Println("Purchase has successfully been added")
	default:
		fmt.Println("end of command")
	}
	time.Sleep(2 * time.Second)
	clear()
}
