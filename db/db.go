package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main(){
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	db, _ := leveldb.OpenFile("./LEVELDB/chain", nil)
	defer db.Close()
	
	_ = db.Put([]byte("0"), []byte("v"), nil)
	_ = db.Put([]byte("1"), []byte("v"), nil)
	_ = db.Put([]byte("2"), []byte("vv"), nil)
	_ = db.Put([]byte("3"), []byte("vvv"), nil)
	_ = db.Put([]byte("4"), []byte("vvvv"), nil)
	
	data, _ := db.Get([]byte("4"), nil)

	var v = string(data)

	fmt.Println(v)
}