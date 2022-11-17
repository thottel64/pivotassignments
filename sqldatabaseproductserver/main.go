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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
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
	err = seed(db)
	if err != nil {
		log.Fatalln(err)
	}
	err = seed(db)
	if err != nil {
		log.Fatalln(err)
	}
	db.Query("SELECT ID, Name, Price FROM products )")
}

func seed(db *sql.DB) error {
	//err := os.Remove("products.db")
	//if err != nil {
	//	return err
	//}
	sqlStatement := "CREATE TABLE products (ID integer primary key , Name ,Price REAL); delete from products"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("error with sqlstatement")
		return err
	}
	transaction, err := db.Begin()
	if err != nil {
		return err
	}
	statement, err := transaction.Prepare("INSERT INTO products(ID, Name, Price) VALUES(?,?,?) ")
	if err != nil {
		return err
	}
	defer statement.Close()
	for _, product := range products {
		_, err := statement.Exec(product.ID, product.Name, product.Price)
		if err != nil {
			return err
		}

	}
	err = transaction.Commit()
	return nil
}

//func query(db *sql.DB) error {
//	rows, err := db.Query("select ID from table ( products )")
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var id int
//		var name string
//		var price float64
//		err = rows.Scan(&id, &name, &price)
//		if err != nil {
//			return err
//		}
//		fmt.Println(id, name, price)
//	}
//	err = rows.Err()
//	if err != nil {
//		return err
//	}
//	statement, err := db.Prepare("select Name from products where ID = ?")
//	if err != nil {
//		return err
//	}
//	defer statement.Close()
//	var name string
//	err = statement.QueryRow("3").Scan(&name)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(name)
//
//	_, err = db.Exec("delete from products")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rows, err = db.Query("select ID, Name from products")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var id int
//		var name string
//		err = rows.Scan(&id, &name)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(id, name)
//	}
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//	return nil
//}

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
