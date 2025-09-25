package main

import (
	"fmt"
	"strings"
)

// ConvertCurrency convierte un monto en USD a la moneda especificada.
// Soporta: EUR (Euros), LB/GBP (Libras Esterlinas), WON/KRW (Won Surcoreano), BTC (Bitcoin).
// Tasas aproximadas actualizadas a octubre 2024 (fijas para demo; usa API para valores reales).
// Retorna el monto convertido y un error si la moneda no es válida.
// Ejemplo: ConvertCurrency(100, "EUR") -> 92.0, nil
func ConvertCurrency(usd float64, targetCurrency string) (float64, error) {
	target := strings.ToUpper(strings.TrimSpace(targetCurrency))

	switch target {
	case "EUR", "EUROS":
		return usd * 0.92, nil // 1 USD ≈ 0.92 EUR
	case "LB", "LIBRAS", "GBP":
		return usd * 0.77, nil // 1 USD ≈ 0.77 GBP
	case "WON", "KRW":
		return usd * 1330, nil // 1 USD ≈ 1330 KRW
	case "BTC":
		return usd * 0.0000167, nil // 1 USD ≈ 0.0000167 BTC (volátil)
	default:
		return 0, fmt.Errorf("moneda no soportada: %s. Opciones: EUR, LB, WON, BTC", target)
	}
}

// PrintConversion imprime el resultado de la conversión en formato legible.
// Maneja errores y muestra el símbolo de la moneda.
// Ejemplo: PrintConversion(100, "BTC") -> "100.00 USD = 0.0017 BTC"
func PrintConversion(usd float64, targetCurrency string) {
	result, err := ConvertCurrency(usd, targetCurrency)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var currencySymbol string
	switch strings.ToUpper(targetCurrency) {
	case "EUR", "EUROS":
		currencySymbol = "EUR"
	case "LB", "LIBRAS", "GBP":
		currencySymbol = "GBP"
	case "WON", "KRW":
		currencySymbol = "KRW"
	case "BTC":
		currencySymbol = "BTC"
	default:
		currencySymbol = targetCurrency
	}

	fmt.Printf("%.2f USD = %.4f %s\n", usd, result, currencySymbol)
}
