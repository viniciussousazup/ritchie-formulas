package main

import (
    "os"
	"formula/pkg/formula"
)

func main() {
	input1 := os.Getenv("SAMPLE_TEXT")
	input2 := os.Getenv("SAMPLE_LIST")
	input3 := os.Getenv("SAMPLE_BOOL")

	formula.Input{
    	Text:    input1,
    	List:    input2,
    	Boolean: input3,
    }.Run()
}