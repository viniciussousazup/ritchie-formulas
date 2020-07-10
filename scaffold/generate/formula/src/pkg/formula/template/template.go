package template

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	templateDirName = "templates"
)

var (
	ErrTemplateDirNotFound = errors.New("TemplateDir not found")
)

type Manager interface {
	TemplateDir() (string, error)
	Languages() ([]string, error)
	LangTemplateFiles(lang string) ([]File, error)
	ResolverNewPath(oldPath, newDir, lang string) (string, error)
}

type File struct {
	Path  string
	IsDir bool
}

func NewManager() Manager {
	return DefaultManager{}
}

func NewManagerCustom(templateDir string) Manager {
	return DefaultManager{templateDir}
}

type DefaultManager struct {
	templateDir string
}

func (tm DefaultManager) ResolverNewPath(oldPath, newDir, lang string) (string, error) {
	tplD, err := tm.TemplateDir()
	if err != nil {
		return "", err
	}
	oldDir := path.Join(tplD, lang)
	return strings.Replace(oldPath, oldDir, newDir, 1), nil
}

func (tm DefaultManager) Languages() ([]string, error) {
	tplD, err := tm.TemplateDir()
	if err != nil {
		return nil, err
	}

	dirs, err := ioutil.ReadDir(tplD)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, d := range dirs {
		if d.IsDir() {
			result = append(result, d.Name())
		}
	}

	return result, nil
}

func (tm DefaultManager) TemplateDir() (string, error) {

	if tm.templateDir != "" {
		return tm.templateDir, nil
	}

	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir, _ := path.Split(executable)
	dir = path.Join(dir, "..")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, f := range files {
		if f.Name() == templateDirName && f.IsDir() {
			result := path.Join(dir, templateDirName)
			tm.templateDir = result
			return result, nil
		}
	}

	return "", ErrTemplateDirNotFound

}

func (tm DefaultManager) LangTemplateFiles(lang string) ([]File, error) {
	tplD, err := tm.TemplateDir()
	if err != nil {
		return nil, err
	}

	langDir := path.Join(tplD, lang)

	return readDirRecursive(langDir)

}

func readDirRecursive(dir string) ([]File, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fileNames []File
	for _, f := range files {
		if f.IsDir() {
			dirFiles, err := readDirRecursive(path.Join(dir, f.Name()))
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, dirFiles...)
		}
		fileNames = append(fileNames, File{
			Path:  path.Join(dir, f.Name()),
			IsDir: f.IsDir(),
		})

	}
	return fileNames, nil
}
