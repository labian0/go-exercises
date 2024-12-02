package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// mapping des données du json dans un struct bien gourmand
type ServerResponse struct {
	Message string `json:"message"`
}

func post(jsonData []byte) (ServerResponse, error) {
	// faire la requête avec les données json
	response, err := http.Post(SERVER_URL+"/ex1", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		return ServerResponse{}, err
	}

	defer response.Body.Close()

	// lire le corps de la réponnse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ServerResponse{}, err
	}

	// TODO unmarshal le body (parser les données brutes en un struct)

	// TODO renvoyer la struct
}

func genererNombre() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(10000)
	return randomNumber
}
