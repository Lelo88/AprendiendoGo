package main

import "fmt"

type Product struct {
	Id int64
	Name string
	Price float64
	Description string
	Category string
}

var products=[]Product{
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

func Save(product Product, products *[]Product){
	*products = append(*products, product)
}

func GetAll(products []Product){
	for _,product := range products{
		fmt.Println(product.Id, product.Name, product.Category)
	}
}

func main(){
	
	GetAll(products)
	newProduct:=Product{
		Id: 987,
		Name: "Azucar",
		Price: 26.6,
		Description: "",
		Category: "A",
	}

	Save(newProduct, &products)
	
	GetAll(products)
}