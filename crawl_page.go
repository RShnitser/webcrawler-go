package main

import(
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.RWMutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool){
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	count, ok := cfg.pages[normalizedURL]
	if ok{
		cfg.pages[normalizedURL] = count + 1
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) checkPageCount() (shouldEnd bool){
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()
	count := len(cfg.pages)
	if count > cfg.maxPages{
		return true
	}
	return false
}

func (cfg *config)crawlPage(rawCurrentURL string){
	cfg.concurrencyControl<- struct{}{}
	
	defer func(){
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	if cfg.checkPageCount(){
		return
	}
	
	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrent.Hostname(){
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil{
		return
	}


	first := cfg.addPageVisit(normalized)
	if !first{
		return
	}

	//fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil{
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil{
		return
	}

	for _, url := range urls{
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
