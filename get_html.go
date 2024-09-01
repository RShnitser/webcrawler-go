package main

import(
	"fmt"
	"io"
	"strings"
	"net/http"
)

func getHTML(rawURL string) (string, error){
	res, err := http.Get(rawURL)
	if err != nil{
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399{
		return "", fmt.Errorf("error: %v", res.StatusCode)
	}

	header := res.Header.Get("content-type")
	if !strings.Contains(header, "text/html"){
		return "", fmt.Errorf("error: invalid content-type %v", header)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	return string(data), nil
}