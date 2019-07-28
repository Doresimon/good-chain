package store

import (
	"github.com/syndtr/goleveldb/leveldb"
)

// Operator .
type Operator struct {
	db   *leveldb.DB
	Path string
}

// Open .
func (op *Operator) Open() {
	op.db, _ = leveldb.OpenFile(op.Path, nil)
}

// Read .
func (op *Operator) Read(key []byte) []byte {
	data, _ := op.db.Get(key, nil)
	return data
}

// Write .
func (op *Operator) Write(key []byte, value []byte) error {
	err := op.db.Put(key, value, nil)
	return err
}

// Close .
func (op *Operator) Close() error {
	return op.db.Close()
}
