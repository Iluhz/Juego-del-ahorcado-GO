package main

import (
	"fmt"
	"math/rand"
)

func main() {

	palabras := []string{"gato", "perro", "casa", "programacion", "libro", "pako", "estructura", "idea", "golang", "tarea"}
	elegirAzar := rand.Intn(10)
	palabra := palabras[elegirAzar]

	var letrasUsadas [20]string
	totalLetrasUsadas := 0

	errores := 0

	fmt.Println("=== ahorcado ===")
	fmt.Println("La palabra q debes adivinar tiene", len(palabra), "letras")

	for {

		fmt.Println("")
		if errores == 0 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("      |")
			fmt.Println("      |")
			fmt.Println("      |")
			fmt.Println("=========")
		}
		if errores == 1 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("  O   |")
			fmt.Println("      |")
			fmt.Println("      |")
			fmt.Println("=========")
		}
		if errores == 2 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("  O   |")
			fmt.Println("  |   |")
			fmt.Println("      |")
			fmt.Println("=========")
		}
		if errores == 3 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("  O   |")
			fmt.Println(" /|   |")
			fmt.Println("      |")
			fmt.Println("=========")
		}
		if errores == 4 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("  O   |")
			fmt.Println(" /|L |")
			fmt.Println("      |")
			fmt.Println("=========")
		}
		if errores == 5 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("  O   |")
			fmt.Println(" /|L  |")
			fmt.Println(" /    |")
			fmt.Println("=========")
		}
		if errores == 6 {
			fmt.Println("  +---+")
			fmt.Println("  |   |")
			fmt.Println("  O   |")
			fmt.Println(" /|L  |")
			fmt.Println(" / L  |")
			fmt.Println("=========")
		}

		fmt.Print("La palabra es: ")
		letrasAdivinadas := 0
		for i := 0; i < len(palabra); i++ {
			letraDeLaPalabra := string(palabra[i])

			encontrada := false
			for j := 0; j < totalLetrasUsadas; j++ {
				if letraDeLaPalabra == letrasUsadas[j] {
					encontrada = true
				}
			}

			if encontrada {
				fmt.Print(letraDeLaPalabra + " ")
				letrasAdivinadas++
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println("")

		if totalLetrasUsadas > 0 {
			fmt.Print("Letras usadas: ")
			for i := 0; i < totalLetrasUsadas; i++ {
				fmt.Print(letrasUsadas[i] + " ")
			}
			fmt.Println("")
		}

		if letrasAdivinadas == len(palabra) {
			fmt.Println("Adivinaste la palabra, la palabra ear:", palabra)
			break
		}

		if errores == 6 {
			fmt.Println("no adivjiaste la palabra, la palabra era", palabra)
			break
		}

		fmt.Print("Escribe una letra: ")
		var entrada string
		fmt.Scan(&entrada)

		letra := string(entrada[0])

		fmt.Println("Ingresaste la letra:", letra)

		letrasUsadas[totalLetrasUsadas] = letra
		totalLetrasUsadas++

		letraCorrecta := false
		for i := 0; i < len(palabra); i++ {
			if letra == string(palabra[i]) {
				letraCorrecta = true
			}
		}

		if letraCorrecta {
			fmt.Println("La letra", letra, "esta en la palabra")
		} else {
			fmt.Println("La letra", letra, "no esta en la palabra")
			errores++
		}

	}
}
