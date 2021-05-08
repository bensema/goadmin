/*
CBC/ECB/CFB
ECB 256
*/
package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ====================== 256 pkcs5Padding
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cpher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cpher.BlockSize(); bs <= len(origData); bs, be = bs+cpher.BlockSize(), be+cpher.BlockSize() {
		cpher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cpher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cpher.BlockSize(); bs < len(encrypted); bs, be = bs+cpher.BlockSize(), be+cpher.BlockSize() {
		cpher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 32)
	copy(genKey, key)
	for i := 32; i < len(key); {
		for j := 0; j < 32 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte, err error) {
	encrypted = make([]byte, aes.BlockSize+len(origData))
	block, err := aes.NewCipher(key)
	if err != nil {
		return encrypted, err
	}
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return encrypted, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, nil
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte, err error) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}
