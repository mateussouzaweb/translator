package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mateussouzaweb/translator/i18n"
)

func main() {

	// Command line flags
	version := flag.Bool("version", false, "Print program version")
	source := flag.String("source", "", "Path to the source files")
	destination := flag.String("destination", "", "Path to the destination JSON file")
	extensions := flag.String("extensions", ".go,.html,.js", "Source file extensions. Comma separated")
	remove := flag.Bool("remove", true, "Remove not found terms")

	// Parse values
	flag.Parse()

	if *version {
		fmt.Println("Translator version 0.0.4")
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

	err = i18n.Extract(&content, &context, i18n.ExtractFinder{
		Format: `(__|translate)\(:string\)`,
		Use:    []int64{2},
	})

	if err != nil {
		panic(err)
	}

	err = i18n.Extract(&content, &context, i18n.ExtractFinder{
		Format: `(translate\s):string`,
		Use:    []int64{2},
	})

	if err != nil {
		panic(err)
	}

	err = i18n.Extract(&content, &context, i18n.ExtractFinder{
		Format: `(__p|pluralize)\(:var,:string,:string(,:string)?\)`,
		Use:    []int64{3, 4, 5},
	})

	if err != nil {
		panic(err)
	}

	// Merge and remove not found again terms
	i18n.Merge(&context, &destinationContext, *remove)

	// Write new terms result
	err = destinationContext.Write()

	if err != nil {
		panic(err)
	}

}
