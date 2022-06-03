package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los usuarios")
	db.Connect()
	users := models.ListUsers()
	db.Close()
	output, _ := json.Marshal(users) // Marshall devuelve 2 valores: Los valores transformados en tipo byte y un error
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
