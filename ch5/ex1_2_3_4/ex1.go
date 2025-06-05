package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}
	var res []string

	fmt.Println("links:")
	for _, link := range visit(res, doc) {
		fmt.Println(link)
	}

	types := make(map[string]int)
	htmlTypes(doc, types)
	fmt.Println("types:")
	for t, n := range types {
		fmt.Println(t, n)
	}

	printTextNodes(doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}

	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func htmlTypes(n *html.Node, types map[string]int) {
	if n.Type == html.ElementNode {
		types[n.Data]++

	}
	if n.FirstChild != nil {
		htmlTypes(n.FirstChild, types)
	}
	if n.NextSibling != nil {
		htmlTypes(n.NextSibling, types)
	}
}

func printTextNodes(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("%v", n.Data)
	}
	if n.FirstChild != nil {
		printTextNodes(n.FirstChild)
	}
	if n.NextSibling != nil {
		printTextNodes(n.NextSibling)
	}
}
