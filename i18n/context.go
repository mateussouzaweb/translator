package i18n

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Placeholders struct
type Placeholders = map[string]string

// Terms struct
type Terms = map[string]string

// Context struct
type Context struct {
	Code  string
	Alias string
	File  string
	Terms Terms
}

// Translate method return the translated term.
// If the translation not exists, return the current term
func (c *Context) Translate(term string) string {

	if translated, ok := c.Terms[term]; ok {
		return translated
	}

	return term
}

// Pluralize detect correct term and return its translation
func (c *Context) Pluralize(count int64, singular string, plural string, zero string) string {

	term := singular
	if count == 0 {
		term = zero
	} else if count > 1 {
		term = plural
	}

	return c.Translate(term)
}

// Replace placeholders with real value
func (c *Context) Replace(translation string, placeholders Placeholders) string {

	for placeholder, value := range placeholders {
		translation = strings.ReplaceAll(translation, placeholder, value)
	}

	return translation
}

// Load translations from file
func (c *Context) Load() error {

	content, err := ReadFile(c.File)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(content), &c.Terms)

	return err
}

// Write terms into file
func (c *Context) Write() error {

	content, err := json.MarshalIndent(c.Terms, "", "    ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.File, content, 0644)

	return err
}
