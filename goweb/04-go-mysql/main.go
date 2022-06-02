package main

import "gomysql/db"

func main() {
	db.Connect()
	db.Ping() // Verificamos la conexión
	db.Close()
	// Si la ejecutamos después de cerrar la BD nos dará un panic
	//db.Ping()
}
