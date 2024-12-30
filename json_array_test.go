package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Kriti",
		LastName:  "Mauludin",
		Age:       20,
		Hobbies:   []string{"code", "learn"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Kriti","LastName":"Mauludin","Age":20,"Hobbies":["code","learn"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "KM",
		Addresses: []Address{
			{
				Street:     "JL.Code",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Test",
				Country:    "Indonesia",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"KM","LastName":"","Age":0,"Hobbies":null,"Addresses":[{"Street":"JL.Code","Country":"Indonesia","PostalCode":"9999"},{"Street":"Test","Country":"Indonesia","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"JL.Code","Country":"Indonesia","PostalCode":"9999"},{"Street":"Test","Country":"Indonesia","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addresses := []Address{
		{
			Street:     "JL.Code",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Test",
			Country:    "Indonesia",
			PostalCode: "8888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
