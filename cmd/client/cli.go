package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	mrand "math/rand"
	"os"
	"time"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/crypto/hdk"
	"github.com/Doresimon/good-chain/crypto/rand"
	"github.com/Doresimon/good-chain/middleware/application"
	"github.com/Doresimon/good-chain/p2p"
	"github.com/multiformats/go-multiaddr"
	"github.com/urfave/cli"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

var rw *bufio.ReadWriter

func main() {
	startP2P()
	app := App()
	err := app.Run(os.Args)

	if err != nil {
		console.Fatal("error")
	}

}

// AppCommands ...
var appCommands = []cli.Command{
	{
		Name:        "new-org",
		Aliases:     []string{"new-org"},
		Usage:       "create a new org account",
		Description: "create a new org account",
		Action: func(c *cli.Context) error { // read config
			go func() {
				logBytes := newOrg()

				msg := p2p.NewMessage(p2p.LOG, logBytes)
				data := p2p.Serialize(msg)
				_, err := rw.Write(data)
				if err != nil {
					panic(err)
				}
				err = rw.Flush()
				if err != nil {
					panic(err)
				}
			}()

			select {}
		},
	},
}

// App ...
func App() *cli.App {
	var app = cli.NewApp()

	app.Name = "node"
	app.Usage = "run a node of good chain"
	app.Commands = appCommands

	return app
}

func newOrg() []byte {
	// key setup
	key := []byte("good chain key")
	seed, _ := rand.Bytes(256)
	var err error
	masterPrivKey, masterChainCode, err := hdk.GenerateMasterKey(key, seed)
	if err != nil {
		panic(err)
	}

	content := new(application.OrgCreation)
	// content.Name = "fudan"
	// content.Extra = "复旦大学, 中国上海, 邯郸路220号"
	content.Name = "tongji"
	content.Extra = "同济大学, 中国上海, 地址不详"
	content.PublicKey = masterPrivKey.Public().HexString()
	content.ChainCode = fmt.Sprintf("%x", masterChainCode)

	body := new(chain.Body)
	body.Type = "ORG"
	body.Action = "CREATE"
	body.Timestamp = uint32(time.Now().Unix())
	body.ContentBytes, err = json.Marshal(content)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	sigBytes, err := masterPrivKey.Sign(bodyBytes)
	if err != nil {
		panic(err)
	}

	log := chain.NewLog(masterPrivKey.Public().Bytes(), sigBytes, body)

	logBytes, err := log.Marshal()
	if err != nil {
		panic(err)
	}

	return logBytes
}

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
