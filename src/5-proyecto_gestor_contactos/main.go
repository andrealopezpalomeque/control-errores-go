package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)


type Contact struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//guardar contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	//crear el archivo
	file, err := os.Create("contacts.json")
	if err != nil{
		return err
	}
	defer file.Close()

	//crear un ENCODER para el archivo
	encoder := json.NewEncoder(file)

	//recuperar un error
	err = encoder.Encode(contacts)

	if err != nil{
		return err
	}

	return nil
}

//recuperar contactos de un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil{
		return err
	}
	defer file.Close()

	//crear un DECODER para el archivo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)

	if err != nil{
		return err
	}

	return nil
}




func main(){

	//slice de de contactos

	var contacts []Contact

	//cargar contacto desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil{
		fmt.Println("error al cargar los contactos: ", err)
		
	}

	//para ingresar un nuevo contacto creo una instancia de BUFIO
	//para leer desde la consola
	reader := bufio.NewReader(os.Stdin)

	for {
		//mostrar opciones al usuario
		fmt.Print("----gestor de contactos----\n",
			"1. Agregar contacto\n",
			"2. Mostrar contactos\n",
			"3. Salir\n",
			"Elija una opcion: ")



			//leer la opcion del usuario
	var option int
	_, err := fmt.Scanln(&option)
	if err != nil{
		fmt.Println("error al leer la opcion: ", err)
		return
	}

	//manejar las opciones del usuario
	switch option {
	case 1:
		//crear
		var c Contact
		fmt.Print("Nombre: ")
		c.Name, _ = reader.ReadString('\n')
		fmt.Print("Email: ")
		c.Email, _ = reader.ReadString('\n')
		fmt.Print("Telefono: ")
		c.Phone, _ = reader.ReadString('\n')
	
		//agregar el contacto al slice
		contacts = append(contacts, c)

		//guardar el contacto en el archivo
		if err := saveContactsToFile(contacts); err != nil{
			fmt.Println("error al guardar el contacto: ", err)
			return
		}

	case 2:
		//mostrar
		fmt.Println("------------")
		for index, contact := range contacts{
			fmt.Printf("%d. %s %s %s\n", index+1, contact.Name, contact.Email, contact.Phone)
		}
		fmt.Println("------------")

	case 3:
		//salir del programa
		return
	
	default:
		fmt.Println("opcion no valida")
	}
}

}