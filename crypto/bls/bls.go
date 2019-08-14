package bls

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"reflect"

	"golang.org/x/crypto/bn256"
)

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize = 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize = 64
	// SeedSize is the size, in bytes, of private key seeds. These are the private key representations used by RFC 8032.
	SeedSize = 32
)

// PrivateKey is the type of Ed25519 private keys. It implements crypto.Signer.
type PrivateKey struct {
	v     *big.Int
	bytes []byte
}

// Public returns the PublicKey corresponding to priv.
func (priv PrivateKey) Public() *PublicKey {
	var pk = new(PublicKey)
	pk.v = new(bn256.G2).ScalarBaseMult(priv.v)
	pk.bytes = pk.v.Marshal()
	return pk
}

// Sign returns the signature of message signed by private key
func (priv PrivateKey) Sign(message []byte) (signature []byte, err error) {
	sig := Sign(priv.v, message)
	return sig.Marshal(), nil
}

// HexString returns the hex string of private key
func (priv PrivateKey) HexString() string {
	return fmt.Sprintf("%x", priv.bytes)
}

// Bytes returns the Bytes Slice of private key
func (priv PrivateKey) Bytes() []byte {
	var ret = make([]byte, len(priv.bytes))
	copy(ret, priv.bytes)
	return ret
}

// PublicKey is the type of Ed25519 public keys.
type PublicKey struct {
	v     *bn256.G2
	bytes []byte
}

// HexString returns the hex string of public key
func (pub PublicKey) HexString() string {
	return fmt.Sprintf("%x", pub.bytes)
}

// Bytes returns the Bytes Slice of public key
func (pub PublicKey) Bytes() []byte {
	var ret = make([]byte, len(pub.bytes))
	copy(ret, pub.bytes)
	return ret
}

// Verify reports whether sig is a valid signature of message by publicKey. It
// will panic if len(publicKey) is not PublicKeySize.
func (pub PublicKey) Verify(message, sigBytes []byte) bool {
	sig, _ := new(bn256.G1).Unmarshal(sigBytes)
	return Verify(pub.v, message, sig)
}

// KeyGenerate () (*big.Int, *bn256.G2)
// Key Generation. For a particular user, pick random x <-$- Zp,
// and compute v = g2^x. The user’s
// public key is v <--- G2. The user’s secret key is x <--- Zp.
// func KeyGenerate() (*big.Int, *bn256.G2, PrivateKey, PublicKey) {
func KeyGenerate() (sk *PrivateKey, pk *PublicKey, err error) {
	sk = new(PrivateKey)
	pk = new(PublicKey)

	sk.v, pk.v, err = bn256.RandomG2(rand.Reader)
	if err != nil {
		return
	}

	buf1 := sk.v.Bytes()
	sk.bytes = make([]byte, len(buf1))
	copy(sk.bytes, buf1)

	buf2 := pk.v.Marshal()
	pk.bytes = make([]byte, len(buf2))
	copy(pk.bytes, buf2)

	return
}

// Sign (sk *big.Int, msg string) *bn256.G1
// Signing. For a particular user, given the secret key x and a message M 2 f0; 1g∗, compute
// h H(M), where h 2 G1, and σ hx. The signature is σ 2 G1.
func Sign(sk *big.Int, msg []byte) *bn256.G1 {
	h := hashToG1(msg)
	sig := new(bn256.G1).ScalarMult(h, sk)
	return sig
}

// Verify (pk *bn256.G2, msg string, sig *bn256.G1) bool
// Verification. Given user’s public key v, a message M, and a signature σ, compute h=H(M)
// accept if e(σ, g2) = e(h, v) holds.
func Verify(pk *bn256.G2, msg []byte, sig *bn256.G1) bool {
	g2 := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	h := hashToG1(msg)
	lp := bn256.Pair(sig, g2)
	rp := bn256.Pair(h, pk)
	ok := reflect.DeepEqual(lp.Marshal(), rp.Marshal())
	return ok
}

// Aggregate (sigs []*bn256.G1) *bn256.G1
// Aggregation. For the aggregating subset of users U ⊆ U, assign to each user an index i, ranging
// from 1 to k = jUj. Each user ui 2 U provides a signature σi 2 G1 on a message Mi 2 f0; 1g∗
// of his choice. The messages Mi must all be distinct. Compute σ Qk i=1 σi. The aggregate
// signature is σ <- G1.
func Aggregate(sigs []*bn256.G1) (*bn256.G1, error) {
	if len(sigs) <= 1 {
		err := fmt.Errorf("sigs' length is less than 2")
		return nil, err
	}

	asig := new(bn256.G1).ScalarBaseMult(big.NewInt(0))
	for _, sig := range sigs {
		asig.Add(asig, sig)
	}

	return asig, nil
}

// AVerify verify msgs & pks & asig
// Aggregate Verification. We are given an aggregate signature σ <-- G1 for an aggregating subset
// of users U, indexed as before, and are given the original messages Mi <-- {1, 0}∗ and public
// keys vi <-- G2 for all users ui <-- U. To verify the aggregate signature σ,
// 1. ensure that the messages Mi are all distinct, and reject otherwise; and
// 2. compute hi = H(Mi) for 1 ≤ i ≤ k = |U|, and accept if e(σ, g2) = MullAll(e(hi, vi)) holds.
func AVerify(asig *bn256.G1, msgs [][]byte, pks []*bn256.G2) (ok bool, err error) {
	ok = false
	if len(msgs) != len(pks) {
		err = fmt.Errorf("messages and public keys have different quantity")
		return
	}
	g2 := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	hs := make([]*bn256.G1, len(msgs), len(msgs))
	for i, msg := range msgs {
		hs[i] = hashToG1([]byte(msg))
	}
	lp := bn256.Pair(asig, g2)
	rp := bn256.Pair(hs[0], pks[0])
	for i := 1; i < len(pks); i++ {
		rp.Add(rp, bn256.Pair(hs[i], pks[i]))
	}
	ok = reflect.DeepEqual(lp.Marshal(), rp.Marshal())
	return
}

// hashToG1 map a message to G1 point, (naive version)
func hashToG1(msg []byte) *bn256.G1 {
	hash := sha256.Sum256(msg)
	bn := new(big.Int).SetBytes(hash[:])
	return new(bn256.G1).ScalarBaseMult(bn)
}
