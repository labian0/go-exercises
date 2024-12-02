package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type ResponseData struct {
	// Define the fields according to the JSON response structure
	Key string `json:"key"`
}

func post(jsonData []byte, url string) ResponseData {
	// faire la requête avec les données json
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		return ResponseData{}
	}

	defer response.Body.Close()

	// lire le corps de la réponnse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ResponseData{}
	}

	// traiter le body
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return ResponseData{}
	}

	//renvoyer la donnée
}

func genererNombre() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(100)
	return randomNumber
}

func ex1() {
	// JEU DEVIN
	choix := genererNombre()

	json := []byte(fmt.Sprintf(`{"key":"%d"}`, choix))

	response := post(json, "https://url/")

	//TODO: traiter la réponse et trouver le nombre
}
