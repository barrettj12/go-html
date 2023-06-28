package html_test

import (
	"strings"
	"testing"

	"github.com/barrettj12/go-html"
	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	expected := `
<html>
	<head>
		<title>Page Title</title>
	</head>
	<body>
		<h1>My First Heading</h1>
		<p>My first paragraph.</p>	
	</body>
</html>`

	doc :=
		html.HTML(
			html.Head(
				html.Title("Page Title"),
			),
			html.Body(
				html.H1("My First Heading"),
				html.P("My first paragraph."),
			),
		)

	rendered := doc.Render()
	assert.Equal(t, rendered, deformat(expected))
}

// remove indentation from HTML string
func deformat(raw string) string {
	var stripped string
	for _, line := range strings.Split(raw, "\n") {
		stripped += strings.TrimSpace(line)
	}
	return stripped
}
