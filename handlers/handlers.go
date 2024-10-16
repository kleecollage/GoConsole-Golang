package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"
	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("----------------------------------------------------------------------")
	for rows.Next() {
		// Instancia del modelo contact
		contact := models.Contact{}

		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}
		// validando valor nulo de email
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "xxxxxxxxxx"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("----------------------------------------------------------------------")
	}
}

func GetContactById(db *sql.DB, contactID int) {
	// Consulta
	query := "SELECT * FROM contact WHERE id = ?"
	row := db.QueryRow(query, contactID)
	// Instancia del modelo contact
	contact := models.Contact{}
	var valueEmail sql.NullString
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No existen contactos con el ID: ", contactID)
		}
		log.Fatal(err)
	}
	// validando posible valor nulo de email
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "xxxxxxxxxx"
	}
	fmt.Println("\nCONTACTO:")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("----------------------------------------------------------------------")
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES(?, ?, ?)"
	// Ejecutar sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Nuevo contacto registrado con exito")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto actualizado con exito")
}

func DeleteContact(db *sql.DB, contactID int) {
	// Consulta
	query := "DELETE FROM contact WHERE id = ?"
	// Ejecutar consulta
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto eliminado con exito")
}
