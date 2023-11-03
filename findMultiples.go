package main

import (
	"fmt"
)

func main() {
	var multiple1, multiple2, max int

	// Receba os valores de entrada
	fmt.Print("Digite o primeiro múltiplo: ")
	fmt.Scan(&multiple1)

	fmt.Print("Digite o segundo múltiplo: ")
	fmt.Scan(&multiple2)

	fmt.Print("Digite o número máximo da sequência: ")
	fmt.Scan(&max)

	fmt.Printf("Múltiplos de %d e %d até %d:\n", multiple1, multiple2, max)

	for i := 1; i <= max; i++ {
		if i%multiple1 == 0 && i%multiple2 == 0 {
			fmt.Println("ping pang")
		}
		if i%multiple1 == 0 {
			fmt.Println("ping")
		}
		if i%multiple2 == 0 {
			fmt.Println("pang")
		}
	}
}
