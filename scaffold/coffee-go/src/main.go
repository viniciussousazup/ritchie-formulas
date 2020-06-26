package main

import (
	"coffee/pkg/coffee"
	"fmt"
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
	fmt.Println("Teste Teste ")
	ioutil.WriteFile(output, []byte("V1=a;V2=b;V3=c;V4=d"), 0755)
	os.Setenv("RIT_OUTPUT_FILE", "teste1")
	fmt.Printf("f:%s=%s\n","RIT_OUTPUT_FILE",os.Getenv("RIT_OUTPUT_FILE"))
	os.Setenv("TESTE_ENV", "teste2")
	fmt.Printf("f:%s=%s\n","TESTE_ENV",os.Getenv("TESTE_ENV"))

	return coffee.Inputs{
		Name:       name,
		CoffeeType: coffeeType,
		Delivery:   delivery,
	}
}
