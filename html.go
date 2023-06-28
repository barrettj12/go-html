package html

import (
	"fmt"
)

// GlobalAttributes is the set of attributes which can be used with all HTML
// elements.
// https://www.w3schools.com/tags/ref_standardattributes.asp
type GlobalAttributes struct{}

// ELEMENTS

// Element is the interface implemented by all HTML elements.
type Element interface {
	Render() string
}

// Tag represents a generic HTML tag.
type Tag struct {
	name       string
	attributes map[string]string
	contents   []Element
}

func (t *Tag) Render() string {
	var rendered string

	// Opening tag
	rendered += "<"
	rendered += t.name
	for k, v := range t.attributes {
		rendered += fmt.Sprintf(` %s="%s"`, k, v)
	}
	rendered += ">"

	// Render contents
	for _, c := range t.contents {
		rendered += c.Render()
	}

	// Closing tag
	rendered += "</"
	rendered += t.name
	rendered += ">"

	return rendered
}

func HTML(contents ...Element) Element {
	return &Tag{
		name:     "html",
		contents: contents,
	}
}

func Head(contents ...Element) Element {
	return &Tag{
		name:     "head",
		contents: contents,
	}
}

func Body(contents ...Element) Element {
	return &Tag{
		name:     "body",
		contents: contents,
	}
}

func Title(title string) Element {
	return &Tag{
		name:     "title",
		contents: []Element{String(title)},
	}
}

func H1(heading string) Element {
	return &Tag{
		name:     "h1",
		contents: []Element{String(heading)},
	}
}

func P(para string) Element {
	return &Tag{
		name:     "p",
		contents: []Element{String(para)},
	}
}

// String represents a string as an HTML value
type String string

func (s String) Render() string {
	return string(s)
}
