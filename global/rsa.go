package global

//import "github.com/bensema/library/crypto"

var (
	PrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC7bA8oDFASBP+6JWc30k+cHcA0DQhUcuAKQCxxlC0pZgziwfcq
O7hYQJCrdH9IlSbB6Bftn8Y9Ji3J+QqWB82QOlykmaC2F5mEAs4NgupYHrA1cCvt
nDIF9M2UvXThwCwy1ovNJIlhvdaPyUQ1YdHdgk+4XXGkIl6AOTn93wYL+wIDAQAB
AoGAfWeZYJeUf+oeXiQmw2ASaogxkeJif3b514IG+txFt5yT5KoaQoUBHPPemQpz
HFLIQxiT4ih1EXdZTfo4DhkcfkyRyE5kKy9ZwLxEoY6NsIHrqy5p2qmzgSA6bUwL
XaxGyVWozrPaVbDFdaLWXJl5YUq6+Te6tJTo5SswM4bDV3kCQQD2frG50lXvOyl9
y63/KnMvvIVAJ/2BG5LnETkPHDrmNZ2RE5UqkkqeABHn8eKlNKA/4b2zRNPtwO9Y
+rvoK/hdAkEAwqY4ajciYKEXsUJAw9xLqQec0RcQBkTBzyZTV3lFIc1nikNgk1Vz
qT6t48YF4w6nj+XNrGNZgTTllwl66edwNwJAHQo2FAdUuneE3t4lJJ+yrFgQdst/
UTuXZgOgbkhMJB0C8DmXZEmR4uVtNp9HrWDy1DPbLoiYUzVcvWXkx2iOvQJBAIBN
zrv8sbVeGl7mXAh+qoS1luGgQRjQs6vXCHKNZktcuNZDiWI0nnO99CNCwVikrUDF
6qeqKTJo6rl/Lz0FKEsCQQC9K7x5bVSBlPf2PJK23ySZfqdiISjnckZMcpqrc6AO
zGtB7hhM3y09I4GxLIikqSD6PPSk1lh/+Ssocyj7vB1R
-----END RSA PRIVATE KEY-----`)
	PublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC7bA8oDFASBP+6JWc30k+cHcA0
DQhUcuAKQCxxlC0pZgziwfcqO7hYQJCrdH9IlSbB6Bftn8Y9Ji3J+QqWB82QOlyk
maC2F5mEAs4NgupYHrA1cCvtnDIF9M2UvXThwCwy1ovNJIlhvdaPyUQ1YdHdgk+4
XXGkIl6AOTn93wYL+wIDAQAB
-----END PUBLIC KEY-----
`)
)

func init() {
	//PrivateKey, PublicKey, _ = crypto.GenRsaKey()
	//fmt.Println(string(PrivateKey))
	//fmt.Println(string(PublicKey))
}
