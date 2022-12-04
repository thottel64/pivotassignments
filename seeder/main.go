package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	products, err := InitProducts()
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	products, err = seed(db, products)
	if err != nil {
		return
	}
	fmt.Println(products)
}

func seed(db *sql.DB, products []Product) ([]Product, error) {
	transaction, err := db.Begin()
	if err != nil {
		return nil, err
	}
	statement, err := transaction.Prepare("INSERT INTO products(ID, Name, Price) VALUES(?,?,?) ")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	for _, product := range products {
		_, err := statement.Exec(product.ID, product.Name, product.Price)
		if err != nil {
			return nil, err
		}

	}
	err = transaction.Commit()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT ID, Name, Price FROM products LIMIT 5)")
	if err != nil {
		fmt.Println("error with query")
		return nil, err
	}
	fmt.Println(rows)
	return products, nil
}

func InitProducts() ([]Product, error) {
	var products []Product
	response, err := http.Get("https://gist.githubusercontent.com/jboursiquot/259b83a2d9aa6d8f16eb8f18c67f5581/raw/9b28998704fb06f127f13540a4f6e3812f50774b/products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	byteslice, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(byteslice, &products)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s", response.Status)

	return products, err
}
