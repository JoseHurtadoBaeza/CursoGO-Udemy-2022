package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/xml"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los usuarios")

	//rw.Header().Set("content-type", "application/json")
	rw.Header().Set("content-type", "text/xml")

	db.Connect()
	users := models.ListUsers()
	db.Close()
	//output, _ := json.Marshal(users) // Marshall devuelve 2 valores: Los valores transformados en tipo byte y un error
	output, _ := xml.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtiene un usuario")
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Actualiza un usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Elimina un usuario")
}
