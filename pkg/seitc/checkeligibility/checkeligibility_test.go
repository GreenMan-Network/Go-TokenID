package checkeligibility

import (
	"log"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("Hello World")

	err := HelloWorldRequest()
	if err != nil {
		t.Error("Error in HelloWorldRequest: ", err)
	}
}

func TestCheckEligibilityRequestToJson(t *testing.T) {
	t.Log("TestCheckEligibilityRequestToJson")

	testRequest := CheckEligibilityRequest{
		ConsumerID:     "wiremockConsumerUserId1",
		ConversationID: "d432561f-9ba1-4131-86d3-d02fb21fa2c4",
		EnrollDevice:   true,
		DeviceData: DeviceData{
			DeviceID:          "36af69220a9547cf95b624ae",
			DeviceName:        "MyPhone",
			DevicePlatform:    "ANDROID",
			DeviceType:        "PHONE",
			Language:          "",
			Locale:            "en-US",
			NfcCapable:        true,
			OsVersion:         "Android 5.0.3",
			WalletID:          "ccd406a9-cda4-47a4-8ef7-d96f08d1ebc7",
			DevicePhoneNumber: HelperFunction("6506198963"),
			ImeiNumber:        HelperFunction("930850333"),
		},
		EncData:          "eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiUlNBMV81In0.N9maegHiu-smJRRb0zR3a-ppflWKo1IOUKEc7Z-QDikv1o5OAn5gZal7EhsikFEj7gRidNP_LmFdhDGlp3-h1QXld5sunTeipJo9fQGylZ3D1naOajxn2Z0msnRFaRdsZ2PcxtPDR9CBy-0h81qpCcdm893xAYgJWH315HIFxMYUuYJWRsXs9XLyHxcQDQoPH2B4us0iCXEjA6E7_h8UPxVsRnpupC92qq6xXhM_uFGe3D112LjPfnCIJxO8pDbbPKXdh_YtnaUuN5P5GS5mJk-q-D34JIOtBK5ZWzZB1WozOay2Kj9gU34twffmAuvg6FROaHXIYJnAZoqw0p4a2w.PZ5doswn3zRr28rsKW2-7w.BhGnF_Dnz9RSLW70Gq-TJfZn2D9b47kJXB0HiQZ7Lbhl4_k8wHSi9rnQn0z5lEJny4H7J2AZqbU3jtjZJUptSllorR_XtXSY4TgCOYEAvX5IWJ1NbZp7_mHDnaMP0DTp65o415jYMVzwkoW0QzDOD4i88bmSkGBBH6JUz38A_mlV13cuFTMTb2IFyhzEoCLZq_SVaUGNgCbIzTSzENzzJUrAJkloNVCI5oVyT4agGHBM3QJjmOki4yGKhP5jCzRJr_ZAGptM-DN8GgyNV4Sg5seBTPCB05EhEePdx3dHXGQMCv8tRxsipFnd4eC3Xsf_NTK6DGGqw8QW0csWiD383AxF9o_jtaRz9JR2_FNRiMGzj3_NcTjs2rYDEiQJNkW_1DLITl1uJFNIl2fQFc4mvANHwVHxvqQTQ16TTtz6O0zmBMcmnLDITA3A3sMD08Yc6YiAxZ8WxtUTITlOea90TolqvqXb5mYbrg-HgkuXlMi-TdwW2656VuoQsmZQuXhP8i7kEJpaVbpjc8tRUjWSSwf3J64nMKK11tIKxAQJe4FtRoaQN77mpe96WokDzw7e.9fTTrmfheGFPAGMZrlSbXQ",
		MessageID:        "30e1bda2-75ca-41c0-9135-cc9f06158c4c",
		TokenLocation:    "DEVICE_MEMORY",
		PresentationMode: "NFCHCE",
		TokenRequestorID: "77745896114",
		PanSource:        "KEYENTERED",
	}

	requestJson, err := testRequest.ToJson()
	if err != nil {
		t.Error("Error converting request to json: ", err)
		return
	}

	log.Println("Request: ", string(requestJson))

}

func TestCheckEligibilityResponseToJson(t *testing.T) {
	t.Log("TestCheckEligibilityRequestToJson")

	var responseObj CheckEligibilityResponse

	responseObj.ConversationID = "d432561f-9ba1-4131-86d3-d02fb21fa2c4"

	responseJson, err := responseObj.ToJson()

	if err != nil {
		t.Error("Error converting response to json: ", err)
		return
	}

	t.Log("Response object !!!: ", string(responseJson))
}

func TestRequest(t *testing.T) {
	t.Log("TestRequest")

	testEncData := EncDataOfCheckEligibilityRequest{
		EncPanData: PanData{
			Pan:        "4514231274459132",
			PanExpDate: "1230",
			Cvv:        "123",
		},
		EncCardholderData: &CardholderData{
			CardholderName: HelperFunction("Cody Payne"),
			BillingAddress: &BillingAddress{
				Line1:       HelperFunction("389 North Mango Avenue"),
				City:        HelperFunction("Buffalo"),
				State:       HelperFunction("OR"),
				PostalCode:  HelperFunction("04009"),
				CountryCode: HelperFunction("US"),
			},
		},
	}

	jweString, err := testEncData.GetJWE()
	if err != nil {
		t.Error("Error creating jwe: ", err)
		return
	}

	testRequest := CheckEligibilityRequest{
		ConsumerID:     "wiremockConsumerUserId1",
		ConversationID: "d432561f-9ba1-4131-86d3-d02fb21fa2c4",
		EnrollDevice:   true,
		DeviceData: DeviceData{
			DeviceID:          "36af69220a9547cf95b624ae",
			DeviceName:        "MyPhone",
			DevicePlatform:    "ANDROID",
			DeviceType:        "PHONE",
			Language:          "",
			Locale:            "en-US",
			NfcCapable:        true,
			OsVersion:         "Android 5.0.3",
			WalletID:          "ccd406a9-cda4-47a4-8ef7-d96f08d1ebc7",
			DevicePhoneNumber: HelperFunction("6506198963"),
			ImeiNumber:        HelperFunction("930850333"),
		},
		EncData:          jweString,
		MessageID:        "30e1bda2-75ca-41c0-9135-cc9f06158c4c",
		TokenLocation:    "DEVICE_MEMORY",
		PresentationMode: "NFCHCE",
		TokenRequestorID: "10006345780",
		PanSource:        "KEYENTERED",
	}

	response, err := CheckEligibility(&testRequest)
	if err != nil {
		t.Error("Error checking eligibility: ", err)
		return
	}

	log.Println("Response Obj: ", *response)

	reponseJson, err := response.ToJson()
	if err != nil {
		t.Error("Error converting response to json: ", err)
		return
	}

	log.Println("Response: ", reponseJson)
}

func TestEncryptDataOfCheckEligibilityRequest(t *testing.T) {
	testData := EncDataOfCheckEligibilityRequest{
		EncPanData: PanData{
			Pan:        "4514231274459132",
			PanExpDate: "1230",
			Cvv:        "123",
		},
		EncCardholderData: &CardholderData{
			CardholderName: HelperFunction("Cody Payne"),
			BillingAddress: &BillingAddress{
				Line1:       HelperFunction("389 North Mango Avenue"),
				City:        HelperFunction("Buffalo"),
				State:       HelperFunction("OR"),
				PostalCode:  HelperFunction("04009"),
				CountryCode: HelperFunction("US"),
			},
		},
	}

	jweString, err := testData.GetJWE()
	if err != nil {
		t.Error("Error creating jwe: ", err)
		return
	}

	t.Log("JWE: ", jweString)
}
