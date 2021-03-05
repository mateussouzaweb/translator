package i18n

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetFiles retrieve files from source path
func GetFiles(root string, extensions []string) ([]string, error) {

	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		match := false

		for _, item := range extensions {
			if ext == item {
				match = true
				break
			}
		}

		if !match {
			return nil
		}

		files = append(files, path)

		return nil
	})

	return files, err
}

// ReadFileContent retrieve file content from filepath
func ReadFileContent(filepath string) (string, error) {

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

// ReadContent retrieve files content from file list
func ReadContent(files []string) (string, error) {

	buf := bytes.NewBuffer(nil)

	for _, filepath := range files {

		content, err := ReadFileContent(filepath)

		if err != nil {
			return "", err
		}

		buf.WriteString(content)

	}

	return buf.String(), nil
}
