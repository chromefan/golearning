package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main()  {

	db, err := leveldb.OpenFile("/Users/luohuanjun/go/leveldb/data", nil)
	if err !=nil {
	}
	key := "test"
	key2 := "test2"
	value := "hello"
	db.Put([]byte(key),[]byte(value),nil)
	db.Put([]byte(key2),[]byte(value),nil)
	res ,err:= db.Get([]byte(key),nil)
	if err !=nil {
	}
	fmt.Println(string(res))
	iter := db.NewIterator(nil, nil)
	for  iter.Next() {
		// Use key/value.
		fmt.Println(string(iter.Key()),string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	defer db.Close()
}
