package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	db.Ping() // Verificamos la conexión

	db.CreateTable(models.UserSchema)

	db.Close()
	// Si la ejecutamos después de cerrar la BD nos dará un panic
	//db.Ping()
}
