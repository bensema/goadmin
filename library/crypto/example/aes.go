/*
CBC/ECB/CFB
*/
package main

import (
	"encoding/base64"
	"encoding/hex"
	"library/crypto"
	"log"
)

func main() {
	origData := []byte("Hello World")                 // 待加密的数据
	key := []byte("9ce62c1836d128cfc875c9026db7564c") // 加密的密钥
	log.Println("原文：", string(origData))

	//log.Println("------------------ CBC模式 --------------------")
	//encrypted := crypto.AesEncryptCBC(origData, key)
	//log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	//log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	//decrypted := crypto.AesDecryptCBC(encrypted, key)
	//log.Println("解密结果：", string(decrypted))

	log.Println("------------------ ECB模式 --------------------")
	encrypted := crypto.AesEncryptECB(origData, key)
	log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := crypto.AesDecryptECB(encrypted, key)
	log.Println("解密结果：", string(decrypted))

	//log.Println("------------------ CFB模式 --------------------")
	//encrypted, _ = crypto.AesEncryptCFB(origData, key)
	//log.Println("密文(hex)：", hex.EncodeToString(encrypted))
	//log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))
	//decrypted, _ = crypto.AesDecryptCFB(encrypted, key)
	//log.Println("解密结果：", string(decrypted))
}
