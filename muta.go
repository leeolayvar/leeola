package main

import (
	"errors"
	"fmt"

	"./.muta/plugins/prettyurls"
	"github.com/leeola/muta"
	"github.com/leeola/muta-frontmatter"
	"github.com/leeola/muta-markdown"
	"github.com/leeola/muta-sass"
	"github.com/leeola/muta-template"
)

type BlogPost struct {
	Title    string `yaml:"title"`
	PageName string `yaml:"pageName"`
}

type Page struct {
	Title    string `yaml:"title"`
	PageName string `yaml:"pageName"`
}

func fmTyper(st string) (t interface{}, err error) {
	switch st {
	case "blogpost":
		t = &BlogPost{}
	case "page":
		t = &Page{}
	default:
		err = errors.New(fmt.Sprintf("Unknown FrontMatter type '%s'", st))
	}
	return
}

func Markdown() (*muta.Stream, error) {
	s := muta.Src("./*.md").
		Pipe(frontmatter.FrontMatter(fmTyper)).
		Pipe(markdown.Markdown()).
		Pipe(template.Template("./.muta/templates")).
		Pipe(prettyurls.Prettyurls()).
		Pipe(muta.Dest("./build"))
	return s, nil
}

func Sass() (*muta.Stream, error) {
	s := muta.Src("./.muta/sass/*.scss").
		Pipe(sass.Sass()).
		Pipe(muta.Dest("./build/css"))
	return s, nil
}

func main() {
	muta.Task("markdown", Markdown)
	muta.Task("sass", Sass)

	muta.Task("build", "markdown", "sass")
	muta.Task("default", "build")
	muta.Te()
}