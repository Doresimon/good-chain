package hdk

import (
	"fmt"
	"testing"

	"github.com/Doresimon/good-chain/crypto/rand"
	"golang.org/x/crypto/bn256"
)

func TestGenerateMasterKey(t *testing.T) {
	key := []byte("good chain seed")
	seed, _ := rand.Bytes(256)
	_, _, err := GenerateMasterKey(key, seed)
	if err != nil {
		t.Error(err)
	}
}

func TestPriv2Priv(t *testing.T) {
	key := []byte("good chain seed")
	seed, _ := rand.Bytes(256)
	masterPrivKey, masterChainCode, err := GenerateMasterKey(key, seed)
	if err != nil {
		t.Error(err)
	}

	_, _, _ = Priv2Priv(masterPrivKey, masterChainCode, 0)

}
func TestPub2Pub(t *testing.T) {
	key := []byte("good chain seed")
	seed, _ := rand.Bytes(256)

	masterPrivKey, masterChainCode, err := GenerateMasterKey(key, seed)
	if err != nil {
		t.Error(err)
	}

	childPrivKey, childChainCode, valid := Priv2Priv(masterPrivKey, masterChainCode, 0)
	if !valid {
		fmt.Printf("childPrivKey = %x\n", childPrivKey)
		fmt.Printf("childChainCode = %x\n", childChainCode)
		t.Errorf("valid == false")
		t.Fail()
		return
	}

	childPubKey := new(bn256.G2).ScalarBaseMult(childPrivKey)
	masterPubKey := new(bn256.G2).ScalarBaseMult(masterPrivKey)

	childPubKeyX, childChainCodeX, valid := Pub2Pub(masterPubKey, masterChainCode, 0)
	if !valid {
		t.Errorf("valid == false")
		t.Fail()
		return
	}

	fmt.Printf("masterPrivKey = %x\n", masterPrivKey)
	fmt.Printf("masterPubKey = %x\n", masterPubKey)
	fmt.Printf("masterChainCode = %x\n", masterChainCode)
	fmt.Printf("childPrivKey = %x\n", childPrivKey)
	fmt.Printf("childPubKeyX = %x\n", childPubKeyX)
	fmt.Printf("childChainCode = %x\n", childChainCode)

	// fmt.Printf("childPubKey  = %s\n", childPubKey)
	// fmt.Printf("childPubKeyMarshal  = %s\n", childPubKey.Marshal())
	// fmt.Printf("childPubKeyMarshalHex  = %x\n", childPubKey.Marshal())

	// data := []byte("any + old & data")
	// childPubKeyBase64 := base64.StdEncoding.EncodeToString(childPubKey.Marshal())
	// fmt.Printf("childPubKeyBase64  = %s\n", childPubKeyBase64)

	if string(childPubKey.Marshal()) != string(childPubKeyX.Marshal()) {
		fmt.Printf("childPubKey  = %s\n", childPubKey)
		fmt.Printf("childPubKeyX = %s\n", childPubKeyX)
		t.Errorf("childPubKey is not equal to childPubKeyX")
	}
	if fmt.Sprintf("%x", childChainCode) != fmt.Sprintf("%x", childChainCodeX) {
		fmt.Printf("childChainCode  = %x\n", childChainCode)
		fmt.Printf("childChainCodeX = %x\n", childChainCodeX)
		t.Errorf("childChainCode is not equal to childChainCodeX")
	}

}
