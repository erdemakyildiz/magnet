package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HttGet(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)

		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)

		return "", err
	}

	var response = string(body)

	return response, nil
}
