package i18n

import (
	"regexp"
	"strings"
)

// ExtractFinder struct
type ExtractFinder struct {
	Regex string
	Use   []int64
}

// Extract new translations terms from content
func Extract(content *string, context *Context, finder ExtractFinder) error {

	regex := finder.Regex
	regex = strings.ReplaceAll(regex, `:string`, `\s?[\'\"](.+?)[\'\"]\s?`)
	regex = strings.ReplaceAll(regex, `:var`, `\s?[\'\"]?(.+?)[\'\"]?\s?`)

	r, err := regexp.Compile(regex)

	if err != nil {
		return err
	}

	matches := r.FindAllStringSubmatch(*content, -1)
	terms := Terms{}
	existing := context.Terms

	for _, match := range matches {
		for _, use := range finder.Use {

			term := match[use]

			if term == "" {
				continue
			}

			if _, ok := terms[term]; !ok {
				if translated, ok := existing[term]; ok {
					terms[term] = translated
				} else {
					terms[term] = ""
				}
			}

		}
	}

	context.Terms = terms

	return nil
}
