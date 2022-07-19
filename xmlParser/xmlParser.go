package xmlparser

import "encoding/xml"

type urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []url    `xml:"url"`
}

type url struct {
	Loc string `xml:"loc"`
}

func Generate(urls map[string]string) (x []byte, err error) {
	x = []byte{}
	x = append(x, []byte(xml.Header)...)

	urlSet := &urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for _, u := range urls {
		urlSet.Urls = append(urlSet.Urls, url{Loc: u})
	}
	data, err := xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		return
	}
	x = append(x, data...)
	return
}
