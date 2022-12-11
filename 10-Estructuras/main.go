package main

import (
		"fmt" 
		"time"

		usuario "./Usuario" // manera de importar paquetes creados
)

type pepe struct{ //herencia
	usuario.Usuario
}


func main(){
	user := new(pepe)
	user.AltaUsuario(1,"Leandro Villalba", time.Now(), true)
	fmt.Println(user.Usuario)
}

//tuve que correr el comando go env -w GO111MODULE=auto para que pueda leer el modulo