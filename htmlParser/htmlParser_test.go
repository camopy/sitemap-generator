package htmlparser_test

import (
	"reflect"
	"testing"

	htmlparser "github.com/camopy/sitemap-generator/htmlParser"
)

func TestParseAnchors(t *testing.T) {
	base := "https://www.sitemaps.org/"
	baseWithoutSlash := "https://www.sitemaps.org"

	tests := []struct {
		name string
		base string
		html []byte
		want []string
	}{
		{
			name: "one anchor",
			base: base,
			html: []byte(`<a href="https://www.sitemaps.org/protocol.html">Protocol</a>`),
			want: []string{
				"https://www.sitemaps.org/protocol.html",
			},
		},
		{
			name: "two anchors",
			base: base,
			html: []byte(`<a href="https://www.sitemaps.org/protocol.html">Protocol</a>
			<a href="https://www.sitemaps.org/faq.html">Protocol</a>`),
			want: []string{
				"https://www.sitemaps.org/protocol.html",
				"https://www.sitemaps.org/faq.html",
			},
		},
		{
			name: "two short anchors",
			base: base,
			html: []byte(`<a href="protocol.html">Protocol</a>
			<a href="faq.html">Protocol</a>`),
			want: []string{
				base + "protocol.html",
				base + "faq.html",
			},
		},
		{
			name: "two short anchors and base without slash",
			base: baseWithoutSlash,
			html: []byte(`<a href="protocol.html">Protocol</a>
			<a href="faq.html">Protocol</a>`),
			want: []string{
				baseWithoutSlash + "/protocol.html",
				baseWithoutSlash + "/faq.html",
			},
		},
		{
			name: "zero anchors",
			base: base,
			html: []byte(`<i href="https://www.sitemaps.org/protocol.html">Protocol</i>
			<i href="https://www.sitemaps.org/faq.html">Protocol</i>`),
			want: []string{},
		},
		{
			name: "# id",
			base: base,
			html: []byte(`<a href="#test">Protocol</a>`),
			want: []string{},
		},
		{
			name: "mailto",
			base: base,
			html: []byte(`<a href="mailto:test">Protocol</a>`),
			want: []string{},
		},
		{
			name: "external link",
			base: base,
			html: []byte(`<a href="https://www.sitemaps.org/protocol.html">Protocol</a>
			<a href="https://www.external.org/faq.html">Protocol</a>`),
			want: []string{
				"https://www.sitemaps.org/protocol.html",
			},
		},
		{
			name: "xml",
			base: base,
			html: []byte(`asfd   <a href="https://www.sitemaps.org/protocol.html">Professional Programming</a></h2>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1657979857493/xYDyZtu-x.png" alt="Pr`),
			want: []string{
				"https://www.sitemaps.org/protocol.html",
			},
		},
		{
			name: "xml",
			base: base,
			html: []byte(`asfd   <a href="https://www.sitemaps.org/">Professional Programming</a></h2>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1657979857493/xYDyZtu-x.png" alt="Pr`),
			want: []string{
				"https://www.sitemaps.org/",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := htmlparser.ParseAnchors(tt.html, tt.base)
			if len(tt.want) == 0 && len(got) != 0 {
				t.Errorf("got %v, want %v", got, tt.want)
			} else if len(tt.want) > 0 && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
