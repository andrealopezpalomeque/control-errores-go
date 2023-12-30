package main

import (
	"log"
	"os"
)



func main(){
	log.SetPrefix("main: ") 
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) //se crea un archivo de log

	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close() //se cierra el archivo de log


	log.SetOutput(file) //se le dice a log que escriba en el archivo de log creado
	log.Print("soy un log")

	// log.SetPrefix("main: ") //se agrega un prefijo a los mensajes -> todos van a empezar con "main:"


	// log.Print("Hola mundo 1") //se imprime el tiempo y el mensaje -> 2020/07/21 11:22:33 Hola Mundo
	// log.Println("Hola mundo 2")

	

	// log.Fatal("Registro de error") //se imprime el tiempo y el mensaje y se detiene la ejecución -> 2020/07/21 11:22:33 Hola Mundo
	// log.Print("puedes verme?") // esto NO se ejecuta



}