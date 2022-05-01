// Un detalle a tener en cuenta es que si evitamos usar guiones en las rutas la importación de paquetes se hará de forma más automatizada
// y sino lo que tenemos que hacer es lo que hemos hecho aquí, que básicamente consiste en ejecutar el comando go mod init "nombre del módulo gestor de paquetes"
// y tras esto sólamente tendremos que usar los import cómo se puede ver en este ejemplo.
package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	/*mensajes.Hola()
	mensajes.Imprimir()*/
	/*cua1 := figuras.Cuadrado{Lado: 8}
	cir1 := figuras.Circulo{Radio: 5}

	figuras.Medidas(&cua1)
	figuras.Medidas(&cir1)*/

	p1 := models.Persona{}
	p1.Constructor("Jose", 25)
	fmt.Println(p1)
	fmt.Println(p1.GetNombre())
	p1.SetNombre("Carlos")

	fmt.Println(p1)
}
