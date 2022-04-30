package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {

	jose := User{
		nombre: "Jose",
		email:  "jose@gmail.com",
		activo: true,
	}

	carlos := User{
		nombre: "Carlos",
		email:  "carlos@gmail.com",
		activo: true,
	}

	joseStudent := Student{
		user:   jose,
		codigo: "001arsd",
	}

	fmt.Println(carlos)
	fmt.Println(joseStudent.user.nombre)

}
