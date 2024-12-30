package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("bankNew.json")
	encoder := json.NewEncoder(writer)

	bank := Bank{
		Name:  "Bank Rakyat Indonesia",
		Alias: "BRI",
	}

	encoder.Encode(bank)
}
