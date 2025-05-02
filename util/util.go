package util

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

type MetaFinder struct {
	finderKey   string
	finderValue string
	valueKey    string
}

func GetPageMetaData(body string) (string, string, error) {
	bodyReader := strings.NewReader(string(body))
	doc, err := html.Parse(bodyReader)
	stripper := bluemonday.StripTagsPolicy()
	if err != nil {
		slog.Error("unable to parse page body: " + err.Error())
		return "", "", fmt.Errorf("unable to parse page body")
	}
	title, err := scanNode(doc, "title", &MetaFinder{finderKey: "", finderValue: "", valueKey: ""})
	if err != nil {
		return "", "", err
	}
	desc, err := scanNode(doc, "meta", &MetaFinder{finderKey: "name", finderValue: "description", valueKey: "content"})
	if err != nil {
		return "", "", err
	}
	return stripper.Sanitize(html.UnescapeString(title)), stripper.Sanitize(html.UnescapeString(desc)), nil
}

func scanNode(n *html.Node, tagName string, attr *MetaFinder) (string, error) {
	if n.Type == html.ElementNode && n.Data == tagName {
		if attr.finderKey == "" {
			return n.FirstChild.Data, nil
		}
		check := false
		for _, a := range n.Attr {
			if a.Key == attr.finderKey && a.Val == attr.finderValue {
				check = true
			}
			if a.Key == attr.valueKey && check {
				return a.Val, nil
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		text, err := scanNode(child, tagName, attr)
		if text != "" {
			return text, err
		}
	}
	return "", nil
}
