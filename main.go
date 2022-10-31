package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

var product []Product

func main() {
	InitProducts()

	// the data is now unmarshaled into our struct
	r := mux.NewRouter()
	r.HandleFunc("/products", GetHandler).Methods("GET")
	r.HandleFunc("/products/{id}", GetIDHandler).Methods("GET")
	r.HandleFunc("/products", PostHandler).Methods("POST")
	r.HandleFunc("/products/{id}", PutHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteHandler).Methods("DELETE")
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
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
	err = json.Unmarshal(byteslice, &product)
	if err != nil {
		log.Fatalln(err)
	}
	// step 4 print out the data
	log.Printf("%s", response.Status)
	fmt.Println(len(product))
	o := product[1]
	fmt.Println(o)
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(product)
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
		log.Fatalln(err)
	}
	for _, products := range product {
		if intid > len(product) {
			w.WriteHeader(http.StatusNotFound)
			response := "404 status not found"
			_, err = w.Write([]byte(response))
			return
		}
		if products.ID == intid {
			response, err := json.Marshal(product[intid-1])
			if err != nil {
				log.Print("error 3")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)

			_, err = w.Write(response)
		}
	}
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	// these fields need to be changed to match our JSON struct
	ID := r.FormValue("ID")
	Name := r.FormValue("Name")
	Description := r.FormValue("Description")
	Price := r.FormValue("Price")
	fmt.Fprintf(w, "ID = %d\n", ID)
	fmt.Fprintf(w, "Name = %s\n", Name)
	fmt.Fprintf(w, "Description = %s\n", Description)
	fmt.Fprintf(w, "Price = %d\n", Price)

}
func PutHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	}
	response := "This is a put request"
	_, err = w.Write([]byte(response))
	fmt.Println("r now", r)
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.PostForm", r.PostForm)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intid := 0
	_, err := fmt.Sscan(id, &intid)
	if err != nil {
		log.Fatalln(err)
	}
	for _, products := range product {
		if intid > len(product) {
			w.WriteHeader(http.StatusNotFound)
			response := "404 status not found"
			_, err = w.Write([]byte(response))
			return
		}
		if products.ID == intid {
			products.ID = intid
			products.Price = 0
			products.Description = ""
			products.Name = ""
			response := "Product Deleted"
			if err != nil {
				log.Print("error 3")
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(response))
		}
	}
}
