package htmlparser

import (
	"strings"
)

const ANCHOR = `<a`
const HREF = ` href="`
const QUOTATION = `"`
const HASHTAG = `#`
const SLASH = `/`
const MAILTO = `mailto:`

func ParseAnchors(html []byte, base string) (anchors []string) {
	for i := 0; i < len(html)-1; i++ {
		if getSubstring(html, i, i+2) == ANCHOR {
		AnchorLookup:
			for j := i + 2; j < len(html)-len(HREF); j++ {
				index := j + len(HREF)
				t := getSubstring(html, j, index)
				if t == HREF {
					if getSubstring(html, index, index+1) == HASHTAG {
						break AnchorLookup
					}
					for k := index; k < len(html); k++ {
						if getSubstring(html, k, k+1) == QUOTATION {
							href := getSubstring(html, index, k)
							if isValidHref(href, base) {
								if !hasBasePrefix(href, base) {
									href = appendBaseToHref(href, base)
								}
								anchors = append(anchors, href)
							}
							i = k + 1
							break AnchorLookup
						}
					}
				}
			}
		}
	}
	return
}

func isValidHref(href, base string) bool {
	return !isMailTo(href) && !isExternalHref(href, base)
}

func isMailTo(href string) bool {
	return strings.HasPrefix(href, MAILTO)
}

func isExternalHref(href, base string) bool {
	if !hasHttpPrefix(href) {
		return false
	}
	return !hasBasePrefix(href, base)
}

func hasHttpPrefix(href string) bool {
	return strings.HasPrefix(href, "http")
}

func hasBasePrefix(href, base string) bool {
	return strings.HasPrefix(href, base)
}

func getSubstring(html []byte, start, end int) string {
	return string(html[start:end])
}

func appendBaseToHref(href, base string) string {
	if !strings.HasPrefix(href, SLASH) && !strings.HasSuffix(base, SLASH) {
		return base + SLASH + href
	} else if strings.HasPrefix(href, SLASH) && strings.HasSuffix(base, SLASH) {
		return base + href[1:]
	} else {
		return base + href
	}
}
