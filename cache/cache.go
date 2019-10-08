package cache

import (
	"reflect"
	"time"
)

/**
Data item interface
*/
type Item interface {
	//Always return false if the item has not set expire time
	//Return true if the item has expired
	Expire() bool
}

type ItemImpl struct {
	Date     interface{}
	ExpireAt int64
}

func (item *ItemImpl) Expire() bool {
	if item.ExpireAt <= 0 {
		return false
	}
	return time.Now().UnixNano() > item.ExpireAt
}

/**
cache interface
*/
type Cache interface {
	//Set key to hold the interface value
	//Replace if the key exists
	Set(key string, value interface{})

	//Get the interface value from cache.
	//Return nil and false if the key not found.
	Get(key string) (value interface{}, b bool)

	//Same as Set + Get
	//Get the old interface value and set a new item
	//Return nil and false if the key not found before set.
	GetSet(key string, value interface{}) (oValue interface{}, b bool)

	//Multi Get.
	//Return the value map of all specified keys, ignore keys if not found
	//Return empty map if all keys not found
	MGet(key ...string) map[string]interface{}

	//Set key to hold the interface value with expire duration
	//Same as Set + Expire
	//The smallest unit of timeout is seconds
	SetEx(key string, duration time.Duration, value interface{})

	//Set key to hold string value if key does not exist
	//Return error if the key exist
	SetNx(key string, value interface{}) error

	//Returns the size of the value stored at key
	//Return 0 when key does not exist
	Len(key string) int64

	//Sets the given keys to their respective values
	//Replace existing values with new values
	MSet(map[string]interface{})

	//Sets the given keys to their respective values
	//Return error when at least one key already existed
	MSetNx(map[string]interface{}) error

	//Increments the number stored at key by one.
	//It is set to 0 before performing the operation if not exist
	//Return the value of key after the increment
	//Return error if the value is not number
	Incr(key string) (int64, e error)

	//Increments the number stored at key by increment.
	//It is set to 0 before performing the operation if not exist
	//Return the value of key after the increment
	//Return error if the value is not number
	IncrBy(key string, incr int64) error

	//Increments the number stored at key by float increment.
	//It is set to 0 before performing the operation if not exist
	//Return the value of key after the increment
	//Return error if the value is not number
	IncrByFloat(key string, incr float64) error

	//Decrements the number stored at key by one.
	//It is set to 0 before performing the operation if not exist
	//Return the value of key after the increment
	//Return error if the value is not number
	Decr(key string) error

	//Removes the specified keys. A key is ignored if it does not exist.
	//Return integer,the number of keys deleted, 0 if the key does not exist,
	Del(key ...string) int

	//Returns if key exist.
	//Return integer, the number of keys existing, 0 if no keys exist
	Exists(key ...string) int

	//Set timeout duration on the key from now.
	//Return boolean, false if the key not exist, true if set timeout successfully
	//The smallest unit of timeout is seconds
	Expire(key string, duration time.Duration) bool

	//Set timeout unix timestamp, like Expire
	//Return boolean, false if the key not exist, true if set timeout successfully
	ExpireAt(key string, timestamp int64)

	//Return a random key
	//Return nil if empty
	RandomKey() string

	//Renames key to new key.
	//Return error when the older key not exist
	Rename(key, newkey string) error

	//Like Rename
	//Renames key to newkey if newkey does not yet exist
	//Return error when either keys not exist
	RenameNx(key, newkey string) error

	//Returns the remaining time to live of a key
	//Return integer
	// -1 if the key exist but has no associated expire
	// -2 if the key does not exist
	Ttl(k string) int64

	//Return the representation of a Go type.
	Type(key string) reflect.Type
}
