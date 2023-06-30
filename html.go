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

func HTML(head *Head_, body *Body_) *HTML_ {
	return &HTML_{
		Head: head,
		Body: body,
	}
}

func Head(title string, meta ...Metadata) *Head_ {
	return &Head_{
		Title: title,
		Meta:  meta,
	}
}

func Body(contents ...Element) *Body_ {
	return &Body_{
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

// HTML_ is an <html> element.
type HTML_ struct {
	// TODO: add lang attribute
	Head *Head_
	Body *Body_
}

func (h *HTML_) Render() string {
	rendered := "<html>"
	if h.Head != nil {
		rendered += h.Head.Render()
	}
	if h.Body != nil {
		rendered += h.Body.Render()
	}
	rendered += "</html>"
	return rendered
}

// Head_ is a <head> element.
// TODO: do we really need this? It's kind of an XML hack which doesn't really
// make sense in Go.
type Head_ struct {
	Title string
	Meta  []Metadata
}

func (h *Head_) Render() string {
	rendered := "<head><title>"
	rendered += h.Title
	rendered += "</title>"
	for _, m := range h.Meta {
		rendered += m.Render()
	}
	rendered += "</head>"
	return rendered
}

// Body_ is a <body> element.
type Body_ struct {
	contents []Element
}

func (b *Body_) Render() string {
	rendered := "<body>"
	for _, c := range b.contents {
		rendered += c.Render()
	}
	rendered += "</body>"
	return rendered
}

// https://html.spec.whatwg.org/multipage/dom.html#metadata-content
// base link meta noscript script style template
type Metadata interface {
	Element
}
