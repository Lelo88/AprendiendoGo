package main

import "fmt"

type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
}

/*genero humano*/

type hombre struct{ //go detecta atributos y metodos de manera automatica
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	eshombre bool
}

type mujer struct{ //2. otra manera de hacerlo es asignando el struct hombre a mujer para que se parezca
	/*edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool*/
	hombre
}

func (h *hombre) respirar(){h.respirando = true}
func (h *hombre) comer(){h.comiendo = true}
func (h *hombre) pensar(){h.pensando = true}
func (h *hombre) sexo()string {
	if h.eshombre{
		return "Hombre"
	}else{
		return "Mujer"
	}}

func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("Soy un/a %s y estoy respirando\n", hu.sexo())
}

/*genero animal*/



type perro struct{
	respirando bool
	comiendo bool
	escarnivoro bool
}

func(p *perro) respirar() {p.respirando = true}
func(p *perro) comer() {p.comiendo = true}
func(p *perro) EsCarnivoro() bool {return p.escarnivoro}

func AnimalesRespirar(an animal){
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal)int{
	if an.EsCarnivoro(){
		return 1
	} 
	return 0
}

func main(){
	Pedro:=new(hombre)
	Pedro.eshombre = true
	Maria:=new(mujer)
	Maria.eshombre = false

	HumanosRespirando(Pedro)
	HumanosRespirando(Maria)

	fmt.Println("----------------------------------------------")
	totalCarnivoros :=0
	Dogo:= new(perro)
	Dogo.escarnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros=+AnimalesCarnivoros(Dogo)
	fmt.Println(totalCarnivoros)
}