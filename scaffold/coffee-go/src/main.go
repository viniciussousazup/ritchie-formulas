package main

import (
	"coffee/pkg/coffee"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	coffee.GiveMeSomeCoffee(loadInputs())
}

type output struct {
	V1 string
	V2 string
	V3 string
}

func loadInputs() coffee.Inputs {
	delivery, _ := strconv.ParseBool(os.Getenv("DELIVERY"))
	name := os.Getenv("NAME")
	coffeeType := os.Getenv("COFFEE_TYPE")

	outputF := os.Getenv("RIT_OUTPUT_FILE")
	json, _ := json.Marshal(output{
		V1: "a",
		V2: "b",
		V3: "c;",
	})
	ioutil.WriteFile(outputF, json, 0755)

	return coffee.Inputs{
		Name:       name,
		CoffeeType: coffeeType,
		Delivery:   delivery,
	}
}
