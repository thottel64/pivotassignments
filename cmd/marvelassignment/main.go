package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")

	client := marvelClient{
		baseURL:    "https://gateway.marvel.com/v1/public/",
		publickey:  publicKey,
		privatekey: privateKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	characters, err := client.getCharacters()
	if err != nil {
		log.Fatalln(err)
	}
	for _, character := range characters {
		fmt.Println(character.Name, character.Description)
	}
}

type marvelClient struct {
	baseURL    string
	publickey  string
	privatekey string
	httpClient *http.Client
}

func (c *marvelClient) getCharacters() ([]Character, error) {
	url := c.baseURL + "characters"
	response, err := c.httpClient.Get(c.urlSig(url))
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

func (c *marvelClient) md5hash(ts int64) string {
	tsForHash := strconv.Itoa(int(ts))
	hash := md5.Sum([]byte(tsForHash + c.privatekey + c.publickey))
	return hex.EncodeToString(hash[:])
}

func (c *marvelClient) urlSig(url string) string {
	ts := time.Now().Unix()
	hash := c.md5hash(ts)
	return fmt.Sprintf("%s?ts=%d&apikey=%s&hash=%s", url, ts, c.publickey, hash)
}
