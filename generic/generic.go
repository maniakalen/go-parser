package generic

import (
	"fmt"
	"log/slog"
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

// ParseHtml parses the html body and checks if it contains any of the keywords
func ParseHtml(body string, keywords []string) (bool, error) {
	bodyReader := strings.NewReader(string(body))
	doc, err := html.Parse(bodyReader)
	if err != nil {
		slog.Error("unable to parse page body: " + err.Error())
		return false, fmt.Errorf("unable to parse page body")
	}
	return scanNode(doc, keywords)
}

// returns true if any keyword is found in the body
func scanNode(n *html.Node, keywords []string) (bool, error) {
	if n.Type == html.ElementNode && (n.Data == "title") {
		fields := strings.Fields(n.FirstChild.Data)
		for _, keyword := range keywords {
			if slices.Contains(fields, keyword) {
				return true, nil
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "meta" {
		check := false
		for _, a := range n.Attr {
			if a.Key == "name" && (a.Val == "description" || a.Val == "keywords") {
				check = true
			}
			if a.Key == "content" && check {
				fields := strings.Fields(a.Val)
				for _, keyword := range keywords {
					if slices.Contains(fields, keyword) {
						return true, nil
					}
				}
			}
		}
	}
	if n.Type == html.ElementNode && n.Data == "h1" {
		if n.FirstChild != nil && len(n.FirstChild.Data) > 0 {
			fields := strings.Fields(n.FirstChild.Data)
			for _, keyword := range keywords {
				if slices.Contains(fields, keyword) {
					return true, nil
				}
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if found, err := scanNode(child, keywords); found {
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}
