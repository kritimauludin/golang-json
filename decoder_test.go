package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("bank.json")
	decoder := json.NewDecoder(reader)

	bank := &Bank{}
	decoder.Decode(bank)

	fmt.Println(bank)
}
