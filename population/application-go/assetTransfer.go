/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"io/ioutil"
    "encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Person struct {
    Address         string `json:"Address"`
    City            string `json:"City"`
	Id              string `json:"Id"`
	Name            string `json:"Name"`
    Status          string `json:"Status"`
    Surname         string `json:"Surname"`
    TelephoneNumber string `json:"TelephoneNumber"`
}

type HistData struct {
    Data       Person `json:"Data"`
    Time  string   `json:"Time"`
}

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"

func main() {
	fmt.Println(Green + "============ application-golang starts ============" + Reset)
    fmt.Println(Green + "Initialization" + Reset)

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("population")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract("basic")

    fmt.Println(Green + "Initialization finished" + Reset)
    
    var command string

    fmt.Println(Green + "Write commands separately" + Reset)
    fmt.Println(Green + "Print help for help inforamtion" + Reset)

    for {
        fmt.Scanln(&command)
        switch command {
        case "help":
            fmt.Println("\thelp - Prints help information")
            fmt.Println("\tinsert - Insert person to registery. Arguments: (Address, City, Id, Name, Status, Surname, TelephoneNumber)")
            fmt.Println("\tupdate - Update person in registery. Arguments: (Address, City, Id, Name, Status, Surname, TelephoneNumber)")
            fmt.Println("\tread - Read person data. Arguments: (Id)")
            fmt.Println("\tgetHist - Read person data logs. Arguments: (Id)")
            fmt.Println("\texit - Exit the application");
        case "insert":
            var address, city, id, name, status, surname, telephoneNumber string
            fmt.Println("Address:")
            fmt.Scanln(&address)
            fmt.Println("City:")
            fmt.Scanln(&city)
            fmt.Println("Id:")
            fmt.Scanln(&id)
            fmt.Println("Name:")
            fmt.Scanln(&name)
            fmt.Println("Status:")
            fmt.Scanln(&status)
            fmt.Println("Surname:")
            fmt.Scanln(&surname)
            fmt.Println("TelephoneNumber:")
            fmt.Scanln(&telephoneNumber)
            
            _, err := contract.SubmitTransaction("AddPerson", address, city, id, name, status, surname, telephoneNumber)
            if err != nil {
                log.Printf(Red + "Faild to insert person: %v\n" + Reset, err)
            }
            
        case "update": 
            var address, city, id, name, status, surname, telephoneNumber string
            fmt.Println("Address:")
            fmt.Scanln(&address)
            fmt.Println("City:")
            fmt.Scanln(&city)
            fmt.Println("Id:")
            fmt.Scanln(&id)
            fmt.Println("Name:")
            fmt.Scanln(&name)
            fmt.Println("Status:")
            fmt.Scanln(&status)
            fmt.Println("Surname:")
            fmt.Scanln(&surname)
            fmt.Println("TelephoneNumber:")
            fmt.Scanln(&telephoneNumber)

            result, err := contract.SubmitTransaction("ChangePersonData", address, city, id, name, status, surname, telephoneNumber)
            if err != nil {
                log.Printf(Red + "Faild to update person: %v" + Reset, err)
            }
            fmt.Println(string(result))
        case "read": 
            var id string
            fmt.Println("Id:")
            fmt.Scanln(&id)

            result, err := contract.SubmitTransaction("GetPerson", id)
            if err != nil {
                log.Printf(Red + "Faild to read person data: %v" + Reset, err)
            }
            fmt.Println(string(result));
        case "getHist": 
            var id string
            fmt.Println("Id:")
            fmt.Scanln(&id)
            
            result, err := contract.SubmitTransaction("GetPersonHistory", id)
            if err != nil {
                log.Printf(Red + "Faild to read person data: %v" + Reset, err)
            }
            
            var data []HistData
            err = json.Unmarshal(result, &data)
            if err != nil {
                log.Printf(Red + "Faild to unmarshal data: %v" + Reset, err)
            }
            for _, v := range data {
                fmt.Println("{")
                fmt.Printf("\t%+v\n", v.Time)
                fmt.Printf("\t%+v\n", v.Data)
                fmt.Println("}")
            }

        case "exit":
            return
        default:
            fmt.Println(Red + "Wrong command" + Reset)
        }
        fmt.Println(Green + "Enter command" + Reset)
    }

	fmt.Println("============ application-golang ends ============")
}

func populateWallet(wallet *gateway.Wallet) error {
	fmt.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
