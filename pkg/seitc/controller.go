package seitc

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/GreenMan-Network/Go-TokenID/pkg/visa"
)

var url = "https://cert.api.visa.com"

//var url = "https://httpbin.org/"

func HttpPostRequest(endpoint string, requestData RequestData) ([]byte, error) {
	requestUrl := url + endpoint

	requestJson, err := requestData.ToJson()
	if err != nil {
		log.Println("seitc.HttpPostRequest - Error marshalling request to JSON: ", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(requestJson))
	if err != nil {
		log.Println("seitc.HttpPostRequest - Error creating the POST request: ", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := visa.ExecuteRequest(req)
	if err != nil {
		log.Println("seitc.HttpPostRequest - Error making POST request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	return parseResponse(resp)
}

// Verify the response and extract the body
func parseResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("seitc.parseResponse - Error reading response body: ", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("seitc.parseResponse - Http status code: " + resp.Status)
		log.Println("seitc.parseResponse - Response message: " + string(body))

		return nil, fmt.Errorf(ErrHttpStatusNot2xx.Error() + resp.Status)
	}

	if len(body) == 0 {
		return nil, nil
	}

	return body, nil
}
