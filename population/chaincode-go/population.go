/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"
    "encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
//Insert struct field in alphabetic order => to achieve determinism accross languages
// golang keeps the order when marshal to json but doesn't order automatically
type Person struct {
    Address         string `json:"Address"`
    City            string `json:"City"`
	Id              string `json:"Id"`
	Name            string `json:"Name"`
    Status          string `json:"Status"`
    Surname         string `json:"Surname"`
    TelephoneNumber string `json:"TelephoneNumber"`
}

func (s *SmartContract) AddPerson(ctx contractapi.TransactionContextInterface, address string, city string, id string, name string, status string, surname string, telephoneNumber string) error {
    exists, err := s.PersonExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the person with id: %s already exists", id)
	}

    person := Person{
		Address:            address,
		City:               city,
		Id:                 id,
		Name:               name,
		Status:             status,
        Surname:            surname, 
        TelephoneNumber:    telephoneNumber,
	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		return err
	}

    err = ctx.GetStub().SetEvent("Change", personJSON)
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	return ctx.GetStub().PutState(id, personJSON)
}

func (s *SmartContract) GetPerson(ctx contractapi.TransactionContextInterface, id string) (*Person, error) {
    personJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if personJSON == nil {
		return nil, fmt.Errorf("the person %s does not exist", id)
	}

	var person Person
	err = json.Unmarshal(personJSON, &person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (s *SmartContract) ChangePersonData(ctx contractapi.TransactionContextInterface, address string, city string, id string, name string, status string, surname string, telephoneNumber string) error {
    exists, err := s.PersonExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the person with id %s does not exist", id)
	}

	// overwriting original asset with new asset
	person := Person{
		Address:            address,
		City:               city,
		Id:                 id,
		Name:               name,
		Status:             status,
        Surname:            surname, 
        TelephoneNumber:    telephoneNumber,
	}
	personJSON, err := json.Marshal(person)
	if err != nil {
		return err
	}

    err = ctx.GetStub().SetEvent("Change", personJSON)
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	return ctx.GetStub().PutState(id, personJSON)
}


func (s *SmartContract) PersonExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	personJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return personJSON != nil, nil
}

func main() {
	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
