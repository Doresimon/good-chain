package store

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Operator struct {
	db   *leveldb.DB
	Path string
}

func (this *Operator) Open() {
	this.db, _ = leveldb.OpenFile(this.Path, nil)
}

func (this *Operator) Read(key []byte) []byte {
	data, _ := this.db.Get(key, nil)
	return data
}

func (this *Operator) Write(key []byte, value []byte) error {
	err := this.db.Put(key, value, nil)
	return err
}

func (this *Operator) Close() error {
	return this.db.Close()
}
