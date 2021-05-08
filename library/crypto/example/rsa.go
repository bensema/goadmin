package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"library/crypto"
	"time"
)

func main() {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey, _ := crypto.GenRsaKey()
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData, _ := crypto.RsaSignWithSha256([]byte(data), prvKey)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if b, _ := crypto.RsaVerySignWithSha256([]byte(data), signData, pubKey); b == true {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}
	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext, _ := crypto.RsaEncrypt([]byte(`0000116156214087293FlVUfCV{"nonce":"tOGrhLMH","tmpAes":"fz0wKojksh9DSu6TSOhmmTal2QakfOJ1"}`), pubKey)
	fmt.Println(len(ciphertext))
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	start := time.Now()
	sourceData, _ := crypto.RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
	fmt.Println(time.Now().Sub(start))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
