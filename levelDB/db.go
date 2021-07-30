package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
)

func main()  {

	db, err := leveldb.OpenFile("/Users/luohuanjun/go/leveldb/data", nil)
	if err !=nil {
	}

	key := "my-image"
	//key2 := "test2"
	//value := "hello"

	f, err := ioutil.ReadFile("/Users/luohuanjun/test.jpg")
	if err != nil {
		fmt.Println("read fail", err)
	}
	fgzip := GZipBytes(f)
	fmt.Println(len(fgzip))
	db.Put([]byte(key),fgzip,nil)
	//db.Put([]byte(key2),[]byte(value),nil)
	res ,err:= db.Get([]byte(key),nil)
	if err !=nil {
	}
	fuzip := UGZipBytes(res)
	fmt.Println(len(fuzip))
	err = ioutil.WriteFile("/Users/luohuanjun/test-gzip.jpg", fuzip, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
/*	iter := db.NewIterator(nil, nil)
	for  iter.Next() {
		// Use key/value.
		fmt.Println(string(iter.Key()),string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()*/
	defer db.Close()
}
func Write1()  {
	fileName := "file/test2"
	strTest := "测试测试"
	var d = []byte(strTest)
	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}

//压缩
func GZipBytes(data []byte) []byte{
	var in bytes.Buffer
	in.Write(data)
	g := gzip.NewWriter(&in)
	g.Write(data)
	g.Flush()
	g.Close()
	return in.Bytes()
}
//解压
func UGZipBytes(data []byte) []byte{
	var in bytes.Buffer
	in.Write(data)
	fmt.Printf("len data--- %v",len(data))
	r,err := gzip.NewReader(&in)
	fmt.Println(err)
	return nil
	defer r.Close()
	undata, _ := ioutil.ReadAll(r)
	fmt.Println("ungzip size:", len(undata))
	return undata
}