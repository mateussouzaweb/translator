package i18n

import (
	"bytes"
	"io/ioutil"
)

// ReadFile retrieve file content from filepath
func ReadFile(filepath string) (string, error) {

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

// ReadFiles retrieve files content from file list
func ReadFiles(files []string) (string, error) {

	buf := bytes.NewBuffer(nil)

	for _, filepath := range files {

		content, err := ReadFile(filepath)

		if err != nil {
			return "", err
		}

		buf.WriteString(content)

	}

	return buf.String(), nil
}
