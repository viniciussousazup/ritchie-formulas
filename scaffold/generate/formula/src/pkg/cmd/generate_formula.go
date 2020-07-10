package cmd

import (
	"errors"
	"formula/pkg/formula"
	"formula/pkg/formula/creator"
	"formula/pkg/formula/template"
	"formula/pkg/prompt"
	"strings"
)

var (
	ErrNotAllowedCharacter = prompt.NewError(`not allowed character on formula name \/,><@-`)
	ErrDontStartWithRit    = prompt.NewError("Rit formula's command needs to start with \"rit\" [ex.: rit group verb <noun>]")
	ErrTooShortCommand     = prompt.NewError("Rit formula's command needs at least 2 words following \"rit\" [ex.: rit group verb]")
)

const notAllowedChars = `\/><,@-`

type GenerateFormulaCmd struct {
	currentDir      string
	inTextValidator prompt.InputTextValidator
	inList          prompt.InputList
	inText          prompt.InputText
	inBool          prompt.InputBool
	tplM            template.Manager
}

func NewCreateFormulaCmd(
	currentDIr string,
	inTextValidator prompt.InputTextValidator,
	inList prompt.InputList,
	inBool prompt.InputBool,
	tplM template.Manager,
) GenerateFormulaCmd {
	return GenerateFormulaCmd{
		currentDir:      currentDIr,
		inTextValidator: inTextValidator,
		inList:          inList,
		inBool:          inBool,
		tplM:            tplM,
	}
}

func (c GenerateFormulaCmd) Run() error {
	formulaCmd, err := c.inTextValidator.Text(
		"Enter the new formula command: ",
		c.surveyCmdValidator,
		"You must create your command based in this example [rit group verb noun]",
	)
	if err != nil {
		return err
	}

	if strings.ContainsAny(formulaCmd, notAllowedChars) {
		return ErrNotAllowedCharacter
	}

	validLang, err := c.tplM.Languages()
	if err != nil {
		return err
	}
	lang, err := c.inList.List("Choose the language: ", validLang)
	if err != nil {
		return err
	}
	println(lang)

	useCurrentDir, err := c.inBool.Bool("Create formula on current dir ?", prompt.DefaultBoolOpts)
	if err != nil {
		return err
	}

	var outPutDir string
	if !useCurrentDir {
		outPutDir, err = c.inText.Text("Output path (e.g.: /home/user/github):", true)
		if err != nil {
			return err
		}
	} else {
		outPutDir = c.currentDir
	}

	err = creator.NewCreator().Create(formula.Create{
		FormulaCmd: formulaCmd,
		Lang:       lang,
		OutPutDir:  outPutDir,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c GenerateFormulaCmd) surveyCmdValidator(cmd interface{}) error {
	if len(strings.TrimSpace(cmd.(string))) < 1 {
		return errors.New("this input must not be empty")
	}

	s := strings.Split(cmd.(string), " ")
	if s[0] != "rit" {
		return ErrDontStartWithRit
	}

	if len(s) <= 2 {
		return ErrTooShortCommand
	}
	return nil
}
