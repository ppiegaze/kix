package main

type Content struct {
	Items          []Item         `mapstructure:"items"`
	ModularContent map[string]any `mapstructure:"modular_content"`
	Pagination     Pagination     `mapstructure:"pagination"`
}

type Item struct {
	System   System         `mapstructure:"system"`
	Elements map[string]any `mapstructure:"elements"`
}

type Pagination struct {
	Skip     int    `mapstructure:"skip"`
	Limit    int    `mapstructure:"limit"`
	Count    int    `mapstructure:"count"`
	NextPage string `mapstructure:"next_page"`
}

type System struct {
	Id               string   `mapstructure:"id"`
	Name             string   `mapstructure:"name"`
	CodeName         string   `mapstructure:"codename"`
	Language         string   `mapstructure:"language"`
	Type             string   `mapstructure:"type"`
	Collection       string   `mapstructure:"collection"`
	SitemapLocations []string `mapstructure:"sitemap_locations"`
	LastModified     string   `mapstructure:"last_modified"`
	WokflowStep      string   `mapstructure:"workflow_step"`
}

type Article struct {
	Title           SimpleElement   `mapstructure:"title"`
	BodyCopy        RichTextElement `mapstructure:"body_copy"`
	RelatedArticles SimpleElement   `mapstructure:"related_articles"`
	ElementAuthor   SimpleElement   `mapstructure:"author"`
	URL             SimpleElement   `mapstructure:"url"`
}

type NavigationItem struct {
	Title    SimpleElement `mapstructure:"title"`
	URL      SimpleElement `mapstructure:"url"`
	SubItems SimpleElement `mapstructure:"subitems"`
}

type Author struct {
	Name SimpleElement   `mapstructure:"name"`
	Bio  RichTextElement `mapstructure:"bio"`
}

type SimpleElement struct {
	Type  string `mapstructure:"type"`
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

type RichTextElement struct {
	Type           string `mapstructure:"type"`
	Name           string `mapstructure:"name"`
	ModularContent string `mapstructure:"modular_content"`
	Value          string `mapstructure:"value"`
}
