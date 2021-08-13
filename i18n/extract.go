package i18n

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

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

// ExtractTerms from final regex
func ExtractTerms(content *string, toContext *Context, regex *regexp.Regexp) error {

	matches := regex.FindAllStringSubmatch(*content, -1)

	for _, match := range matches {
		for index, term := range match {

			if index == 0 || term == "" {
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
func Extract(content *string, toContext *Context, pattern string) error {

	regex := "(?i)" + pattern
	regex = strings.ReplaceAll(regex, `:term`, `\s?\'(.+?)\'\s?`)
	regex = strings.ReplaceAll(regex, `:variable`, `\s?\'?(?:.+?)\'?\s?`)
	compiled, err := regexp.Compile(regex)

	if err != nil {
		return err
	}

	err = ExtractTerms(content, toContext, compiled)

	if err != nil {
		return err
	}

	regex = "(?i)" + pattern
	regex = strings.ReplaceAll(regex, `:term`, `\s?\"(.+?)\"\s?`)
	regex = strings.ReplaceAll(regex, `:variable`, `\s?\"?(?:.+?)\"?\s?`)
	compiled, err = regexp.Compile(regex)

	if err != nil {
		return err
	}

	err = ExtractTerms(content, toContext, compiled)

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
