package utils

import (
	"io"
	"os"
	"path/filepath"
)

func Copy(source, target string) (written int64, err error) {

	// source file
	file1, err := os.Open(source)
	if err != nil {
		return
	}

	// create target file

	if _, err := os.Stat(filepath.Dir(target)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(target), os.ModeDir|0755)
	}
	file2, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}

	defer file1.Close()
	defer file2.Close()
	written, err = io.Copy(file2, file1)
	if err != nil {
		return
	}
	return
}
