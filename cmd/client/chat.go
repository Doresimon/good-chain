package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	mrand "math/rand"
	"os"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/p2p"
	"github.com/Doresimon/good-chain/types"
	"github.com/multiformats/go-multiaddr"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

func main() {
	// chat()
	log()
}

func log() {
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

	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go func() {
		tx := new(types.Transaction)
		tx.Type = "create"
		tx.Content = "the count infomation"
		tx.KeyPath = "x/1/2/3"
		tx.TimeStamp = "123123123123"
		tx.Nonce = "1"

		sender := []byte("sender")
		bn := []byte("123")
		sig := []byte("sig")
		txBytes, err := json.Marshal(tx)
		if err != nil {
			panic(err)
		}
		lg := chain.NewLog(sender, bn, txBytes, sig)
		lgBytes, err := lg.Marshal()
		if err != nil {
			panic(err)
		}
		msg := p2p.NewMessage(p2p.HELLO, lgBytes)
		data := p2p.Serialize(msg)
		_, err = rw.Write(data)
		if err != nil {
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			panic(err)
		}
	}()

	select {}
}

func chat() {
	sourcePort := 8081
	dest := "/ip4/127.0.0.1/tcp/8080/p2p/QmdST4GrZs1RGkaaVet1d4AishPE811hjohtVZV2YWnUrx"

	// r := rand.Reader
	r := mrand.New(mrand.NewSource(int64(sourcePort)))
	// Creates a new RSA key pair for this host.
	privKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		panic(err)
	}

	// 0.0.0.0 will listen on any interface device.
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", sourcePort))

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
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

	// Add the destination's peer multiaddress in the peerstore.
	// This will be used during connection and stream creation by libp2p.
	host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := host.NewStream(context.Background(), info.ID, "/chat")
	if err != nil {
		panic(err)
	}

	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Create a thread to read and write data.
	go writeData(rw)
	go readData(rw)

	// Hang forever.
	select {}
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
