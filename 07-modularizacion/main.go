// Un detalle a tener en cuenta es que si evitamos usar guiones en las rutas la importación de paquetes se hará de forma más automatizada
// y sino lo que tenemos que hacer es lo que hemos hecho aquí, que básicamente consiste en ejecutar el comando go mod init "nombre del módulo gestor de paquetes"
// y tras esto sólamente tendremos que usar los import cómo se puede ver en este ejemplo.
package main

import "paquetes/mensajes"

func main() {
	mensajes.Hola()
}
