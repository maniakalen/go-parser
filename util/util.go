package util

import (
	"fmt"
	"log/slog"
	"strings"

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
	return title, desc, nil
}

func scanNode(n *html.Node, tagName string, attr *MetaFinder) (string, error) {
	if n.Type == html.ElementNode && n.Data == tagName {
		if attr.finderKey == "" {
			return n.FirstChild.Data, nil
		}
		check := false
		for _, a := range n.Attr {
			if a.Key == attr.finderKey && a.Val == attr.finderValue {
				fmt.Printf("%+v %+v\n", a, attr)
				check = true
			}
			fmt.Printf("%+v %+v %+v\n", a.Key, attr.valueKey, check)
			if a.Key == attr.valueKey && check {
				fmt.Printf("A %+v\n", a)
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
