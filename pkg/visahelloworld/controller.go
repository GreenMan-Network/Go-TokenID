package visahelloworld

import (
	"io"
	"log"
	"net/http"

	"github.com/GreenMan-Network/Go-TokenID/pkg/visa"
)

func HelloWorldRequest() error {

	req, err := http.NewRequest(http.MethodGet, "https://sandbox.api.visa.com/vdp/helloworld", nil)
	if err != nil {
		log.Println("visahelloworld.HelloWorldRequest - Error creating the POST request: ", err)
		return err
	}

	// Do GET something
	resp, err := visa.ExecuteRequest(req)
	if err != nil {
		log.Println("visahelloworld.HelloWorldRequest - Error making POST request: ", err)
		return err
	}
	defer resp.Body.Close()

	// Dump response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))

	return nil
}
