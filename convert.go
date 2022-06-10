package main

import (
	"fmt"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/mitchellh/mapstructure"
)

const (
	primaryContent = "primary"
	modularContent = "modular"
)

// Convert map to Markdown
func convert(contentMap map[string]any) error {

	// Decode the map into a Content struct
	var content Content
	err := mapstructure.Decode(contentMap, &content)
	if err != nil {
		return err
	}

	// We know that Content.Items is a []Item
	// Depending on the Item type we handle the conversion differently
	for _, item := range content.Items {
		switch item.System.Type {
		case "article":
			convertArticle(item, primaryContent)
		case "navigation_item":
			convertNavigationItem(item, primaryContent)
		case "author":
			convertAuthor(item, primaryContent)
		}
	}

	// We know that the Content.ModuleContent is a map[string]any
	// Each value in that map is convertible to an Item
	// We iterate and convert.
	// Depending on the Item type we handle the conversion differently
	for _, m := range content.ModularContent {

		var item Item
		err := mapstructure.Decode(m, &item)
		if err != nil {
			return err
		}

		switch item.System.Type {
		case "article":
			convertArticle(item, modularContent)
		case "navigation_item":
			convertNavigationItem(item, modularContent)
		case "author":
			convertAuthor(item, modularContent)
		}
	}
	return nil
}

func convertArticle(item Item, contentType string) error {
	// We know this Item contains an Article
	var article Article
	err := mapstructure.Decode(item.Elements, &article)
	if err != nil {
		return err
	}

	// Construct file name
	var fn = outdir + "/" + item.System.CodeName + "-" + contentType + ".md"

	// Open output file
	fo, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fo.Close()

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(article.BodyCopy.Value)
	if err != nil {
		return err
	}

	fmt.Fprintln(fo, "Name: "+item.System.Name)
	fmt.Fprintln(fo, "Title: "+article.Title.Value)
	fmt.Fprintln(fo, "URL Slug: "+article.URL.Value)
	fmt.Fprintln(fo, "Author: "+article.ElementAuthor.Value)
	fmt.Fprintln(fo, "Related: "+item.System.LastModified)
	fmt.Fprintln(fo, "Type: "+item.System.Type)
	fmt.Fprintln(fo, "Last Modified: "+item.System.LastModified)
	fmt.Fprintln(fo, "------------------")
	fmt.Fprintln(fo, "")
	fmt.Fprintln(fo, "# "+article.Title.Value)
	fmt.Fprint(fo, markdown)
	return nil
}

func convertNavigationItem(item Item, contentType string) error {
	// We know this Item contains a NavigationItem
	var navigationItem NavigationItem
	err := mapstructure.Decode(item.Elements, &navigationItem)
	if err != nil {
		return err
	}

	// Construct file name
	var fn = outdir + "/" + item.System.CodeName + "-" + contentType + ".md"

	// Open output file
	fo, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fo.Close()

	fmt.Fprintln(fo, "Name: "+item.System.Name)
	fmt.Fprintln(fo, "Title: "+navigationItem.Title.Value)
	fmt.Fprintln(fo, "URL Slug: "+navigationItem.URL.Value)
	fmt.Fprintln(fo, "Subitems: "+navigationItem.SubItems.Value)
	fmt.Fprintln(fo, "Type: "+item.System.Type)
	fmt.Fprintln(fo, "Last Modified: "+item.System.LastModified)
	fmt.Fprintln(fo, "------------------")

	return nil
}

func convertAuthor(item Item, contentType string) error {
	// We know this Item contains an Author
	var author Author
	err := mapstructure.Decode(item.Elements, &author)
	if err != nil {
		return err
	}

	// Construct file name
	var fn = outdir + "/" + item.System.CodeName + "-" + contentType + ".md"

	// Open output file
	fo, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fo.Close()

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(author.Bio.Value)
	if err != nil {
		return err
	}

	fmt.Fprintln(fo, "Name: "+author.Name.Value)
	fmt.Fprintln(fo, "Type: "+item.System.Type)
	fmt.Fprintln(fo, "Last Modified: "+item.System.LastModified)
	fmt.Fprintln(fo, "------------------")
	fmt.Fprintln(fo, "")
	fmt.Fprintln(fo, "# "+author.Bio.Name)
	fmt.Fprint(fo, markdown)

	return nil
}
