package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
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

	// the data is now unmarshaled into our struct
	r := mux.NewRouter()
	r.HandleFunc("/productss", GetHandler).Methods("GET")
	r.HandleFunc("/productss/{id}", GetIDHandler).Methods("GET")
	r.HandleFunc("/productss", PostHandler).Methods("POST")
	r.HandleFunc("/productss/{id}", PutHandler).Methods("PUT")
	r.HandleFunc("/productss/{id}", DeleteHandler).Methods("DELETE")
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func InitProducts() {
	// step 1: use the HTTP package to get the data and defer the request to prevent a data leak
	response, err := http.Get("https://gist.githubusercontent.com/jboursiquot/259b83a2d9aa6d8f16eb8f18c67f5581/raw/9b28998704fb06f127f13540a4f6e3812f50774b/productss.json")
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
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intid := 0
	_, err := fmt.Sscan(id, &intid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if intid > len(products) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, productss := range products {

		if productss.ID == intid {
			response, err := json.Marshal(products[intid-1])
			if err != nil {
				log.Print("error 3")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)

			_, err = w.Write(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	var newProduct Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()
	err = json.Unmarshal(reqBody, &newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)

}
func PutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intid := 0
	_, err := fmt.Sscan(id, &intid)
	if err != nil {
		log.Fatalln(err)
	}
	if intid > len(products) {
		w.WriteHeader(http.StatusNotFound)
		response := "404 status not found"
		_, err = w.Write([]byte(response))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for index, p := range products {
		if p.ID == intid {
			products = append(products[:index], products[:index+1]...)
			err = json.NewDecoder(r.Body).Decode(&products)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			p.ID = intid
			products = append(products, p)
			json.NewEncoder(w).Encode(products)
			return
		}
	}
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intid := 0
	_, err := fmt.Sscan(id, &intid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if intid > len(products) {
		w.WriteHeader(http.StatusNotFound)
		response := "404 status not found"
		_, err = w.Write([]byte(response))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for index, p := range products {
		if p.ID == intid {
			products = append(products[:index], products[index+1:]...)
			w.WriteHeader(http.StatusOK)
		}
	}
}
