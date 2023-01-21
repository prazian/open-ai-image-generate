package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
	Size   string `json:"size"`
	N      int    `json:"n"`
}

type Image struct {
	Created int64  `json:"created"`
	Data    []Data `json:"data"`
}

type Data struct {
	URL string `json:"url"`
}

func main() {

	// Define the flags for the prompt and size parameters
	prompt := flag.String("p", "", "Prompt used for the image generation")
	size := flag.String("s", "", "Size of the image")
	flag.Parse()

	// Check that the prompt and size parameters are not empty
	if *prompt == "" || *size == "" {
		log.Fatal("prompt and size parameters are required")
	}

	// Set the API endpoint and API key
	endpoint := "https://api.openai.com/v1/images/generations"
	// Load the value of apiKey from a local file called token and remove the new line character
	token, err := os.ReadFile("./secrets/token")
	token = bytes.TrimSuffix(token, []byte("\n"))
	if err != nil {
		log.Fatal(err)
	}
	apiKey := string(token)

	// Create a new instance of the RequestBody struct and set the prompt and size
	body := RequestBody{Prompt: *prompt, Size: *size, N: 1}

	// Marshal the struct into a JSON string
	jsonStr, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	// Add the API key and model parameter to the request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("model", "image-alpha-001")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Request failed: ", resp.Status)
	}

	// Create a new JSON decoder
	dec := json.NewDecoder(resp.Body)

	// Unmarshal the response into an Image struct
	var image Image
	err = dec.Decode(&image)
	if err != nil {
		log.Fatal(err)
	}

	// Print the URL of the generated image
	fmt.Println(image.Data[0].URL)
}
