package main

import(
	"fmt"
	"sort"
)

type page struct{
	url string
	count int
}

func (cfg *config)printReport(){
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", cfg.baseURL.String())
	fmt.Println("=============================")

	pageCounts := make([]page, len(cfg.pages))
	for k, v := range cfg.pages{
		pageCounts = append(pageCounts, page{k, v})
	}

	sort.Slice(pageCounts, func(a, b int) bool{
		return pageCounts[a].count < pageCounts[b].count
	})

	for _, page := range pageCounts{
		fmt.Printf("Found %v internal links to %s\n", page.count, page.url)
	}
}