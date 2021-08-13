package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mateussouzaweb/translator/i18n"
)

func main() {

	var patterns []string

	patterns = append(patterns,
		`(?:__|translate)\(:term\)`,
		`(?:translate\s):term`,
		`(?:__p|pluralize)\(:variable,:term,:term(,:term)?\)`,
	)

	// Command line flags
	version := flag.Bool(
		"version",
		false,
		"Print program version",
	)

	source := flag.String(
		"source",
		"",
		"Path to the source files",
	)

	destination := flag.String(
		"destination",
		"",
		"Path to the destination JSON file",
	)

	extensions := flag.String(
		"extensions",
		".go,.html,.js", "Source file extensions. Comma separated",
	)

	remove := flag.Bool(
		"remove",
		true,
		"Remove not found terms",
	)

	flag.Func(
		"add-pattern",
		"Declare additional regex pattern to extract translations.\nUse :term to represent translatable terms and :variable to represent non translatable variables.\nIf you use regex group, make sure to use non capturing group, as every capturing group is considered a translation term",
		func(value string) error {
			patterns = append(patterns, value)
			return nil
		})

	// Parse values
	flag.Parse()

	if *version {
		fmt.Println("Translator version 0.0.5")
		return
	}

	// Instantiate existing terms from file
	destinationContext := i18n.Context{
		File: *destination,
	}

	err := destinationContext.Load()

	if err != nil {
		panic(err)
	}

	// Retrieve source files content
	extensionsList := strings.Split(*extensions, ",")
	files, err := i18n.ExtractFiles(*source, extensionsList)

	if err != nil {
		panic(err)
	}

	content, err := i18n.ReadFiles(files)

	if err != nil {
		panic(err)
	}

	// Extract new terms to context
	context := i18n.Context{
		Terms: i18n.Terms{},
	}

	for _, pattern := range patterns {

		err = i18n.Extract(&content, &context, pattern)

		if err != nil {
			panic(err)
		}

	}

	// Merge and remove not found again terms
	i18n.Merge(&context, &destinationContext, *remove)

	// Write new terms result
	err = destinationContext.Write()

	if err != nil {
		panic(err)
	}

}
