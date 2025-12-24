package parser

import (
	"regexp"
	"strings"
)

type CodeResult struct {
	Html string
	Css  string
	Js   string
}

// ParseMultiFile extracts html, css, js code from markdown content
func ParseMultiFile(content string) *CodeResult {
	res := &CodeResult{}

	htmlRe := regexp.MustCompile("(?i)```html\\s*\\n([\\s\\S]*?)```")
	cssRe := regexp.MustCompile("(?i)```css\\s*\\n([\\s\\S]*?)```")
	jsRe := regexp.MustCompile("(?i)```(?:js|javascript)\\s*\\n([\\s\\S]*?)```")

	if match := htmlRe.FindStringSubmatch(content); len(match) >= 2 {
		res.Html = strings.TrimSpace(match[1])
	}
	if match := cssRe.FindStringSubmatch(content); len(match) >= 2 {
		res.Css = strings.TrimSpace(match[1])
	}
	if match := jsRe.FindStringSubmatch(content); len(match) >= 2 {
		res.Js = strings.TrimSpace(match[1])
	}

	// Fallback/Legacy: If specific blocks not found, try generic block for HTML
	if res.Html == "" {
		re := regexp.MustCompile("(?s)```\\w*\\n(.*?)\\n```")
		if matches := re.FindStringSubmatch(content); len(matches) >= 2 {
			res.Html = strings.TrimSpace(matches[1])
		} else {
			// Even weaker fallback: if content looks like HTML but no ticks
			if strings.Contains(content, "<html") || strings.Contains(content, "<div") {
				res.Html = content
			}
		}
	}

	return res
}
