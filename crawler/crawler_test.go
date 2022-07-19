package crawler_test

import (
	"reflect"
	"testing"

	"github.com/camopy/sitemap-generator/crawler"
)

func TestStartCrawler(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("Skipping crawler test in short mode")
	// }

	c := crawler.New("https://www.sitemaps.org/", 1, 1)
	c2 := crawler.New("https://www.sitemaps.org/", 2, 1)
	c4 := crawler.New("https://www.sitemaps.org/", 4, 1)
	tests := []struct {
		name string
		g    *crawler.Crawler
		base string
		want map[string]string
	}{
		{
			name: "test1",
			g:    c,
			want: map[string]string{
				"https://www.sitemaps.org/":             "https://www.sitemaps.org/",
				"https://www.sitemaps.org/protocol.php": "https://www.sitemaps.org/protocol.php",
				"https://www.sitemaps.org/faq.php":      "https://www.sitemaps.org/faq.php",
				"https://www.sitemaps.org/terms.php":    "https://www.sitemaps.org/terms.php",
			},
		},
		{
			name: "depth 2",
			g:    c2,
			want: map[string]string{
				"https://www.sitemaps.org/":                        "https://www.sitemaps.org/",
				"https://www.sitemaps.org/protocol.php":            "https://www.sitemaps.org/protocol.php",
				"https://www.sitemaps.org/protocol.php#index":      "https://www.sitemaps.org/protocol.php#index",
				"https://www.sitemaps.org/protocol.php#lastmoddef": "https://www.sitemaps.org/protocol.php#lastmoddef",
				"https://www.sitemaps.org/protocol.php#validating": "https://www.sitemaps.org/protocol.php#validating",
				"https://www.sitemaps.org/faq.php":                 "https://www.sitemaps.org/faq.php",
				"https://www.sitemaps.org/terms.php":               "https://www.sitemaps.org/terms.php",
				"https://www.sitemaps.org/index.php":               "https://www.sitemaps.org/index.php",
			},
		},
		{
			name: "depth 4",
			g:    c4,
			want: map[string]string{
				"https://www.sitemaps.org/":                        "https://www.sitemaps.org/",
				"https://www.sitemaps.org/protocol.php":            "https://www.sitemaps.org/protocol.php",
				"https://www.sitemaps.org/protocol.php#index":      "https://www.sitemaps.org/protocol.php#index",
				"https://www.sitemaps.org/protocol.php#lastmoddef": "https://www.sitemaps.org/protocol.php#lastmoddef",
				"https://www.sitemaps.org/protocol.php#validating": "https://www.sitemaps.org/protocol.php#validating",
				"https://www.sitemaps.org/faq.php":                 "https://www.sitemaps.org/faq.php",
				"https://www.sitemaps.org/terms.php":               "https://www.sitemaps.org/terms.php",
				"https://www.sitemaps.org/index.php":               "https://www.sitemaps.org/index.php",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.Start()
			if err != nil {
				t.Errorf("Crawler.Start() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot: \n%v\n, \nwant: \n%v\n", got, tt.want)
			}
		})
	}
}
