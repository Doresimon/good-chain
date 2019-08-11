package rand

import (
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	ret, err := Bytes(256)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("ret = %x\n", ret)
}
