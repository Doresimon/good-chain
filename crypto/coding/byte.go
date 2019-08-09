package coding

import "fmt"

// Uint32ToBytes turn i to Big Endian 4 bytes
func Uint32ToBytes(i uint32) []byte {
	ret := make([]byte, 4, 4)
	ret[0] = byte((i) >> 24)
	ret[1] = byte((i) >> 16)
	ret[2] = byte((i) >> 8)
	ret[3] = byte((i) >> 0)
	return ret
}

// BytesToUint32 turn 4 bytes to uint32, Big Endian
func BytesToUint32(b []byte) uint32 {
	if len(b) != 4 {
		panic(fmt.Errorf("byte length is not 4"))
	}

	var ret = uint32(0)
	ret += uint32(b[0]) << 24
	ret += uint32(b[1]) << 16
	ret += uint32(b[2]) << 8
	ret += uint32(b[3]) << 0
	return ret
}
