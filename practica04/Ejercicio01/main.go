package main

import "fmt"

//struct Product
type Product struct {
	Id int64
	Name string
	Price float64
	Description string
	Category string
}

//defino un slice con el struct Product
type Products []Product 

var products=Products{
	{
	   Id: 987,
	   Name: "Manzana",
	   Price: 45.6,
	   Description: "",
	   Category: "F",
   },
   {
	   Id: 345,
	   Name: "Carne",
	   Price: 100.53,
	   Description: "Carne de vaca",
	   Category: "C",
   },
}

func (prods *Product) Save(){
	products = append(products, *prods)
}

func (prods Products) GetAll(){
	for _,product := range prods{
		fmt.Println(product.Id, product.Name, product.Category)
	}
}

func getByID(id int64) (ide int64, message string){
	if id<=0{
		id = ide
		message = "El Id ingresado no es correcto"
		return ide, message 
	}

	for _,prod := range products {
		if prod.Id == id{
			ide = id
			message = prod.Name
			return ide, message
		} else {
			message = "No existe ese producto"
			return ide, message
		}
	}
	return ide, message
}

func main(){
	
	products.GetAll()
	newProduct:=Product{
		Id: 456,
		Name: "Azucar",
		Price: 26.6,
		Description: "",
		Category: "A",
	}

	newProduct.Save()

	products.GetAll()

	otroProducto, mensaje := getByID(45)
	otroProducto2, mensaje2 := getByID(987)
	fmt.Println(otroProducto, mensaje)
	fmt.Println(otroProducto2, mensaje2)
}