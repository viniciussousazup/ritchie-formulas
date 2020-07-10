package creator

import (
	"fmt"
	"formula/pkg/formula"
	"formula/pkg/formula/template"
	"formula/pkg/stream"
	"path"
	"strings"
)

const (
	helpFileTemplate = "help placeholder for %s"
)

type CreateManager struct {
	tplM template.Manager
	dir  stream.DirManager
	file stream.FileManager
}

func NewCreator() CreateManager {
	return CreateManager{}
}

func (c CreateManager) Create(cf formula.Create) error {

	formulaPath := cf.FormulaPath()
	formulaCmdName := cf.FormulaCmdName()

	err := c.dir.Create(formulaPath)
	if err != nil {
		return err
	}

	if err := c.createHelpFiles(formulaCmdName, cf.OutPutDir); err != nil {
		return err
	}

	if err := c.applyLangTemplate(formulaPath, cf.Lang, cf.OutPutDir); err != nil {
		return err
	}

	return nil
}

func (c CreateManager) createHelpFiles(formulaCmdName, outPutDir string) error {
	dirs := strings.Split(formulaCmdName, " ")
	for i := 0; i < len(dirs); i++ {
		d := dirs[0 : i+1]
		tPath := path.Join(outPutDir, path.Join(d...))
		helpPath := fmt.Sprintf("%s/help.txt", tPath)
		if !c.file.Exists(helpPath) {
			folderName := path.Base(tPath)
			tpl := fmt.Sprintf(helpFileTemplate, folderName)
			err := c.file.Write(helpPath, []byte(tpl))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c CreateManager) applyLangTemplate(lang, outPutDir, formulaPath string) error {

	tFiles, err := c.tplM.LangTemplateFiles(lang)
	if err != nil {
		return err
	}

	for _, f := range tFiles {
		if f.IsDir {
			newPath := resolverNewPath(f.Path, outPutDir, formulaPath)
			err := c.dir.Create(newPath)
			if err != nil {
				return err
			}
		} else {
			tpl, err := c.file.Read(f.Path)
			if err != nil {
				return err
			}
			newPath := resolverNewPath(f.Path, outPutDir, formulaPath)
			err = c.file.Write(newPath, tpl)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func resolverNewPath(oldPath, outPutDir, formulaPath string) string {
	return strings.Replace(oldPath, outPutDir, formulaPath, 1)
}
