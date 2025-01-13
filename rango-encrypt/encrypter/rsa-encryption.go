package encrypter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func RsaEncryption(byteContent []byte) []byte {
	// Public key string (PEM format)
	pubKeyStr := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1XJGE9mOqFYu/TOpxxST
c6XG7IaVzs245xwqTND1p/Ctm2du7Z/eciyp3ubTm+wl9QdzRj7fd4EnUAFuR6Av
eFna3guJxWxUHpgaYOZG2T+mSeQdBzVPvK3TZyx/Fwdc46tiOk7LIZPup1gv3Gw5
PwGJkYrH50ncZ+96gfD1QXH/NG55rX3KffhkeZfwPNaiveAv08OI8dlupxChQWWN
SSfpdVRXja+o14MWt/rPiiUgPRoue7+z2Q6xsAqdO87Q2s80ugP6ogllXMEVoxug
jwx5sCxWnT5Fi220lFwqpJwWQXlji5SE2WVvUKeRmXacM8B0WSxHjnYO4cx2NoZH
gQIDAQAB
-----END PUBLIC KEY-----`

	// Step 1: Decode the PEM block
	block, _ := pem.Decode([]byte(pubKeyStr))
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println("failed to decode PEM block containing public key")
	}

	// Step 2: Parse the DER-encoded public key
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse public key:", err)
	}

	// Step 3: Assert that it is an *rsa.PublicKey
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("not an RSA public key")
	}

	ciphertext, _ := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, byteContent)

	return ciphertext

}
