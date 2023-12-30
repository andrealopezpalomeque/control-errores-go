package main

import (
	"fmt"
	"strconv"
)

func divider(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//creo un error personalizado
		//return 0, errors.New("No se puede dividir por cero") // con el paquete ERRORS
		return 0, fmt.Errorf("No se puede dividir por cero") //tambien con el paquete FMT

	}

	return dividendo / divisor, nil
}

func main() {
	str := "123"
	strconv.Atoi(str)

	num, err := strconv.Atoi(str) //trata de convertir el string a un entero, si no puede devuelve un error

	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("El número es:", num)
	}

	fmt.Print("--------------------\n")
	//======================================================FUNCION DIVIDER===============================================
	result, err := divider(10, 0) //

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("El resultado es: ", result)
	}

}