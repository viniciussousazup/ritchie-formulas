package formula

import (
	"fmt"
	"github.com/fatih/color"
)

type Input struct {
	Text string
	List string
	Boolean string
}

func(in Input)Run()  {
	fmt.Println("Hello world!")
	color.Green(fmt.Sprintf("You receive %s in text.", in.Text ))
	color.Red(fmt.Sprintf("You receive %s in list.", in.List ))
	color.Yellow(fmt.Sprintf("You receive %s in boolean.", in.Boolean ))
}