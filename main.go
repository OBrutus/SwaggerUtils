package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const PrintPath = false

func main() {
	fmt.Print("Enter the swagger endpoint URL for getting count: ")
	var swaggerUrl string
	fmt.Scanln(&swaggerUrl)

	resp, err := http.Get(swaggerUrl)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var keyValue map[string]interface{}
	err = json.Unmarshal(body, &keyValue)
	if err != nil {
		log.Fatalln(err)
	}

	if keyValue["paths"] == nil {
		log.Fatalln("Invalid swagger endpoint")
		os.Exit(404)
	}

	pathMap := keyValue["paths"].(map[string]interface{})

	if PrintPath {
		printPath(pathMap)
	}

	// Only endpoints are counted here, irrespective of the methods, like GET, POST, etc
	length := len(keyValue["paths"].(map[string]interface{}))
	fmt.Println("length = ", length)
}

func printPath(pathMap map[string]interface{}) {
	for k, v := range pathMap {
		fmt.Printf("%s: %s \n", k, v)
	}
}
