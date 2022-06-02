package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	db.Ping() // Verificamos la conexión

	fmt.Println(db.ExistsTable("users"))
	db.CreateTable(models.UserSchema, "users")

	db.Close()
	// Si la ejecutamos después de cerrar la BD nos dará un panic
	//db.Ping()
}
