/**
 * memcache-objects.go
 *
 * How to store objects in memcache using bradfitz/gomemcache package
 * Memcache stores data as []byte, so you need to first encode the
 * object prior to storing.

 * I typically encode using JSON, often the data is already or will
 * soon be in JSON format, plus human readable helps debugging. If
 * space and performance is a concern, see "encoding/gob" for another
 * format.
 *
 */
package main

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// define Dog object type
type Dog struct {
	Name  string
	Color string
}

func main() {

	// connect to memcache server
	mc := memcache.New("127.0.0.1:11211")

	// try to pull from memcache
	fetchItem, err := mc.Get("dog")

	// check for cache hit
	if err != memcache.ErrCacheMiss {
		if err != nil {
			fmt.Println("Error fetching from memcache", err)
		} else {
			fmt.Println("Cache hit!")

			dog, err := DecodeData(fetchItem.Value)
			if err != nil {
				fmt.Println("Error decoding data from memcache", err)
			} else {
				fmt.Println("Dog name is: ", dog.Name)
			}
		}
	}

	// create instance of object and set properties
	spot := Dog{Name: "Spot", Color: "brown"}

	// create memcache item to store
	setItem := memcache.Item{Key: "dog", Value: EncodeData(spot), Expiration: 300}

	err = mc.Set(&setItem)
	if err != nil {
		fmt.Println("Error setting memcache item", err)
	}

	// run twice
}

func DecodeData(raw []byte) (dog Dog, err error) {
	err = json.Unmarshal(raw, &dog)
	return dog, err
}

func EncodeData(dog Dog) []byte {
	enc, err := json.Marshal(dog)
	if err != nil {
		fmt.Println("Error encoding Action to JSON", err)
	}
	return enc
}
