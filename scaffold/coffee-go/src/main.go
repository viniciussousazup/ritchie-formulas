package main

import (
	"coffee/pkg/coffee"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	coffee.GiveMeSomeCoffee(loadInputs())
}

func loadInputs() coffee.Inputs {
	delivery, _ := strconv.ParseBool(os.Getenv("DELIVERY"))
	name := os.Getenv("NAME")
	coffeeType := os.Getenv("COFFEE_TYPE")

	outputDir := os.Getenv("RIT_OUTPUT_DIR")
	ioutil.WriteFile(outputDir+"/v1", []byte("a"), 0755)
	ioutil.WriteFile(outputDir+"/v2", []byte("b"), 0755)
	ioutil.WriteFile(outputDir+"/v3", []byte("c"), 0755)

	return coffee.Inputs{
		Name:       name,
		CoffeeType: coffeeType,
		Delivery:   delivery,
	}
}
