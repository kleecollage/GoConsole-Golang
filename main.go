package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	// fmt.Println(numAleatorio)
	var numIngresado int
	var intentos int
	const maxIntentos = 5

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un numero del 0 al 100 (intentos restantes %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("\n¡Felicidades, Adivinsate el numero!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("Estoy pensando en un numero mayor\n")
		} else if numIngresado > numAleatorio {
			fmt.Println("Estoy pensando en un numero menor\n")
		}
	}
	fmt.Printf("\nSe acabaron los intentos :( \nEl numero que pense es: %d\n", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Println("\n¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("\nGAME OVER")
	default:
		fmt.Println("Creo que no has entendido bien las opciones")
		jugarNuevamente()
	}
}
