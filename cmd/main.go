package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mateussouzaweb/translator/i18n"
)

func main() {

	// Command line flags
	source := flag.String("source", "", "Path to the source files")
	extensions := flag.String("extensions", ".go,.html,.js", "Source file extensions. Comma separated")
	destination := flag.String("destination", "", "Path to the destination JSON file")
	version := flag.Bool("version", false, "Print program version")

	// Parse values
	flag.Parse()

	if *version {
		fmt.Println("Translator version 0.0.1")
		return
	}

	// Instantiate existing translations
	context := i18n.Context{
		File: *destination,
	}

	err := context.Load()

	if err != nil {
		panic(err)
	}

	// Retrieve source files content
	extensionsList := strings.Split(*extensions, ",")
	files, err := i18n.GetFiles(*source, extensionsList)

	if err != nil {
		panic(err)
	}

	content, err := i18n.ReadContent(files)

	if err != nil {
		panic(err)
	}

	// Extract new translations to context
	err = i18n.Extract(&content, &context, i18n.ExtractFinder{
		Regex: `(__|translate)\(:string\)`,
		Use:   []int64{2},
	})

	if err != nil {
		panic(err)
	}

	err = i18n.Extract(&content, &context, i18n.ExtractFinder{
		Regex: `(__p|pluralize)\(:var,:string,:string(,:string)?\)`,
		Use:   []int64{3, 4, 5},
	})

	if err != nil {
		panic(err)
	}

	// Write new translations result
	err = context.Write()

	if err != nil {
		panic(err)
	}

}
