package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"
)

func main() {
	// Estalecer conexion a la bd
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Selecciona una opcion: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Print("Ingrese el ID del contacto que desea obtener: ")
			var contactId int
			fmt.Scanln(&contactId)
			handlers.GetContactById(db, contactId)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Print("Ingrese el ID del contacto que desea borrar: ")
			var contactId int
			fmt.Scanln(&contactId)
			handlers.DeleteContact(db, contactId)
			handlers.ListContacts(db)
		case 6:
			fmt.Println("Chupeme")
			return
		default:
			fmt.Println("Esa opcion es digna de un bautista")
		}
	}
}

func inputContactDetails(option int) models.Contact {
	// Leer la entrada del usuario usando bufio
	reader := bufio.NewReader(os.Stdin)
	var contact models.Contact

	if option == 4 {
		fmt.Print("Ingrese el ID del contacto que desea actualizar: ")
		var contactId int
		fmt.Scanln(&contactId)

		contact.Id = contactId
	}

	fmt.Print("Ingrese el nombre: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electronico: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el numero telefonico: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
