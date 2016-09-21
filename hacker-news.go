package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP GET failed: %v\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "HTML parse failed: %v\n", err)
		os.Exit(1)
	}

	resp.Body.Close()

	visitNode(doc.FirstChild)
}

func visitNode(n *html.Node) {
	if (n.Type == html.ElementNode) && (n.Data == "a") {
		for _, attr := range n.Attr {
			if (attr.Key == "class") && (attr.Val == "storylink") {
				fmt.Printf("%s\n", n.FirstChild.Data)

				break
			}
		}
	}

	if n.FirstChild != nil {
		visitNode(n.FirstChild)
	}

	if n.NextSibling != nil {
		visitNode(n.NextSibling)
	}
}
