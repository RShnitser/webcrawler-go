package main

import(
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int){
	parsedBase, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if parsedBase.Hostname() != parsedCurrent.Hostname(){
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil{
		return
	}

	count, ok := pages[normalized]
	if ok{
		pages[normalized] = count + 1
		return
	}else{
		pages[normalized] = 1
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil{
		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil{
		return
	}

	for _, url := range urls{
		crawlPage(rawBaseURL, url, pages)
		fmt.Println(url)
	}
}
