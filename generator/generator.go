package generator

import (
	"github.com/camopy/sitemap-generator/crawler"
	xmlparser "github.com/camopy/sitemap-generator/xmlParser"
)

type generator struct {
	crawler *crawler.Crawler
}

func New(baseUrl string, maxDepth, parallel int) *generator {
	return &generator{
		crawler: crawler.New(baseUrl, maxDepth, parallel),
	}
}

func (g *generator) Generate() ([]byte, error) {
	urls, err := g.findUrls()
	if err != nil {
		return nil, err
	}
	return g.generateXml(urls)
}

func (g *generator) findUrls() (map[string]string, error) {
	return g.crawler.Start()
}

func (g *generator) generateXml(urls map[string]string) ([]byte, error) {
	return xmlparser.Generate(urls)
}
