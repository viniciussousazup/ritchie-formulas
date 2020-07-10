package main

import (
	"formula/pkg/cmd"
	"formula/pkg/formula/template"
	"formula/pkg/prompt"
	"os"
)

func main() {

	currentDir := os.Getenv("CURRENT_PWD")

	// prompt
	inputTextValidator := prompt.NewSurveyTextValidator()
	inputList := prompt.NewSurveyList()
	inputBool := prompt.NewSurveyBool()
	templateManager := template.NewManager()

	generateFormulaCmd := cmd.NewCreateFormulaCmd(
		currentDir,
		inputTextValidator,
		inputList,
		inputBool,
		templateManager,
	)

	err := generateFormulaCmd.Run()
	if err != nil {
		panic(err)
	}

}
