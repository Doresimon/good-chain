package p2p

import (
	"context"
	"encoding/hex"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
)

// Service TODO
type Service struct {
	sourcePort      int
	sourceMultiAddr multiaddr.Multiaddr
	host            host.Host
	pid             protocol.ID
	privKey         crypto.PrivKey
}

// NewService create & run a new p2p service
func NewService() {
	S := new(Service)
	S.sourcePort = 8080

	var testOnlyPrivKeyBytes, _ = hex.DecodeString(testOnlyPrivKeyHex)
	var testOnlyPrivKey, _ = crypto.UnmarshalPrivateKey(testOnlyPrivKeyBytes)

	// r := mrand.New(mrand.NewSource(int64(S.sourcePort)))
	// // r :=rand.Reader
	// privKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	// if err != nil {
	// 	panic(err)
	// }

	S.privKey = testOnlyPrivKey
	S.sourceMultiAddr, _ = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", S.sourcePort))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	host, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(S.sourceMultiAddr),
		libp2p.Identity(S.privKey),
	)
	if err != nil {
		panic(err)
	}

	// Set a function as stream handler.
	// This function is called when a peer connects, and starts a stream with this protocol.
	// Only applies on the receiving side.
	host.SetStreamHandler("/chat", handleStream)
	host.SetStreamHandler("/log", handleLog)

	// Let's get the actual TCP port from our listen multiaddr, in case we're using 0 (default; random available port).
	var port string
	for _, la := range host.Network().ListenAddresses() {
		if p, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
			port = p
			break
		}
	}

	if port == "" {
		panic("was not able to find actual local port")
	}

	fmt.Printf("host.ID().Pretty() = %s\n", host.ID().Pretty())
	fmt.Printf("host.port = %s\n", port)
	fmt.Printf("addr = /ip4/127.0.0.1/tcp/%v/p2p/%s\n", port, host.ID().Pretty())
	fmt.Println("You can replace 127.0.0.1 with public IP as well.")
	fmt.Printf("\nWaiting for incoming connection\n\n")

	// <-make(chan int)
	select {}
}
