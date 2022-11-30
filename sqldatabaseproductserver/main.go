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

var products []Product

func main() {
	InitProducts()
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	products, err = seed(db)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(products)
}

func seed(db *sql.DB) ([]Product, error) {
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
		return err
	}
	rows, err := db.Query("SELECT ID, Name, Price FROM products LIMit 5)")
	if err != nil {
		fmt.Println("error with query")
		return nil, err
	}
	fmt.Println(rows)
	return products, nil
}

func InitProducts() {
	// step 1: use the HTTP package to get the data and defer the request to prevent a data leak
	response, err := http.Get("https://gist.githubusercontent.com/jboursiquot/259b83a2d9aa6d8f16eb8f18c67f5581/raw/9b28998704fb06f127f13540a4f6e3812f50774b/products.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	//step 2 read the data and save it to a slice of bytes
	byteslice, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//step 3 unmarshal the slice of bytes into a struct
	err = json.Unmarshal(byteslice, &products)
	if err != nil {
		log.Fatalln(err)
	}
	// step 4 print out the data
	log.Printf("%s", response.Status)

}
