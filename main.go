package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //_ ignora deixando a gente compilar mesmo sem esta sendo usado visualmente
	"github.com/google/uuid"
)


type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
		}
	}


func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
    if err!= nil {
        panic(err)
    }
    defer db.Close()

	product := NewProduct("IPHONE", 4889.94)
	err = insertProduct(db, product)
	if err!= nil {
        panic(err)
    }
	fmt.Println("Product inserted successfully")


}


func insertProduct(db *sql.DB,  product *Product) error{
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err!= nil {
        return err
    }
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err!= nil {
        return err
    }
	return nil

}

