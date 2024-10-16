package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json: "name"`
	Email string `json: "email"`
	Phone string `json: "phone"`
}

// Guardar contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// codificacion o serializacion de datos
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Slice de contactos
	var contacts []Contact
	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	// instancia de bufio
	rader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar menu de opciones
		fmt.Println(" .....:::   GESTOR DE CONTACTOS   :::.....\n",
			"1. Agregar nuevo contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion: ")
		// Leer la oopcion seleccionada
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion", err)
			return
		}
		// Manejar la opcion del usuario
		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = rader.ReadString('\n')
			fmt.Print("Correo: ")
			c.Email, _ = rader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = rader.ReadString('\n')
			// Agregar un contacto a Slice
			contacts = append(contacts, c)
			// Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto")
			}
		case 2:
			// Mostrar todos los contactos
			fmt.Println("==================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("==================================================")
		case 3:
			// salir del programa
			return
		default:
			fmt.Println("Opcion invalida")
		}

	}
}
