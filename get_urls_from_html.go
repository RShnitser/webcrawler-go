package main

import(
	"strings"
	"fmt"
	"golang.org/x/net/html"
	"net/url"
)

func getURLsFromHTML(htmlBody, baseURL *url.URL) ([]string, error){
	
	htmlReader := strings.NewReader(htmlBody)
	htmlTree, err := html.Parse(htmlReader)
	if err != nil{
		return nil, fmt.Errorf("couldn't parse html body: %w", err)
	}
	
	
	var result []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					parsedLink, err := url.Parse(a.Val)
					if err != nil{
						fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
						continue
					}
					resolvedURL := parsedBase.ResolveReference(parsedLink)
					result = append(result, resolvedURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseNodes(c)
		}
	}
	traverseNodes(htmlTree)

	return result, nil
}