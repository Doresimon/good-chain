package main

import (
	"bufio"
	"context"
	"fmt"
	mrand "math/rand"
	"os"

	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/p2p"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
)

var rw *bufio.ReadWriter

func startP2P() *bufio.ReadWriter {
	sourcePort := 8081
	dest := "/ip4/127.0.0.1/tcp/8080/p2p/QmdST4GrZs1RGkaaVet1d4AishPE811hjohtVZV2YWnUrx"

	r := mrand.New(mrand.NewSource(int64(sourcePort)))
	privKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		panic(err)
	}

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", sourcePort))

	host, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(privKey),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("This node's multiaddresses:")
	for _, la := range host.Addrs() {
		fmt.Printf(" - %v\n", la)
	}
	fmt.Println()

	// Turn the destination into a multiaddr.
	maddr, err := multiaddr.NewMultiaddr(dest)
	if err != nil {
		console.Fatal(err.Error())
	}

	// Extract the peer ID from the multiaddr.
	info, err := peerstore.InfoFromP2pAddr(maddr)
	if err != nil {
		console.Fatal(err.Error())
	}
	host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	s, err := host.NewStream(context.Background(), info.ID, "/log")
	if err != nil {
		panic(err)
	}

	rw = bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	return rw

}

func readData(rw *bufio.ReadWriter) {
	for {
		msg, err := p2p.ReadOneMessage(rw)
		if err != nil {
			console.Warn(err.Error())
			return
		}

		fmt.Printf("msg.Type = %d\n", msg.Type)
		fmt.Printf("msg.Content = %s\n", msg.Content)
	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		msg := p2p.NewMessage(p2p.HELLO, []byte(sendData))
		_, err = rw.Write(p2p.Serialize(msg))
		if err != nil {
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			panic(err)
		}
	}

}
