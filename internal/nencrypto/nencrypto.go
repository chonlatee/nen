package nencrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKeyPEM() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	pk, err := encodePrivateKeyToPEM(privateKey)
	if err != nil {
		return "", "", err
	}

	publicKey := privateKey.PublicKey

	pb, err := encodePublicKeyToPEM(&publicKey)
	if err != nil {
		return "", "", err
	}

	return pk, pb, err

}

func encodePrivateKeyToPEM(key *rsa.PrivateKey) (string, error) {
	k, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return "", err
	}

	p := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: k,
	}

	r := pem.EncodeToMemory(p)

	return string(r), nil
}

func encodePublicKeyToPEM(key *rsa.PublicKey) (string, error) {
	k, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}

	p := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: k,
	}

	r := pem.EncodeToMemory(p)

	return string(r), nil
}
