package decrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
	"path/filepath"
)

func DecryptFile(aes32Key []byte, folderPath string, fileName string) {

	filePath := filepath.Join(folderPath, fileName)

	ciphertext, _ := os.ReadFile(filePath)

	block, err := aes.NewCipher(aes32Key)
	if err != nil {
		panic("Invalid key provided!")
	}
	if len(ciphertext) < aes.BlockSize {
		panic("Invalid key provided!")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// Move the decrypted file
	destinationPath := filepath.Join(folderPath, fileName[:len(fileName)-6])
	err = os.WriteFile(destinationPath, ciphertext, 0644)

	// Delete the encrypted file
	os.Remove(filePath)
}