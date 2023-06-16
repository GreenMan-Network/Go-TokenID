package visa

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var (
	certFile = flag.String("cert", "../../../assets/visasecurity/cert.pem", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "../../../assets/visasecurity/privateKey.pem", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "../../../assets/visasecurity/DigiCertGlobalRootCA.pem", "A PEM eoncoded CA's certificate file.")
	userID   = flag.String("User ID", "Basic VVkySFJKTkZBOFNKU0tRSDZLSUEyMUNBTlh4cGluNzdyVkNUbkY4SE84TmxiOWtUazp3WmQ3M0JBNXJZTHZtRTVXV0w=", "User ID provided by Visa.")
)

var client *http.Client

func loadClient() error {
	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Println("visa.loadClient - Error loading client cert: ", err)
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile(*caFile)
	if err != nil {
		log.Println("visa.loadClient - Error reading CA cert file: ", err)
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	// in the http client configuration, add TLS configuration and add the RootCAs
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	return nil
}

// Executes a request adding all the necessary data for Visa authentication (certificates, etc)
func ExecuteRequest(req *http.Request) (*http.Response, error) {

	if client == nil {
		err := loadClient()
		if err != nil {
			log.Println("visa.ExecuteRequest - Error loading client: ", err)
			return nil, err
		}
	}

	req.Header.Add("Authorization", *userID)

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Println("visa.ExecuteRequest - Error dumping request: ", err)
	} else {
		log.Println("visa.ExecuteRequest - POST REQUEST:\n", string(requestDump))
	}

	// Do GET something
	resp, err := client.Do(req)
	if err != nil {
		log.Println("visa.ExecuteRequest - Error executing request: ", err)
		return nil, err
	}

	return resp, nil
}
