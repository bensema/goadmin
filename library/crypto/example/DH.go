package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"library/crypto"
	"math/big"
	"time"
)

// DH密钥协商算法
func main() {
	start := time.Now()
	// p为16位质数
	//p := big.NewInt(2)
	//g := crypto.GenPrimeNumber(2)
	g := big.NewInt(2)

	p := crypto.GenPrimeNumber(512)

	fmt.Println("g:", g)
	fmt.Println("p:", p)

	// EXP(a, b, c) = (a ** b) % c
	// 服务器生成A发给客户端
	a := crypto.GenPrimeNumber(100)
	A := big.NewInt(0).Exp(g, a, p)
	fmt.Println("a:", a)
	fmt.Println("A ", A)

	// 客户端生成B发给服务器
	b := crypto.GenPrimeNumber(100)
	B := big.NewInt(0).Exp(g, b, p)
	fmt.Println("b:", b)
	fmt.Println("B ", B)

	// 如果收到的客户端B为一个任意数,则双方协商的密钥将不一致，双方也将无法解密出正确数据
	//B.SetString("123", 0)

	// 服务器拿到客户端的B，生成密钥SA5
	SA := big.NewInt(0).Exp(B, a, p)
	fmt.Println("SA", SA)

	// 客户端拿到服务器的A，生成密钥SB
	SB := big.NewInt(0).Exp(A, b, p)
	fmt.Println("SB", SB)

	fmt.Println(fmt.Sprintf("SSA: %x", md5.Sum([]byte(SA.String()))))
	fmt.Println(fmt.Sprintf("SSB: %x", md5.Sum(SB.Bytes())))

	if SA.String() != SB.String() {
		fmt.Println("SA != SB, 密钥协商失败！")
		return
	}

	// 最终SA=SB，完成密钥协商
	key := []byte(SA.String())
	// 将16位key拼接成32位key
	key = append(key, key...)

	// 先AES加密
	encrypt := crypto.AesEncryptECB([]byte("测试"), key)
	// 再异或运算加密
	for i, item := range encrypt {
		encrypt[i] = item ^ 28
	}
	fmt.Println("encrypt:", base64.StdEncoding.EncodeToString((encrypt)))

	// 先异或运算解密
	for i, item := range encrypt {
		encrypt[i] = item ^ 28
	}
	// 再AES解密
	decrypt := crypto.AesDecryptECB(encrypt, key)
	fmt.Println("decrypt:", string(decrypt))
	fmt.Println(time.Now().Sub(start))

}
