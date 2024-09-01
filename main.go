package main

import (
	"fmt"
	"os"
)

func main(){
	cmdArgs := os.Args[1:]
	if len(cmdArgs) < 1{
		fmt.Println("no website provided")
		os.Exit(1)
	}else if len(cmdArgs) > 1{
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url := cmdArgs[0]
	fmt.Printf("starting crawl of %v\n", url)

	pages := make(map[string]int)
	crawlPage(url, url, pages)
}