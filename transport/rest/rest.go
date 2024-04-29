package rest

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func Post(url string, input []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(input))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("An error occurred when sending POST request : %s", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("An error occurred reading the response body : %s", err)
		return nil, err
	}

	err = resp.Body.Close()
	return body, nil
}
