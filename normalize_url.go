package main

import( 
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {

	parsed, err := url.Parse(rawURL)
	if err != nil{
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	result := parsed.Host + parsed.Path
	result = strings.ToLower(result)
	result = strings.TrimRight(result, "/")
	return result, nil
}