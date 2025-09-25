// main.go: Programa principal del paquete 'TALLER_PAQUETEs'.
// Proporciona un menú interactivo para usar las funcionalidades:
// 1. Conversor de monedas (currency.go)
// 2. Contador de vocales (vowels.go)
// Usa bufio.Scanner para input del usuario.
package main

// ... (resto del código sin cambios)

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== Taller Práctico: Paquete de Utilidades ===")
	fmt.Println("1. Conversor de Monedas (USD a EUR/LB/WON/BTC)")
	fmt.Println("2. Contador de Vocales")
	fmt.Println("3. Salir")
	fmt.Print("Elige una opción (1-3): ")

	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "1":
		// Conversor de monedas
		fmt.Print("Ingresa el monto en USD: ")
		scanner.Scan()
		usdStr := strings.TrimSpace(scanner.Text())
		usd, err := strconv.ParseFloat(usdStr, 64)
		if err != nil {
			fmt.Println("Error: Monto inválido. Debe ser un número.")
			return
		}

		fmt.Print("Moneda objetivo (EUR, LB, WON, BTC): ")
		scanner.Scan()
		target := strings.TrimSpace(scanner.Text())

		PrintConversion(usd, target)

	case "2":
		// Contador de vocales
		fmt.Print("Ingresa una frase: ")
		scanner.Scan()
		phrase := strings.TrimSpace(scanner.Text())

		if phrase == "" {
			fmt.Println("Error: Frase vacía.")
			return
		}

		PrintVowelCounts(phrase)

	case "3":
		fmt.Println("Saliendo...")
		return

	default:
		fmt.Println("Opción inválida.")
	}
}
