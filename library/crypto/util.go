package crypto

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	lowercaseLen = len(lowercase)

	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	uppercaseLen = len(uppercase)

	digits    = "0123456789"
	digitsLen = len(digits)
	digit1_9  = "123456789"
)

func RandString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandDigitString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		idx, _ := crand.Int(crand.Reader, big.NewInt(10))
		b[i] = digits[int(idx.Int64())]
	}
	return string(b)
}

func RandBigInt(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	switch n {
	case 0:
		return big.NewInt(0)
	case 1:
		b := make([]byte, n)
		for i := 0; i < n; i++ {
			idx, _ := crand.Int(crand.Reader, big.NewInt(9))
			b[i] = digit1_9[int(idx.Int64())]
		}
		p := new(big.Int)
		fmt.Sscan(string(b), p)
		return p
	default:
		b := make([]byte, n)
		for i := 0; i < n; i++ {
			idx, _ := crand.Int(crand.Reader, big.NewInt(10))
			b[i] = digits[int(idx.Int64())]
		}

		idx, _ := crand.Int(crand.Reader, big.NewInt(9))
		f := digit1_9[int(idx.Int64())]
		b[0] = f
		p := new(big.Int)
		fmt.Sscan(string(b), p)
		return p

	}
}

func GenPrimeNumber(n int) *big.Int {
	if n <= 0 {
		panic("n must > 0")
	}
	for {
		p := RandBigInt(n)
		if p.ProbablyPrime(20) {
			return p
		}
	}
}
