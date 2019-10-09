package data

import (
	"time"
)

/**
Data item interface
*/
type Item interface {
	//Always return false if the item has not set expire time
	//Return true if the item has expired
	Expire() bool
	//Get item expire time
	GetExpireAt() int64
	//Get value from item
	GetValue() interface{}
}

//simple data item impl
type ItemImpl struct {
	value    interface{}
	expireAt int64
}

func (item *ItemImpl) GetValue() interface{} {
	return item.value
}

func (item *ItemImpl) GetExpireAt() int64 {
	return item.expireAt
}

func (item *ItemImpl) Expire() bool {
	if item.expireAt <= 0 {
		return false
	}
	return time.Now().UnixNano() > item.expireAt
}

/**
Data block interface
*/

type DataBlock interface {
	//data block is full ?
	IsFull() bool

	//flush data block
	Flush()

	//find items by keys
	Get(keys ...string) map[string]Item

	//delete items by keys
	Del(keys ...string)

	//set item to data block
	Set(kv map[string]interface{}, duration time.Duration)

	//delete expire item
	DeleteExpire()

	//init data block
	Init(size int)

	//remove items not expire by lru or other logic
	Eliminate()
}

type DataBlockImpl struct {
	maxSize        int
	latestExpireAt int64
	dataMap        map[string]Item
}

func (data *DataBlockImpl) IsFull() bool {
	return len(data.dataMap) == data.maxSize
}

func (data *DataBlockImpl) Flush() {
	data.dataMap = make(map[string]Item, data.maxSize)
}

func (data *DataBlockImpl) Get(keys ...string) map[string]Item {
	r := make(map[string]Item)
	for _, key := range keys {
		if v, ok := data.dataMap[key]; ok {
			r[key] = v
		}
	}
	return r
}

func (data *DataBlockImpl) Del(keys ...string) {
	for _, key := range keys {
		delete(data.dataMap, key)
	}
}

func (data *DataBlockImpl) Set(kv map[string]interface{}, duration time.Duration) {
	var expireAt int64 = 0
	if duration != 0 {
		expireAt = time.Now().Add(duration).UnixNano()
	}

	for k, v := range kv {
		data.dataMap[k] = &ItemImpl{
			v,
			expireAt,
		}
	}
	//remember the latest expire time
	if expireAt < data.latestExpireAt {
		data.latestExpireAt = expireAt
	}
}

func (data *DataBlockImpl) DeleteExpire() {
	if time.Now().UnixNano() < data.latestExpireAt {
		return
	}
	for k, item := range data.dataMap {
		if item.Expire() {
			delete(data.dataMap, k)
		}
		//remember the latest expire time
		if item.GetExpireAt() < data.latestExpireAt {
			data.latestExpireAt = item.GetExpireAt()
		}
	}
}

func (data *DataBlockImpl) Init(size int) {
	if len(data.dataMap) > 0 {
		return
	}
	data.maxSize = size
	data.Flush()
}

func (data *DataBlockImpl) Eliminate() {
	data.DeleteExpire()
}
