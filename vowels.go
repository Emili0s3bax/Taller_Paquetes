package main

import (
	"fmt"
	"strings"
)

// CountVowels cuenta las ocurrencias de cada vocal (a, e, i, o, u) en la frase.
// Ignora mayúsculas/minúsculas y caracteres no vocales.
// Retorna un map con las vocales como keys y conteos como values.
// Complejidad: O(n), donde n es la longitud de la frase.
// Ejemplo: CountVowels("Hello") -> map['e':1 'o':1]
func CountVowels(phrase string) map[rune]int {
	phrase = strings.ToLower(phrase) // Convierte a minúscula para ignorar case
	counts := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}

	for _, char := range phrase {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			counts[char]++
		}
	}

	return counts
}

// PrintVowelCounts imprime el conteo de vocales de manera formateada.
// Incluye la frase original para referencia.
// Ejemplo: PrintVowelCounts("Hola") -> Muestra 'a':1, 'o':1
func PrintVowelCounts(phrase string) {
	counts := CountVowels(phrase)
	fmt.Printf("Frase: \"%s\"\n", phrase)
	fmt.Println("Conteo de vocales:")
	for vowel, count := range counts {
		fmt.Printf("  %c: %d\n", vowel, count)
	}
}
