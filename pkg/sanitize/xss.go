package sanitize

import (
	"github.com/microcosm-cc/bluemonday"
)

// Remover HTML por texto simples
func StrictHTML(input string) string {
	p := bluemonday.StrictPolicy()

	return p.Sanitize(input)
}
