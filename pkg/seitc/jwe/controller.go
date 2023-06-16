package jwe

import (
	"crypto/rsa"
	"log"

	"github.com/GreenMan-Network/Go-KMS/pkg/rsakey"
	"github.com/GreenMan-Network/Go-Utils/pkg/fileutils"
	"github.com/go-jose/go-jose/v3"
)

var publicKeyPEMfile = "../../../assets/server_cert.pem"

func CreteJWE(data []byte) (string, error) {

	// Load public key
	publicKey, err := LoadPublicKey()
	if err != nil {
		log.Println("jwe.CreteJWE - Error loading public key: ", err)
		return "", err
	}

	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	encrypter, err := jose.NewEncrypter(jose.A128CBC_HS256, jose.Recipient{Algorithm: jose.RSA1_5, Key: publicKey}, nil)
	if err != nil {
		log.Println("jwe.CreteJWE - Error creating encrypter: ", err)
		return "", err
	}

	// Encrypt a sample plaintext. Calling the encrypter returns an encrypted
	// JWE object, which can then be serialized for output afterwards. An error
	// would indicate a problem in an underlying cryptographic primitive.
	object, err := encrypter.Encrypt(data)
	if err != nil {
		log.Println("jwe.CreteJWE - Error encrypting data: ", err)
		return "", err
	}

	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized, err := object.CompactSerialize()
	if err != nil {
		log.Println("jwe.CreteJWE - Error serializing data: ", err)
		return "", err
	}

	return serialized, nil
}

func LoadPublicKey() (*rsa.PublicKey, error) {

	// Load pem file
	pemFileString, err := fileutils.ReadFileToBytes(publicKeyPEMfile)
	if err != nil {
		log.Println("jwe.LoadPublicKey - Error reading pem file: ", err)
		return nil, err
	}

	// Load public key
	publicKey, err := rsakey.PemCertificateToPublicKey(string(pemFileString))
	if err != nil {
		log.Println("jwe.LoadPublicKey - Error parsing public key pem: ", err)
		return nil, err
	}

	return publicKey, nil
}
