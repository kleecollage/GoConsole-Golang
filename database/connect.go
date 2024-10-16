package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	// cadena de conexion a mysql
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// abrimos la conexion a la bd
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	// verificar la conexion
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Conexion a la BD MySql exitosa")

	return db, nil
}
