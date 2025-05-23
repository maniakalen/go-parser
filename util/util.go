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
	if len(title) > 1 && title[len(title)-2:] == "</" {
		title += "i>"
	}
	if len(desc) > 1 && desc[len(desc)-2:] == "</" {
		desc += "i>"
	}
	return stripper.Sanitize(html.UnescapeString(title)), stripper.Sanitize(html.UnescapeString(desc)), nil
}

func scanNode(n *html.Node, tagName string, attr *MetaFinder) (string, error) {
	field := ""
	if n.Type == html.ElementNode && n.Data == tagName {
		if attr.finderKey == "" && n.FirstChild != nil {
			field = n.FirstChild.Data
		}
		check := false
		for _, a := range n.Attr {
			if a.Key == attr.finderKey && a.Val == attr.finderValue {
				check = true
			}
			if a.Key == attr.valueKey && check {
				field = a.Val
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		text, err := scanNode(child, tagName, attr)
		if text != "" {
			return text, err
		}
	}
	return field, nil
}
