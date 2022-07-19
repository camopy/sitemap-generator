package xmlparser_test

import (
	"strings"
	"testing"

	xmlparser "github.com/camopy/sitemap-generator/xmlParser"
)

func TestGenerate(t *testing.T) {
	inputUrls := map[string]string{
		"http://example.com/":      "http://example.com/",
		"http://example.com/page1": "http://example.com/page1",
		"http://example.com/page2": "http://example.com/page2",
		"http://example.com/page3": "http://example.com/page3",
	}
	expectedXml := []string{
		`<url><loc>http://example.com/</loc></url>`,
		`<url><loc>http://example.com/page1</loc></url>`,
		`<url><loc>http://example.com/page2</loc></url>`,
		`<url><loc>http://example.com/page3</loc></url>`,
	}

	inputUrls2 := map[string]string{
		"http://example.com/":      "http://example.com/",
		"http://example.com/page1": "http://example.com/page1",
	}

	expectedXml2 := []string{
		`<url><loc>http://example.com/</loc></url>`,
		`<url><loc>http://example.com/page1</loc></url>`,
	}

	tests := []struct {
		name  string
		input map[string]string
		want  []string
	}{
		{
			name:  "generate xml",
			input: inputUrls,
			want:  expectedXml,
		},
		{
			name:  "generate xml",
			input: inputUrls2,
			want:  expectedXml2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := xmlparser.Generate(tt.input)
			if err != nil {
				t.Errorf("generateXml() error = %v", err)
				return
			}
			for _, u := range tt.input {
				if !strings.Contains(string(got), u) {
					t.Errorf("\ngot: \n\n%v, \n\nwant: \n\n%v\n\n", string(got), tt.want)
				}
			}
		})
	}
}
