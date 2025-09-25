# Taller_Paquetes

# Taller Práctico: Paquete de Utilidades en Go

Este es un paquete modular en Golang para el taller de creación de paquetes. Contiene dos funcionalidades principales:
- **Conversor de Monedas**: Convierte montos en USD a EUR, LB (Libras Esterlinas), WON (Surcoreano) o BTC. Usa tasas fijas aproximadas (octubre 2024).
- **Contador de Vocales**: Cuenta las ocurrencias de vocales (a, e, i, o, u) en una frase, ignorando mayúsculas.

## Estructura del Paquete
- `go.mod`: Inicializa el módulo `taller`.
- `main.go`: Menú interactivo para probar las funciones.
- `currency.go`: Funciones para conversión de monedas (`ConvertCurrency`, `PrintConversion`).
- `vowels.go`: Funciones para contar vocales (`CountVowels`, `PrintVowelCounts`).

## Instalación y Uso Local
1. Clona el repositorio: `git clone https://github.com/tu-usuario/taller.git`
2. Entra al directorio: `cd taller`
3. Inicializa dependencias: `go mod tidy`
4. Ejecuta: `go run .`

## Uso como Dependencia en Otro Proyecto
En un nuevo proyecto Go, agrega como dependencia: