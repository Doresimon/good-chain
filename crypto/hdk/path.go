package hdk

import (
	"fmt"
	"strconv"
	"strings"
)

type Path struct {
	Indexes []uint32
	Root    string // eg: pk
}

func (p *Path) String() string {
	var arr = make([]string, 1+len(p.Indexes))
	arr[0] = p.Root

	for i, v := range p.Indexes {
		arr[i] = strconv.Itoa(int(v))
	}

	fmt.Printf("String().arr=%v\n", arr)
	ret := strings.Join(arr, "/")
	return ret
}

func (p *Path) ParseString(str string) error {
	arr := strings.Split(str, "/")
	p.Root = arr[0]
	psLen := len(arr) - 1
	p.Indexes = make([]uint32, psLen)
	for i := 0; i < psLen; i++ {
		t, err := strconv.Atoi(arr[i+1])
		if err != nil {
			return err
		}
		p.Indexes[i] = uint32(t)
	}
	return nil
}

func (p *Path) Marshal() []byte {
	return []byte(p.String())
}

func (p *Path) Unmarshal(strBytes []byte) error {
	return p.ParseString(string(strBytes))
}

func (p *Path) ParentPath() []uint32 {
	last := len(p.Indexes) - 1
	if last < 0 {
		last = 0
	}
	return p.Indexes[0:last]
}

func (p *Path) LastIndex() uint32 {
	last := len(p.Indexes)
	if last == 0 {
		return 0
	}
	return p.Indexes[last-1]
}

// NewPath ...
func NewPath(str string) *Path {
	p := new(Path)
	err := p.ParseString(str)
	if err != nil {
		panic(err)
	}

	return p
}
