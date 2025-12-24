package parser

import (
	"testing"
)

func TestParseMultiFile(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    *CodeResult
	}{
		{
			name:    "Single HTML block",
			content: "Here is the code:\n```html\n<div>Hello</div>\n```",
			want: &CodeResult{
				Html: "<div>Hello</div>",
			},
		},
		{
			name: "Multi-file: HTML+CSS+JS",
			// Using concatenation to avoid backtick issues in tool calling
			content: "Sure!\n" +
				"```html\n<div id=\"app\"></div>\n```\n" +
				"Style:\n" +
				"```css\n#app { color: red; }\n```\n" +
				"Script:\n" +
				"```javascript\nconsole.log('hello');\n```\n",
			want: &CodeResult{
				Html: "<div id=\"app\"></div>",
				Css:  "#app { color: red; }",
				Js:   "console.log('hello');",
			},
		},
		{
			name:    "Fallback: Generic block as HTML",
			content: "```\nJust some text\n```",
			want: &CodeResult{
				Html: "Just some text",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseMultiFile(tt.content)
			if got.Html != tt.want.Html {
				t.Errorf("Html = %v, want %v", got.Html, tt.want.Html)
			}
			if got.Css != tt.want.Css {
				t.Errorf("Css = %v, want %v", got.Css, tt.want.Css)
			}
			if got.Js != tt.want.Js {
				t.Errorf("Js = %v, want %v", got.Js, tt.want.Js)
			}
		})
	}
}
