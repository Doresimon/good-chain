package p2p

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/crypto/coding"
	"github.com/Doresimon/good-chain/types"
	net "github.com/libp2p/go-libp2p-core/network"
)

func handleLog(s net.Stream) {
	console.Info(fmt.Sprintf("got a new connection, protocol = %s, peer = %s", s.Protocol(), s.Conn().RemoteMultiaddr().String()))

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go func() {
		for {
			newMsg, err := ReadOneMessage(rw)
			if err != nil {
				console.Error(err.Error())
				return
			}
			console.Info("new message received")
			fmt.Printf("msg.Type = %d\n", newMsg.Type)
			fmt.Printf("msg.Content = %s\n", newMsg.Content)

			l, err := chain.UnmarshalLog(newMsg.Content)
			if err != nil {
				console.Error(err.Error())
				return
			}
			tx := new(types.Transaction)
			err = json.Unmarshal(l.TX, tx)
			if err != nil {
				console.Error(err.Error())
				return
			}

			fmt.Printf("tx.Type = %s\n", tx.Type)
			fmt.Printf("tx.Content = %s\n", tx.Content)
			fmt.Printf("tx.KeyPath = %s\n", tx.KeyPath)
			fmt.Printf("tx.TimeStamp = %s\n", tx.TimeStamp)
			fmt.Printf("tx.Nonce = %s\n", tx.Nonce)

		}
	}()
}

func handleStream(s net.Stream) {
	console.Info("Got a new connection!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go OnReceiveData(rw)
	// go sendData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}

// OnReceiveData ...
func OnReceiveData(rw *bufio.ReadWriter) {
	for {
		newMsg, err := ReadOneMessage(rw)
		if err != nil {
			console.Warn(err.Error())
			return
		}
		console.Info("new message received")
		fmt.Printf("Type = %d\n", newMsg.Type)
		fmt.Printf("Content = %s\n", newMsg.Content)

		heartbeat := []byte("heart beart")
		msg := NewMessage(HEARTBEAT, heartbeat)
		msgBytes := Serialize(msg)

		_, err = rw.Write(msgBytes)
		if err != nil {
			console.Warn(err.Error())
		}
		rw.Flush()
		console.Info("response sent")
	}
}

// ReadOneMessage ...
func ReadOneMessage(rw *bufio.ReadWriter) (*Message, error) {
	var buf = make([]byte, 4, 4)
	var err error
	for i := 0; i < 4; i++ {
		buf[i], err = rw.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	msgLen := coding.BytesToUint32(buf)

	unreadLen := rw.Available()
	if unreadLen < int(msgLen) {
		return nil, fmt.Errorf("msg borken")
	}

	msgBytes, err := rw.Peek(int(msgLen))
	if err != nil {
		return nil, err
	}

	_, err = rw.Discard(int(msgLen))
	if err != nil {
		return nil, err
	}

	msg, err := Unserialize(append(buf, msgBytes...))
	return msg, err
}

var testOnlyPrivKeyHex = "" +
	"080012a709308204a302010002820101" + "00cd8c842d1e398f3a2af49296a29aa7" +
	"a0af3fd6e9e95d075101c0fbc076a332" + "21afd6c47a174ddda06bfe9d9fb06a29" +
	"f41d7542c1da9b7db2f160bb55a4c5d9" + "288e19808ff26241e0af76908e03cb65" +
	"2c18f970492484c572ed458c4c12a318" + "6bc9a5420b4ebd298b7aefaf2660b549" +
	"25a9a08334fdda31a2ef85631346648c" + "c952aa2c955c3f32e9605882a0550e1b" +
	"b2ba3e8f7dea24568d4d5e2722eb8c72" + "886570f96ca99ad661c5a0be23e24038" +
	"7b4344d667504b81f98ccce302626f20" + "8c90a30ea45a78d034a985511ea1214c" +
	"772a967665ded8d799669679b5aec76f" + "9e2ad70eba37a9cbc088c80f8f92c073" +
	"8756b5b12210c8814078d7c61f184e5d" + "9902030100010282010100bb081e0933" +
	"d46a282287f28e909caae4c23213ab3a" + "05c52f87aa0329688a2c868c8a9eb2fd" +
	"5b83cb8218e77233c31633a34e5c952a" + "068c44f3eff1e6070d6400cbf4169064" +
	"9bd597cdf0d3adf573f0b6446f4c25b1" + "ce022006185f85a4fbb4aff786a0e6cb" +
	"19c2a0c1173147797ddb7622b8044b86" + "a530bf32b52ba69149d2f388007cd077" +
	"f022655a8491fad3c97b12a99802ea38" + "8e42b0c3d804ba2c671c2ef2f7c0046d" +
	"60d3cf09c5c72b064c2aedc0084a6631" + "e11d34e6ed247378e21822da4fbb4eb2" +
	"733fe314390ae322dc4ca42a3e3a3f5f" + "eb5ed45b55e11d9fea7fc2a5c6b3967f" +
	"2a4442ad0c7100b8021baa6bed41a328" + "02a09d3949dcc96ee9710102818100e2" +
	"b4d8f1190f6a7a4487f444fb087fc811" + "5e3aef472f402e2d9da4ec9347fda943" +
	"2d3e909fa099537b839f8fe8d1ee15aa" + "0ad04988ca0be3e6b8cfd155294cf8d6" +
	"462e5c7aa03fc4175ee42489c879e460" + "068a6b312816e3c18c380be846f63bc6" +
	"7d2ca72884aa1c179053705c84783c3a" + "376dba78ffc620fa1458201f5ae3b902" +
	"818100e81bce14921aee4741ff0bdbda" + "4f73c6f4f3de985dd9094edf7dfba76d" +
	"b0195f38a0b17a8db46e317696d8789a" + "dd8dbd86c2211517f1f9634d8a7b161a" +
	"8d79addc83c9ef2a9b83596496b94dea" + "7acfd0bae6611322723b4ff74f04c52a" +
	"03ff75d97a4f6f94d2cdddbd14efa97e" + "d99155498b67e371ffab4bafdc78509d" +
	"d9f8e102818054dac000ae1299161ea4" + "8c876d36a72d6962cdc855ea97a93125" +
	"ba5f5a592dd6b6f02e64cf7abb290628" + "2407276bbc239bbbe46e41a6ee806511" +
	"73a7b4aa7ac700dccfe9897485a98ddd" + "d4a2e0778b3831c100fa5231e12f6a78" +
	"635a019cddf94db7f888b09fc7543075" + "b2e5719b9bad5f54f3db4202ea53d986" +
	"d7dd5e1f14f10281803d7c882618e23b" + "76f2fbae578a487be21ba7b2b2e84183" +
	"5eba83b433913212369058b94b3ed8aa" + "4c3e1f0daf3d4f2daa2455aecfd8ad08" +
	"248f38fd0c48f26b666c4119305b3821" + "1e8030115c9b9df70953498e33b46f50" +
	"5909a98c18f2dd664200df8ec57f462c" + "b4edceeb021a91661792c8f437634e98" +
	"410f9036eb1c35cee102818079cbac71" + "5bbd20409457038822b7bece5f34ee31" +
	"2ea221156d71580a09dd648f28afdaa7" + "1ee5e4388b0f390ce4cbe37a48f1ed53" +
	"d0647b67443373db66e79f4ec6068066" + "47e760d5b40d410adcb2960e703247f0" +
	"e30ca924e182f50eeb73c94377077138" + "5a72c57ea4e0dff42dd31178e800bac2" +
	"55642e122da2823effbe4e61"
