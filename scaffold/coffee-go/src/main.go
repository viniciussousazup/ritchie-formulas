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
	output := os.Getenv("RIT_OUTPUT_FILE")
	ioutil.WriteFile(output, []byte("V1=a;V2=b;V3=c;V4=d"), 0755)

	return coffee.Inputs{
		Name:       name,
		CoffeeType: coffeeType,
		Delivery:   delivery,
	}
}
