package p2p

import (
	"context"
	"crypto/rand"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	protocol "github.com/libp2p/go-libp2p-protocol"
	"github.com/multiformats/go-multiaddr"
)

type Service struct {
	sourcePort      int
	sourceMultiAddr *multiaddr.Multiaddr
	host            *host.Host
	pid             protocol.ID
	privKey
}

func (S *Service) RUN() {

}

func NewService() {
	S := new(Service)
	S.sourcePort = 8080
	S.privKey = crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	S.sourceMultiAddr, _ = multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *S.sourcePort))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	host, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)

}
