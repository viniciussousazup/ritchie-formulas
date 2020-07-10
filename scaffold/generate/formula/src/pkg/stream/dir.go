package stream

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type DirManager interface {
	Create(dir string) error
	Remove(dir string) error
	List(dir string, hiddenDir bool) ([]string, error)
	Exists(dir string) bool
	IsDir(dir string) bool
	Copy(src, dst string) error
}

type DefaultDirManager struct {
	file FileManager
}

func NewDirManager(file FileManager) DirManager {
	return DefaultDirManager{file: file}
}

// Create creates a directory named dir
// A successful call returns err == nil
func (m DefaultDirManager) Create(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// Remove removes dir and any children it contains.
func (m DefaultDirManager) Remove(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return nil
}

func (m DefaultDirManager) Exists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	return true
}

func (m DefaultDirManager) IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return info.IsDir()
}

// List list all directories associate with dir name, case hiddenDir is
// true, we will add hidden directories to your slice of directories.
// If success, List returns a slice of directories and a nil error.
// If error, List returns an empty slice and a non-nil error.
func (m DefaultDirManager) List(dir string, hiddenDir bool) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, f := range fileInfos {
		n := f.Name()
		if f.IsDir() {
			if hiddenDir {
				dirs = append(dirs, n)
				continue
			}

			if !strings.ContainsAny(n, ".") {
				dirs = append(dirs, n)
			}
		}
	}

	return dirs, nil
}

func (m DefaultDirManager) Copy(src, dest string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			if err := m.Create(destPath); err != nil {
				return err
			}
			if err := m.Copy(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			if err := m.file.Copy(sourcePath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}
