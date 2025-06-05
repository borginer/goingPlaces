package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	if node := ElementByID(doc, "toc"); node != nil {
		fmt.Printf("%v, %v\n", node.Type, node.Data)
	}
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}

	if node, ok := forEachNode(doc, pre, nil); ok {
		return node
	}
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (*html.Node, bool) {
	if pre != nil {
		if pre(n) {
			return n, true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, ok := forEachNode(c, pre, post); ok {
			return node, true
		}
	}

	if post != nil {
		if post(n) {
			return n, true
		}
	}
	return nil, false
}
