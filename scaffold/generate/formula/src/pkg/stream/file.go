package stream

import (
	"io"
	"io/ioutil"
	"os"
)

type FileManager interface {
	Exists(path string) bool
	Read(path string) ([]byte, error)
	Write(path string, content []byte) error
	Create(path string, data io.ReadCloser) error
	Remove(path string) error
	List(file string) ([]string, error)
	Copy(src, dst string) error
}

// DefaultFileManager implements FileManager
type DefaultFileManager struct {
}

// NewFileManager returns a FileManage that writes from w
// reads from r, exists from e and removes from re
func NewFileManager() FileManager {
	return DefaultFileManager{}
}

// Read reads the file named by path and returns the contents.
// A successful call returns err == nil
func (f DefaultFileManager) Read(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return b, err
}

// Exists returns true if file path exists
func (f DefaultFileManager) Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Write writes content to a file named by path.
// A successful call returns err == nil
func (f DefaultFileManager) Write(path string, content []byte) error {
	return ioutil.WriteFile(path, content, os.ModePerm)
}

func (f DefaultFileManager) Create(path string, data io.ReadCloser) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	if _, err = io.Copy(file, data); err != nil {
		return err
	}

	return nil
}

// Remove removes the named file
func (f DefaultFileManager) Remove(path string) error {
	if f.Exists(path) {
		return os.Remove(path)
	}
	return nil
}

// List lists all files in dir path
func (f DefaultFileManager) List(dir string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, f := range fileInfos {
		n := f.Name()
		if !f.IsDir() {
			dirs = append(dirs, n)
		}
	}

	return dirs, nil
}

func (f DefaultFileManager) Copy(src, dest string) error {
	input, err := f.Read(src)
	if err != nil {
		return err
	}

	if err := f.Write(dest, input); err != nil {
		return err
	}

	return nil
}
