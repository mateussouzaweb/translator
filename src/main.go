package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Retrieve files from source path
func getFiles(root string, extensions []string) ([]string, error) {

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

// Retrieve file content from filepath
func getFileContent(filepath string) (string, error) {

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

// Retrieve files content from file list
func getFilesContent(files []string) (string, error) {

	buf := bytes.NewBuffer(nil)

	for _, filepath := range files {

		content, err := getFileContent(filepath)

		if err != nil {
			return "", err
		}

		buf.WriteString(content)

	}

	return buf.String(), nil
}

// TranslationsMap type
type TranslationsMap = map[string]string

// TranslationFinder struct
type TranslationFinder struct {
	Regex string
	Use   []int64
}

// Load translations from file
func loadTranslations(file string, translations TranslationsMap) error {

	content, err := getFileContent(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(content), &translations)

	return err
}

// Extract new translations terms from content
func extractTranslations(content string, translations TranslationsMap, existing TranslationsMap, finder TranslationFinder) error {

	regex := finder.Regex
	regex = strings.ReplaceAll(regex, `:string`, `\s?[\'\"](.+?)[\'\"]\s?`)
	regex = strings.ReplaceAll(regex, `:var`, `\s?[\'\"]?(.+?)[\'\"]?\s?`)

	r, err := regexp.Compile(regex)

	if err != nil {
		return err
	}

	matches := r.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		for _, use := range finder.Use {

			term := match[use]

			if term == "" {
				continue
			}

			if _, ok := translations[term]; !ok {
				if translated, ok := existing[term]; ok {
					translations[term] = translated
				} else {
					translations[term] = ""
				}
			}

		}
	}

	return nil
}

// Save translations into file
func saveTranslations(translations TranslationsMap, file string) error {

	content, err := json.MarshalIndent(translations, "", "    ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, content, 0644)

	return err
}

func main() {

	// Command line flags
	source := flag.String("source", "", "Path to the source files")
	extensions := flag.String("extensions", ".html,.js", "Source file extensions. Comma separated")
	destination := flag.String("destination", "", "Path to the destination JSON file")

	// Parse values
	flag.Parse()

	// Instantiate translations
	translations := TranslationsMap{}
	existing := TranslationsMap{}

	err := loadTranslations(*destination, existing)

	if err != nil {
		panic(err)
	}

	// Retrieve new translations
	extensionsList := strings.Split(*extensions, ",")
	files, err := getFiles(*source, extensionsList)

	if err != nil {
		panic(err)
	}

	content, err := getFilesContent(files)

	if err != nil {
		panic(err)
	}

	err = extractTranslations(content, translations, existing, TranslationFinder{
		Regex: `(__|translate)\(:string\)`,
		Use:   []int64{2},
	})

	if err != nil {
		panic(err)
	}

	err = extractTranslations(content, translations, existing, TranslationFinder{
		Regex: `(__p|pluralize)\(:var,:string,:string(,:string)?\)`,
		Use:   []int64{3, 4, 5},
	})

	if err != nil {
		panic(err)
	}

	// Save final result
	err = saveTranslations(translations, *destination)

	if err != nil {
		panic(err)
	}

}
