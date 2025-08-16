package formatters

import (
	"fmt"
	"regexp"
)

// FormatPhoneNumber recebe uma string de 11 digitos e retorna no formato "(XX) X XXXX-XXXX"
func FormatPhoneNumber(phone string) (string, error) {
	re := regexp.MustCompile(`\D`)
	phone = re.ReplaceAllString(phone, "")

	if len(phone) != 11 {
		return "", fmt.Errorf("telefone deve conter 11 digitos")
	}

	// Apply formatting
	formatted := fmt.Sprintf("(%s) %s %s-%s",
		phone[0:2], // DDD
		phone[2:3], // 9
		phone[3:7], // 1234
		phone[7:],  // 5678
	)

	return formatted, nil
}
