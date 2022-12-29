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

/*func getByID(id int) (prod *Products){
	
}*/

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
}