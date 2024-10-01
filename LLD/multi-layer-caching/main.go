/*
Desgin a multi layer caching system
- no of levels can be given as input
- key to be present in lower levels of cache
- statistics
- promote key if present in lower layers


WIP
write mlc methods and implement 2nd layer
add promotion of key
check if it can optimised with go-routines.

*/

package main

import "time"

type Cache interface {
	Set(level int, key string, Value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string) bool
}

type CacheStats struct {
	Hits    int
	Miss    int
	Evicted int
}

type MultiLevelCache struct {
	Cache []Cache
	Stats CacheStats
}

func (mlc *MultiLevelCache) Get(key string) {

}

type InmemoryCache struct {
	Data        map[string]interface{}
	Stats       CacheStats
	accessCount map[string]int
}

func (inmc *InmemoryCache) Set(level int, key string, val interface{}, ttl time.Duration) {
	return
}
func (inmc *InmemoryCache) Get(key string) (interface{}, bool) {

	return nil, false
}

func (inmc *InmemoryCache) Delete(key string) bool {
	return false
}

func newInmemoryCache() *InmemoryCache {
	return &InmemoryCache{
		Data:  map[string]interface{}{},
		Stats: CacheStats{},
	}
}

func newMultiLevelCache(cache []Cache) *MultiLevelCache {
	return &MultiLevelCache{
		Cache: cache,
		Stats: CacheStats{},
	}
}
func main() {

	inmc := newInmemoryCache()
	mlc := newMultiLevelCache([]Cache{
		inmc,
	})

	// WIP
	mlc.Get("key1")

}
