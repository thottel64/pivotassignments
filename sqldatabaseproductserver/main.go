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

func main() {
	var err error
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
	stmt, err := db.Prepare("SELECT * FROM products LIMIT ?;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := stmt.Query(limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = stmt.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	stmt, err := db.Prepare("SELECT * FROM products WHERE ID =?;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := stmt.QueryRow(intid)
	err = stmt.Close()
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
	newproduct := Product{}

	err = json.Unmarshal(request, &newproduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stmt, err := db.Prepare("INSERT INTO products (ID, Name, Price) VALUES (?,?,?);")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = stmt.Exec(newproduct.ID, newproduct.Name, newproduct.Price)
	w.WriteHeader(http.StatusBadRequest)

}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	id, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newproduct := Product{}

	err = json.Unmarshal(request, &newproduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("UPDATE products SET Name =?, Price =?, WHERE ID =?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(newproduct.Name, newproduct.Price, newproduct.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringid := vars["id"]
	id, err := strconv.Atoi(stringid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stmt, err := db.Prepare("DELETE FROM products WHERE ID =?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
