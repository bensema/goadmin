package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func AesCBCEncrypt(rawData, key, iv []byte) ([]byte, error) {
	rawData = PKCS7Padding(rawData)
	ciphertext := make([]byte, len(rawData))
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, rawData)
	return ciphertext, err
}

func AesCBCDecrypt(rawData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(rawData, rawData)

	rawData = PKCS7UnPadding(rawData)
	return rawData, err
}
