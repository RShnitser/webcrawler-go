package main

import (
	"fmt"
	"os"
	"sync"
	"net/url"
	"strconv"
)

func main(){
	cmdArgs := os.Args[1:]
	if len(cmdArgs) < 3{
		fmt.Println("usage URL maxConcurrency maxPages")
		os.Exit(1)
	}else if len(cmdArgs) > 3{
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawURL := cmdArgs[0]

	parsedBase, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("unable to parse url")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(cmdArgs[1])
	if err != nil {
		fmt.Println("maxConcurrency must be an int")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(cmdArgs[2])
	if err != nil {
		fmt.Println("maxPages must be an int")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of %v\n", rawURL)

	cfg := config{
		make(map[string]int),
		parsedBase,
		&sync.RWMutex{},
		make(chan struct{}, maxConcurrency),
		&sync.WaitGroup{},
		maxPages,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	// for normalizedURL, count := range cfg.pages {
	// 	fmt.Printf("%d - %s\n", count, normalizedURL)
	// }

	cfg.printReport()
}