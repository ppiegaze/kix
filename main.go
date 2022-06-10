package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const in = "content.json"
const outdir = "out"

func main() {

	// Read the JSON content dump from Kentico
	fi, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	// Create output directory
	outpath := filepath.Join(".", outdir)
	err = os.MkdirAll(outpath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON to a map[string]any (map of string to any).
	// We use this approach instead of defining structs with custom field tags
	// because the source JSON contains a mixture of known and unknown
	// property names.
	var content map[string]any
	err = json.Unmarshal(fi, &content)
	if err != nil {
		log.Fatal(err)
	}

	// Convert JSON to Markdown
	err = convert(content)
	if err != nil {
		log.Fatal(err)
	}
}
