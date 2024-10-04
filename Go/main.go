package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Iniciar el cronómetro para medir el tiempo de ejecución.
	
	// Definir las letras y la palabra a adivinar.
	letters := "aábcdeéfghiíjklmnñoópqrstuúüvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.,;:!?¿¡-"
	
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)
	
	start := time.Now()
	
	guess := ""
	attempts := 0
	
	// Sembrar el generador de números aleatorios.
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	for {
		// Mantener solo la última parte del guess si excede la longitud de la palabra.
		if attempts >= len(word) {
			guess = guess[len(guess)-len(word):]
		}

		// Romper el bucle si la palabra se adivinó.
		if attempts >= len(word) && guess == word {
			break
		}

		// Generar un índice aleatorio y agregar el carácter correspondiente a guess.
		index := randomGenerator.Intn(len(letters))
		guess += string(letters[index])

		attempts++
	}

	// Medir el tiempo transcurrido.
	elapsed := time.Since(start)

	fmt.Printf("Attempts: %d\n", attempts)
	fmt.Printf("Time: %.2f seconds\n", elapsed.Seconds())
}