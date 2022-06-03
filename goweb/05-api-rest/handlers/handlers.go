package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Lista todos los usuarios")

	rw.Header().Set("content-type", "application/json") // Para responder con json
	//rw.Header().Set("content-type", "text/xml") // Para responder con xml
	// No hay un tipo de dato específico para yaml

	db.Connect()
	users := models.ListUsers()
	db.Close()
	// Marshall devuelve 2 valores: Los valores transformados en tipo byte y un error
	output, _ := json.Marshal(users) // Para responder con json
	//output, _ := xml.Marshal(users) // Para responder con xml
	//output, _ := yaml.Marshal(users) // Para responder con yaml
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Obtiene un usuario")

	rw.Header().Set("content-type", "application/json") // Para responder con json
	//rw.Header().Set("content-type", "text/xml") // Para responder con xml
	// No hay un tipo de dato específico para yaml

	// Obtener ID
	vars := mux.Vars(r) // Mos devuelve un mapa de tipo string
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userId)
	db.Close()
	// Marshall devuelve 2 valores: Los valores transformados en tipo byte y un error
	output, _ := json.Marshal(user) // Para responder con json
	//output, _ := xml.Marshal(users) // Para responder con xml
	//output, _ := yaml.Marshal(users) // Para responder con yaml
	fmt.Fprintln(rw, string(output))

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
