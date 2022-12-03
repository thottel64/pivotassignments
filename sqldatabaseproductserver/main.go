package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("sqlite3", "products.db")
	if err != nil {
		fmt.Println("error opening database")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error pinging database")
		return
	}
	r := mux.NewRouter()
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/products", AddProduct).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	limit := r.FormValue("limit")
	query := "SELECT * FROM products LIMIT " + limit + ";"
	if limit == "" {
		query = "SELECT * FROM products;"
	}
	result, err := db.Query(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()
	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	intid, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if intid > 100 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	query := "SELECT * FROM products WHERE ID = " + stringid + ";"
	result := db.QueryRow(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var product Product
	err = result.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func AddProduct(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newproduct := Product{
		ID:    0,
		Name:  "",
		Price: 0,
	}

	err = json.Unmarshal(request, &newproduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := "INSERT INTO products (ID, Name, Price) VALUES (" + string(newproduct.ID) + ",`" + newproduct.Name + "`," + string(newproduct.Price) + ");"
	_, err = db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	id, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id > 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newproduct := Product{
		ID:    0,
		Name:  "",
		Price: 0,
	}

	err = json.Unmarshal(request, &newproduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := "UPDATE products SET Name = `" + newproduct.Name + "`, Price = " + string(newproduct.Price) + " WHERE ID =" + stringid
	_, err = db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	id, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if id > 100 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	query := "DELETE FROM products WHERE ID =" + stringid
	_, err = db.Exec(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
