# Translator - Translation extraction from any code to JSON files

*Translator* extracts all translatable terms of your project and places everything into a single JSON file that can be read from almost any programing language - including JavaScript.

It was originally designed to use in conjunction with multi-language SPA projects but is extensible to any kind of project that you are developing, whatever the source code language - if can reads JSON, will works.

----

## Features

- Written in Go Language and provided as package module
- Extracts terms from ``__()`` and ``translate()`` functions
- Supports plurals from ``__p()`` and ``pluralize()`` functions
- Supports terms extraction from custom patterns
- Reads terms from the following files by default: ``.html``, ``.js``, ``.jsx``, ``.ts`` ``.tsx``, ``.mjs``, ``.mts`` and ``.go``
- You can include any other file extension to extract translations
- Keeps already translated terms on the destination file
- Remove old terms that are not used anymore
- Can be use to extract and translate Go projects on the fly
- Just works!

----

## CLI - Installation and Usage

To install, just download the most recent binary file:

```bash
REPOSITORY="https://github.com/mateussouzaweb/translator/releases/download/latest"
sudo wget $REPOSITORY/translator -O /usr/local/bin/translator
sudo chmod +x /usr/local/bin/translator
```

To check command flags use:

```bash
translator --help
```

To translate a project, run:

```bash
translator --source ./ --destination ./translations.json
```

Now translate the terms by editing the JSON file. Enjoy!

----

## Usage with Docker

You can run translator directly from a Docker container image. Simply pull the published image from the registry:

```bash
# See help
docker run --rm ghcr.io/mateussouzaweb/translator:latest --help

# Run current folder on statiq
docker run --rm -v "$PWD:/app" \
  ghcr.io/mateussouzaweb/translator:latest \
    --source /app \
    --destination /app/translations.json
```

---

## Usage In Go Projects

Start by requiring the go module:

```bash
go get github.com/mateussouzaweb/translator
```

Load the translation file and start using it:

```go
package main

import (
    "fmt"
    "log"
    "github.com/mateussouzaweb/translator/i18n"
)

translator := i18n.Context{
    Code:  "pt-BR",
    Alias: "pt",
    File:  "translations/pt-BR.json",
}

err := translator.Load()

if err != nil {
    log.Fatal(err)
}

fmt.Println( translator.Translate("Hello World!") )
// Result: Olá Mundo!
```

Enjoy!
