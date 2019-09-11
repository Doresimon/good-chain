package bls

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func Test_BLS_Short_Signature_Scheme(t *testing.T) {
	// key generation
	sk, pk, _ := KeyGenerate()
	// signing
	msg := []byte("hello world")
	sig, _ := sk.Sign(msg)
	// verifying
	ok := pk.Verify(msg, sig)

	fmt.Printf("sk = %x\n", sk.Bytes())
	fmt.Printf("pk = %x\n", pk.Bytes())
	fmt.Printf("sig = %x\n", sig)
	fmt.Printf("hash = %x\n", sha256.Sum256(msg))

	if !ok {
		t.Error("verification failed.")
	}
}

// func Test_BLS_Aggregate_Signature_Scheme(t *testing.T) {
// 	N := 64
// 	sks := make([]*big.Int, N, N)
// 	pks := make([]*bn256.G2, N, N)
// 	msgs := make([][]byte, N, N)
// 	sigs := make([]*bn256.G1, N, N)

// 	for i := 0; i < N; i++ {
// 		sks[i], pks[i] = KeyGenerate()
// 		msgs[i] = []byte("hello world" + strconv.Itoa(i))
// 		sigs[i] = Sign(sks[i], msgs[i])
// 	}

// 	asig, err := Aggregate(sigs)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ok, err := AVerify(asig, msgs, pks)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if !ok {
// 		t.Error("aggregate signature verification failed.")
// 	}
// }

// func Benchmark_KeyGenerate(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		KeyGenerate()
// 	}
// }
// func Benchmark_Sign(b *testing.B) {
// 	sk, _ := KeyGenerate()
// 	msg := []byte("hello world")
// 	for i := 0; i < b.N; i++ {
// 		Sign(sk, msg)
// 	}
// }
// func Benchmark_Verify(b *testing.B) {
// 	sk, pk := KeyGenerate()
// 	msg := []byte("hello world")
// 	sig := Sign(sk, msg)
// 	for i := 0; i < b.N; i++ {
// 		Verify(pk, msg, sig)
// 	}
// }

// func Benchmark_Short_Scheme(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		sk, pk := KeyGenerate()
// 		KeyGenerate()
// 		msg := []byte("hello world")
// 		sig := Sign(sk, msg)
// 		ok := Verify(pk, msg, sig)
// 		if !ok {
// 			b.Error("verification failed.")
// 		}
// 	}
// }
