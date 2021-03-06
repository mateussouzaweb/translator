package i18n

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ExtractFinder struct
type ExtractFinder struct {
	Format string
	Use    []int64
}

// ExtractFiles retrieve files from source path
func ExtractFiles(root string, extensions []string) ([]string, error) {

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

// ExtractTerms from regex
func ExtractTerms(regex string, use []int64, content *string, toContext *Context) error {

	r, err := regexp.Compile(regex)

	if err != nil {
		return err
	}

	matches := r.FindAllStringSubmatch(*content, -1)

	for _, match := range matches {
		for _, use := range use {

			term := match[use]

			if term == "" {
				continue
			}

			if _, ok := toContext.Terms[term]; !ok {
				toContext.Terms[term] = ""
			}

		}
	}

	return nil
}

// Extract new terms from content with finder rules
func Extract(content *string, toContext *Context, finder ExtractFinder) error {

	regex := "(?i)" + finder.Format
	regex = strings.ReplaceAll(regex, `:string`, `\s?\'(.+?)\'\s?`)
	regex = strings.ReplaceAll(regex, `:var`, `\s?\'?(.+?)\'?\s?`)

	err := ExtractTerms(regex, finder.Use, content, toContext)

	if err != nil {
		return err
	}

	regex = "(?i)" + finder.Format
	regex = strings.ReplaceAll(regex, `:string`, `\s?\"(.+?)\"\s?`)
	regex = strings.ReplaceAll(regex, `:var`, `\s?\"?(.+?)\"?\s?`)

	err = ExtractTerms(regex, finder.Use, content, toContext)

	return err
}

// Merge terms from one context into another
// Also can remove old terms from target context if not found on source context
func Merge(context *Context, toContext *Context, removeNotFound bool) {

	for term, translation := range context.Terms {
		if _, ok := toContext.Terms[term]; !ok {
			toContext.Terms[term] = translation
		}
	}

	if removeNotFound {
		for index := range toContext.Terms {
			if _, ok := context.Terms[index]; !ok {
				delete(toContext.Terms, index)
			}
		}
	}

}
