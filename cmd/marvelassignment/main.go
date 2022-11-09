package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("/Users/taylor.hottel/pivotassignments/cmd/marvelassignment/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")
	fmt.Println(publicKey, privateKey)

	client := marvelClient{
		publickey:  publicKey,
		privatekey: privateKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

}

type marvelClient struct {
	publickey  string
	privatekey string
	httpClient *http.Client
}

func (c *marvelClient) getCharacters() ([]Character, error) {
	response, err := c.httpClient.Get("https://gateway.marvel.com:443/v1/public/characters")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var characterResponse CharacterResponse
	err = json.NewDecoder(response.Body).Decode(&characterResponse)
	if err != nil {
		return nil, err
	}
	return characterResponse.Data.Results, nil
}

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CharacterResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int         `json:"offset"`
		Limit   int         `json:"limit"`
		Total   int         `json:"total"`
		Count   int         `json:"count"`
		Results []Character `json:"results"`
	} `json:"data"`
}
