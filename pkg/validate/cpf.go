package validate

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
)

// Valida o CPF calculando os dois ultimos digitos e checando se estão de acordo
func IsValidCPF(cpf string, ctx context.Context) bool {
	if ctx == nil {
		return true
	}
	tracer := otel.Tracer("contacts-service")
	_, span := tracer.Start(ctx, "validateCPF")
	defer span.End()

	start := time.Now()
	log.Printf("[TRACE] Validando CPF: %s", cpf)

	// Usar apenas digitos do CPF
	re := regexp.MustCompile(`[^0-9]`)
	cpf = re.ReplaceAllString(cpf, "")

	// Deve ser de 11 digitos
	if len(cpf) != 11 {
		span.RecordError(fmt.Errorf("CPF inválido (tamanho)"))
		log.Printf("[TRACE] CPF inválido: %s | Tempo: %s", cpf, time.Since(start))
		return false
	}

	// Calculo do primeiro digito
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		current := digit * (10 - i)
		sum += current

	}
	firstDigitMod := sum % 11

	var firstCheck int
	if firstDigitMod >= 2 {
		firstCheck = 11 - firstDigitMod
	}

	if firstCheck != int(cpf[9]-'0') {
		return false
	}

	// Calculo do segundo digito
	sum = 0

	for i := 0; i < 10; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		current := digit * (11 - i)
		sum += current
	}

	secondCheckMod := sum % 11
	var secondCheck int

	if secondCheckMod >= 2 {
		secondCheck = 11 - secondCheckMod
	}
	if secondCheck != int(cpf[10]-'0') {
		return false
	}
	log.Printf("[TRACE] CPF válido: %s | Tempo: %s", cpf, time.Since(start))
	return true
}
