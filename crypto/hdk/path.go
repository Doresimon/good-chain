package hdk

import (
	"strconv"
	"strings"

	"github.com/Doresimon/good-chain/console"
)

type path struct {
	ps   []uint32
	root string // eg: pk
}

func (p *path) Marshal() []byte {
	var arr = make([]string, 1+len(p.ps))
	arr[0] = p.root

	for i, v := range p.ps {
		arr[i] = strconv.Itoa(int(v))
	}

	ret := strings.Join(arr, "/")
	return []byte(ret)
}

func (p *path) Unmarshal(str []byte) error {
	arr := strings.Split(string(str), "/")

	psLen := len(arr) - 1
	p.ps = make([]uint32, psLen)

	for i := 0; i < psLen; i++ {
		t, err := strconv.Atoi(arr[i+1])
		if err != nil {
			return err
		}
		p.ps[i] = uint32(t)
	}
	return nil
}

func (p *path) SetPath(path string) {
	console.Warn("hdk.SetPath is not implemented")
}
