package p2p

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
)

// Service TODO
type Service struct {
	port      int
	multiAddr multiaddr.Multiaddr
	host      host.Host
	pid       protocol.ID
	privKey   crypto.PrivKey
	cs        *chain.Service
}

// NewService create & run a new p2p service
func NewService(cs *chain.Service) *Service {
	service := new(Service)
	service.port = 8080

	var testOnlyPrivKeyBytes, _ = hex.DecodeString(testOnlyPrivKeyHex)
	var testOnlyPrivKey, _ = crypto.UnmarshalPrivateKey(testOnlyPrivKeyBytes)

	// r := mrand.New(mrand.NewSource(int64(S.port)))
	// // r :=rand.Reader
	// privKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	// if err != nil {
	// 	panic(err)
	// }

	service.privKey = testOnlyPrivKey

	multiAddr := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", service.port)
	service.multiAddr, _ = multiaddr.NewMultiaddr(multiAddr)

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	host, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(service.multiAddr),
		libp2p.Identity(service.privKey),
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
		panic("not able to find actual local port")
	}

	service.host = host

	console.Title("Node Info")
	console.Infof("host.ID().Pretty() = %s", host.ID().Pretty())
	console.Infof("host.port = %s", port)
	console.Infof("addr = /ip4/127.0.0.1/tcp/%v/p2p/%s", port, host.ID().Pretty())
	console.Infof("You can replace 127.0.0.1 with public IP as well.")
	console.Infof("Waiting for incoming connection\n")

	return service

	// monitor channel LogTransferPool
	// go func() {
	// 	for {
	// 		var l = new(chain.Log)
	// 		l = <-chain.LogTransferPool
	// 		cs.AddLog(l)
	// 	}
	// }()

	// <-make(chan int)
	// select {}
}
