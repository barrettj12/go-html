# go-html

Write your webpage using nested Go structs, and render it to HTML.

```go
html.HTML(
	html.Head(
		html.Title("Page Title"),
	),
	html.Body(
		html.H1("My First Heading"),
		html.P("My first paragraph."),
	),
)
```
renders to
```html
<html>
	<head>
		<title>Page Title</title>
	</head>
	<body>
		<h1>My First Heading</h1>
		<p>My first paragraph.</p>	
	</body>
</html>
```