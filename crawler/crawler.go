package crawler

import (
	"io/ioutil"
	"net/http"
	"sync"

	htmlparser "github.com/camopy/sitemap-generator/htmlParser"
)

type Crawler struct {
	baseUrl   string
	depth     int
	workers   int
	queueCh   chan anchors
	errCh     chan error
	doneCh    chan bool
	workerCh  chan bool
	urls      map[string]string
	urlMu     sync.Mutex
	crawled   map[string]bool
	crawledMu sync.Mutex
	wg        sync.WaitGroup
}

type anchors struct {
	urls  []string
	depth int
}

func New(baseUrl string, depth, workers int) *Crawler {
	return &Crawler{
		baseUrl:  baseUrl,
		depth:    depth,
		workers:  workers,
		queueCh:  make(chan anchors, 100),
		errCh:    make(chan error),
		doneCh:   make(chan bool),
		workerCh: make(chan bool, workers),
		urls:     make(map[string]string),
		crawled:  make(map[string]bool),
		wg:       sync.WaitGroup{},
	}
}

func (c *Crawler) Start() (map[string]string, error) {
	c.startWorkers()
	c.startBaseUrlCrawling()

	go c.waitWorkersAndCloseQueue()

	for {
		select {
		case anchors := <-c.queueCh:
			for _, url := range anchors.urls {
				if c.alreadyCrawled(url) {
					c.wg.Done()
					continue
				}
				<-c.workerCh
				go c.crawl(url, anchors.depth)
			}
		case err := <-c.errCh:
			return nil, err
		case <-c.doneCh:
			return c.urls, nil
		}
	}
}

func (c *Crawler) waitWorkersAndCloseQueue() {
	c.wg.Wait()
	c.doneCh <- true
}

func (c *Crawler) crawl(url string, depth int) {
	defer func() { c.workerCh <- true }()
	defer c.wg.Done()

	r, err := http.Get(url)
	if err != nil {
		c.errCh <- err
		return
	}
	defer r.Body.Close()

	html, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.errCh <- err
		return
	}
	a := htmlparser.ParseAnchors(html, c.baseUrl)
	if len(a) == 0 {
		return
	}

	c.addAnchors(a)
	c.markAsCrawled(url)

	if depth < c.depth {
		c.enqueue(anchors{urls: a, depth: depth + 1})
	}
}

func (c *Crawler) alreadyCrawled(url string) bool {
	c.crawledMu.Lock()
	_, ok := c.crawled[url]
	c.crawledMu.Unlock()
	return ok
}

func (c *Crawler) addAnchors(anchors []string) {
	c.urlMu.Lock()
	for _, a := range anchors {
		c.urls[a] = a
	}
	c.urlMu.Unlock()
}

func (c *Crawler) markAsCrawled(url string) {
	c.crawledMu.Lock()
	c.crawled[url] = true
	c.crawledMu.Unlock()
}

func (c *Crawler) enqueue(a anchors) {
	c.wg.Add(len(a.urls))
	c.queueCh <- a
}

func (c *Crawler) startWorkers() {
	for i := 0; i < c.workers; i++ {
		c.workerCh <- true
	}
}

func (c *Crawler) startBaseUrlCrawling() {
	c.addAnchors([]string{c.baseUrl})
	c.enqueue(anchors{urls: []string{c.baseUrl}, depth: 1})
}
