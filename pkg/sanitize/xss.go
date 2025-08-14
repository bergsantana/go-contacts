package sanitize

import (
	"github.com/microcosm-cc/bluemonday"
)

// Removes all HTML tags, leaving only plain text
func StrictHTML(input string) string {
	p := bluemonday.StrictPolicy()

	return p.Sanitize(input)
}
