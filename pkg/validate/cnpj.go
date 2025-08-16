package validate

import (
	"regexp"
	"strconv"
)

func IsValidCNPJ(cnpj string) bool {
	re := regexp.MustCompile(`\D`)
	cnpj = re.ReplaceAllString(cnpj, "")

	// Deverá haver exatos 14 digitos
	if len(cnpj) != 14 {
		return false
	}

	// Primeiros 12 digitos
	nums := make([]int, 14)
	for i := 0; i < 14; i++ {
		n, err := strconv.Atoi(string(cnpj[i]))
		if err != nil {
			return false
		}
		nums[i] = n
	}

	weights := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Pesos usados na primeira checagem
	weightsFirstCheck := weights[1:]
	sum := 0

	// Verificação do primeiro digito
	for i := 0; i < 12; i++ {
		sum += nums[i] * weightsFirstCheck[i]
	}
	remainder := sum % 11
	d1 := 0
	if remainder >= 2 {
		d1 = 11 - remainder
	}

	// Verificação do segundo digito
	sum = 0
	for i := 0; i < 13; i++ {
		sum += nums[i] * weights[i]
	}
	remainder = sum % 11
	d2 := 0
	if remainder >= 2 {
		d2 = 11 - remainder
	}

	return nums[12] == d1 && nums[13] == d2
}
