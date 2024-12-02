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

// mapping des données du json dans un struct bien gourmand
type ResponseData struct {
	Key string `json:"key"`
}

func post(jsonData []byte, url string) (ResponseData, error) {
	// faire la requête avec les données json
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		return ResponseData{}, err
	}

	defer response.Body.Close()

	// lire le corps de la réponnse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ResponseData{}, err
	}

	// traiter le body
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return ResponseData{}, err
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

	response, err := post(json, "https://url/")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	//TODO: traiter la réponse et trouver le nombre
}
