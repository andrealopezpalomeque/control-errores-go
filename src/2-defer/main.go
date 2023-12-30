package main

import (
	"fmt"
	"os"
)


func main() {
	//SE AGREGAN A UNA PILA -> el primero en entrar es el ultimo en salir
	defer fmt.Println(3) //pospone la ejecucion de la funcion hasta que termine la funcion actual, en este caso main
	defer fmt.Println(2)	
	defer fmt.Println(1)

	file, err := os.Create("hola.txt")

	if err != nil {
		fmt.Println(err)
		return
	} 

	defer file.Close() //se ejecuta al final de la funcion main

	_, err = file.Write([]byte("Hello World, writing in Go Language"))
	if err != nil {
		fmt.Println(err)
		return
	}

	
	





}