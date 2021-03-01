# Translator - Translation Extraction to JSON i18n files

*Translator* extracts all translatable terms of your project and places everything into a single JSON file that can be read from almost any programing language - including JavaScript.

It was originally designed to use in conjunction with multi-language SPA projects but is extensible to any kind of project that you are developing, whatever the language.

---

## Features

- Extractor written in Go Language
- Extracts translations from ``__()`` and ``translate()`` functions
- Supports plurals from ``__p()`` and ``pluralize()`` functions
- Reads translations from ``.js`` and ``.html`` files by default, but can include any other file extension
- Keeps already translated terms on the destination file
- Remove old translations that are not used anymore
- Just works!

---

## Installation and Usage

To install, just download the binary file and place it on the binaries folder:

```bash
sudo wget https://raw.githubusercontent.com/mateussouzaweb/translator/master/bin/translator -O /usr/local/bin/translator && sudo chmod +x /usr/local/bin/translator
```

To check command flags use:

```bash
translator --help
```

To translate a project, run:

```bash
translator \
    --source /path/to/source/ \
    --destination /path/to/destination/json/file.json
```

Now translate the terms by editing the JSON file. Enjoy!
