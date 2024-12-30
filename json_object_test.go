package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Hobbies   []string
	Addresses []Address
}

type Bank struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Kriti",
		LastName:  "Mauludin",
		Age:       20,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
