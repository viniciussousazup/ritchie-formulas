package hello

import (
	"fmt"
	"github.com/fatih/color"
)

type Hello struct {
	Text    string
	List    string
	Boolean string
}

func (h Hello) Run() {
	fmt.Println("Hello world!")
	color.Green(fmt.Sprintf("You receive %s in text.", h.Text))
	color.Red(fmt.Sprintf("You receive %s in list.", h.List))
	color.Yellow(fmt.Sprintf("You receive %s in boolean.", h.Boolean))
}
