package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
	<body>
		<h1>My First Page</h1>
		<p>I am Uzama</p>
		<p>I am good</p>
		<img src="img.png" width="104" height="200">
		<p>I am good</p>
		<img src="img.png" width="104" height="200">
	</body>
</html>
`

// count img and words in html file
func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(-1)
	}

	words, pics := countWordAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)
}

func countWordAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func visit(node *html.Node, words, pics *int) {
	if node.Type == html.TextNode {
		*words += len(strings.Fields(node.Data))
	} else if node.Type == html.ElementNode && node.Data == "img" {
		*pics++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}
