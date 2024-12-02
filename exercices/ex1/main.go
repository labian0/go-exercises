package main

import "fmt"

const SERVER_URL string = "http://35.180.172.144:1337"

func main() {
	// JEU DEVIN
	choix := genererNombre()

	json := []byte(fmt.Sprintf(`{"guess":"%d"}`, choix))

	response, err := post(json)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	//TODO: traiter la réponse et trouver le nombre
}
