package sanitize

import (
	"regexp"
	"strings"
)

// SanitizeSQLInput remove padrões de SQL injection de uma string
func SanitizeSQLInput(input string) string {
	if input == "" {
		return input
	}

	// Removendo SQL
	sqlPattern := regexp.MustCompile(`(?i)(\b(SELECT|INSERT|DELETE|UPDATE|DROP|UNION|--|;|OR|AND)\b)`)
	clean := sqlPattern.ReplaceAllString(input, "")

	// Collapsar espaços
	clean = strings.Join(strings.Fields(clean), " ")

	return strings.TrimSpace(clean)
}
