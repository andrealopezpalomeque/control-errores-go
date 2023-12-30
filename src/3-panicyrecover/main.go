package main

import "fmt"


func divider(a, b int)  {

	//creo funcion anonima
	defer func() {
		//recupero el control del programa
		//agarro el valor del panico pero sigue ejecutando el programa
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	validateZero(b)
	fmt.Println(a/b)

}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero") //PANIC
	}
}

func main() {

	//panic -> interrupcion del programa cuando se produce un error
	//recover -> recupera el control del programa despues de un panic

	divider(10, 10)
	divider(10, 0)

}