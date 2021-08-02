package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func clear(db *leveldb.DB) error {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		if err := db.Delete(iter.Key(), nil); err != nil {
			return err
		}
	}
	iter.Release()
	if err := iter.Error(); err != nil {
		return err
	}
	return nil
}

func main() {
	// 打开或者创建一个leveldb实例
	db, err := leveldb.OpenFile("leveldb/exampledb", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 读取数据, 不存在的k-v值读取为[]byte("")
	data, err := db.Get([]byte("key"), nil) // return ([]byte, error)
	fmt.Println(string(data))

	// 更新、插入数据
	err = db.Put([]byte("key"), []byte("value"), nil) // return error
	if err != nil {
		panic(err)
	}

	// 删除数据, 允许删除不存在的k-v
	err = db.Delete([]byte("key1"), nil) // return error
	err = db.Delete([]byte("key"), nil)
	if err != nil {
		panic(err)
	}

	_ = clear(db)
	err = db.Put([]byte("1"), []byte("1"), nil)
	err = db.Put([]byte("2"), []byte("2"), nil)
	err = db.Put([]byte("3"), []byte("3"), nil)
	err = db.Put([]byte("4"), []byte("4"), nil)

	// 迭代 leveldb
	iter := db.NewIterator(nil, nil)

	// iter.Next() 迭代器游标后移 return bool
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%s -> %s\n", key, value)
	}

	ok := iter.First()                            // 迭代器游标移动到第一个  return bool
	ok = iter.Last()                              // 迭代器游标移动到最后一个
	ok = iter.Prev()                              // 迭代器游标移动到上一个
	ok = iter.Next()                              // 迭代器游标移动到下一个
	ok = iter.Seek([]byte("3"))                   // 迭代器游标移动到指定元素
	fmt.Printf("seek 3 %t, %s\n", ok, iter.Key()) // seek 3 true, 3
	ok = iter.Seek([]byte("8"))
	fmt.Printf("seek 8 %t, %s\n", ok, iter.Key()) // seek 8 false,
	iter.Prev()
	fmt.Println(string(iter.Key())) // 4

	// 修改数据对db数据生效，对iter数据不生效，iter游标释放后，iter提交更改
	_ = db.Put([]byte("1"), []byte("one"), nil)
	ok = iter.Seek([]byte("1"))
	fmt.Printf("seek 1 %t, value %s\n", ok, iter.Value()) // seek 1 true, value 1
	data, err = db.Get([]byte("1"), nil)
	fmt.Printf("Get data key 1 , value %s\n", data) // Get data key 1 , value one

	iter.Release() // 释放游标
	if err = iter.Error(); err != nil {
		panic(err)
	}

	_ = clear(db)
	err = db.Put([]byte("1-1"), []byte("1-1"), nil)
	err = db.Put([]byte("1-2"), []byte("1-2"), nil)
	err = db.Put([]byte("1-3"), []byte("1-3"), nil)
	err = db.Put([]byte("2"), []byte("2"), nil)
	err = db.Put([]byte("3"), []byte("3"), nil)

	// 根据前缀prefix迭代查询
	iter = db.NewIterator(util.BytesPrefix([]byte("1")), nil)
	for iter.Next() {
		fmt.Printf("%s -> %s\n", iter.Key(), iter.Value())
	}
	iter.Release()

	_ = clear(db)
	err = db.Put([]byte("1-1"), []byte("1-1"), nil)
	err = db.Put([]byte("1-2"), []byte("1-2"), nil)
	err = db.Put([]byte("1-3"), []byte("1-3"), nil)
	err = db.Put([]byte("1-3-1"), []byte("1-3-1"), nil)
	err = db.Put([]byte("2"), []byte("2"), nil)
	err = db.Put([]byte("3"), []byte("3"), nil)

	// 子集查询 key in [start, limit)
	iter = db.NewIterator(&util.Range{
		Start: []byte("1-"),
		Limit: []byte("2"),
	}, nil)
	for iter.Next() {
		fmt.Printf("%s -> %s\n", iter.Key(), iter.Value())
	}
	iter.Release()

	_ = clear(db)
	// 批量操作
	batch := new(leveldb.Batch)
	batch.Put([]byte("1"), []byte("1"))
	batch.Put([]byte("2"), []byte("2"))
	batch.Put([]byte("3"), []byte("3"))
	batch.Put([]byte("3"), []byte("4"))
	batch.Delete([]byte("2"))
	// 写入db
	if err = db.Write(batch, nil); err != nil {
		panic(err)
	}
	iter = db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("%s -> %s\n", iter.Key(), iter.Value())
	}

	// leveldb 快照
	snap, err := db.GetSnapshot()
	if err != nil {
		panic(err)
	}
	snap.Release() // 使用snapshot的release方法释放

	// 使用bloom filter，减少硬盘访问频率
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err = leveldb.OpenFile("path/to/db", o)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
