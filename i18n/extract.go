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

// Extract new terms from content with finder rules
func Extract(content *string, toContext *Context, finder ExtractFinder) error {

	regex := finder.Regex
	regex = strings.ReplaceAll(regex, `:string`, `\s?[\'\"](.+?)[\'\"]\s?`)
	regex = strings.ReplaceAll(regex, `:var`, `\s?[\'\"]?(.+?)[\'\"]?\s?`)

	r, err := regexp.Compile(regex)

	if err != nil {
		return err
	}

	matches := r.FindAllStringSubmatch(*content, -1)

	for _, match := range matches {
		for _, use := range finder.Use {

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

// Merge terms from one context into another
// Also can remove old terms from target context if not found on source context
func Merge(context *Context, toContext *Context, removeNotFound bool) {

	for term, translation := range context.Terms {
		if _, ok := toContext.Terms[term]; !ok {
			context.Terms[term] = translation
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
