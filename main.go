package main

import (
	"fmt"
	"os"
	"sync"
	"net/url"
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

	rawURL := cmdArgs[0]
	fmt.Printf("starting crawl of %v\n", rawURL)

	parsedBase, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("unable to parse url")
		os.Exit(1)
	}

	cfg := config{
		make(map[string]int),
		parsedBase,
		&sync.Mutex{},
		make(chan struct{}, 20),
		&sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()
}