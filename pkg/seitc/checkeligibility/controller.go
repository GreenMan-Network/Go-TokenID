package checkeligibility

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/GreenMan-Network/Go-TokenID/pkg/seitc"
	"github.com/GreenMan-Network/Go-TokenID/pkg/seitc/jwe"
	"github.com/GreenMan-Network/Go-TokenID/pkg/visa"
)

var endpoint = "/unified/tr/v1/checkeligibility"

//var endpoint = "post"

func (str *CheckEligibilityRequest) ToJson() ([]byte, error) {
	return json.Marshal(str)
}

func (str *CheckEligibilityResponse) ToJson() ([]byte, error) {
	return json.Marshal(str)
}

func (str *EncDataOfCheckEligibilityRequest) ToJson() ([]byte, error) {
	return json.Marshal(str)
}

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

// Execute a check eligibility request
func CheckEligibility(request *CheckEligibilityRequest) (*CheckEligibilityResponse, error) {

	responseBytes, err := seitc.HttpPostRequest(endpoint, request)
	if err != nil {
		log.Println("checkeligibility.CheckEligibility - Error making POST request: ", err)
		return nil, err
	}

	log.Println("Response string: ", string(responseBytes))

	var responseObj *CheckEligibilityResponse

	err = json.Unmarshal(responseBytes, responseObj)
	if err != nil {
		log.Println("checkeligibility.CheckEligibility - Error unmarshalling response: ", err)
		return nil, err
	}

	return responseObj, nil
}

func HelperFunction(val string) *string {
	return &val
}

// Encrypt the data of a check eligibility request into a jwe, using the server_cert.pem public key
func (encDataObj *EncDataOfCheckEligibilityRequest) GetJWE() (string, error) {
	objJson, err := encDataObj.ToJson()
	if err != nil {
		log.Println("checkeligibility.EncDataOfCheckEligibilityRequest.GetJWE - Error converting object to json: ", err)
		return "", err
	}

	jweString, err := jwe.CreteJWE(objJson)
	if err != nil {
		log.Println("checkeligibility.EncDataOfCheckEligibilityRequest.GetJWE - Error generating jwe: ", err)
	}

	return jweString, nil
}
