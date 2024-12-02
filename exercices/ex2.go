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

// checker si un nombre est premier
func estPremier(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// generer un nombre premier de n chiffres
func genererPremier(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := intPow(10, n-1)
	max := intPow(10, n) - 1

	for {
		num := r.Intn(max-min+1) + min // Generate a number in the range [min, max]
		if estPremier(num) {
			return num
		}
	}
}

// calculer la puissance d'un nombre
func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func main() {
	choix := genererPremier(19)
	fmt.Println(choix)
	/*
	   json := []byte(fmt.Sprintf(`{"key":"%d"}`, choix))

	   response := post(json, "https://url/")
	*/
}
