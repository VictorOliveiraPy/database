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

	product := NewProduct("CASA", 4889.94)
	err = insertProduct(db, product)
	if err!= nil {
        panic(err)
    }
	fmt.Println("Product inserted successfully")

	product.Price = 100.2
	err = updateProduct(db, product)
	if err!= nil {
        panic(err)
    }
	fmt.Print("Product updated successfully")

	product, err = selectProduct(db, product.ID)
	if err!= nil {
        panic(err)
    }
	fmt.Println("Product retrieved successfully")
	fmt.Println(product.ID)
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

func updateProduct(db *sql.DB, product *Product) error{
	stmt, err := db.Prepare("update products set name = ?, price = ? where id =?")
    if err!= nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(product.Name, product.Price, product.ID)
    if err!= nil {
        return err
    }
    return nil

}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id =?")
    if err!= nil {
        return nil, err
    }
    defer stmt.Close()

    var product Product
    err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
    if err!= nil {
        return nil, err
    }
    return &product, nil

}