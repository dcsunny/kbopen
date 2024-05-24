package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	// DefaultEvery means the clock time of recycling the expired cache items in memory.
	DefaultEvery = 60 // 1 minute
)

// MemoryItem store memory cache item.
type MemoryItem struct {
	val         interface{}
	createdTime time.Time
	lifespan    time.Duration
}

func (mi *MemoryItem) isExpire() bool {
	// 0 means forever
	if mi.lifespan == 0 {
		return false
	}
	return time.Now().Sub(mi.createdTime) > mi.lifespan
}

func (mi *MemoryItem) ttl() time.Duration {
	if mi.lifespan == 0 {
		return -1
	}
	diff := mi.lifespan - time.Now().Sub(mi.createdTime)
	if diff <= 0 {
		return 0
	}
	return diff
}

// MemoryCache is Memory cache adapter.
// it contains a RW locker for safe map storage.
type MemoryCache struct {
	sync.RWMutex
	dur       time.Duration
	items     map[string]*MemoryItem
	Every     int // run an expiration check Every clock time
	isStartGc bool
}
type MemoryConf struct {
	Interval int //单位s
}

// NewMemoryCache new cache
func NewMemoryCache() *MemoryCache {

	cache := &MemoryCache{items: make(map[string]*MemoryItem)}
	cache.startAndGC()
	return cache
}

// Get cache from memory.
// if non-existed or expired, return nil.
func (bc *MemoryCache) Get(name string) interface{} {
	bc.RLock()
	defer bc.RUnlock()
	if itm, ok := bc.items[name]; ok {
		if itm.isExpire() {
			return nil
		}
		return itm.val
	}
	return nil
}

// Set cache to memory.
// if lifespan is 0, it will be forever till restart.
func (bc *MemoryCache) Set(name string, value interface{}, lifespan time.Duration) error {
	bc.Lock()
	defer bc.Unlock()
	bc.items[name] = &MemoryItem{
		val:         value,
		createdTime: time.Now(),
		lifespan:    lifespan,
	}
	return nil
}

// Delete cache in memory.
func (bc *MemoryCache) Delete(name string) error {
	bc.Lock()
	defer bc.Unlock()
	if _, ok := bc.items[name]; !ok {
		return errors.New("key not exist")
	}
	bc.items[name] = nil
	delete(bc.items, name)
	if _, ok := bc.items[name]; ok {
		return errors.New("delete key error")
	}
	return nil
}

// IsExist check cache exist in memory.
func (bc *MemoryCache) IsExist(name string) bool {
	bc.RLock()
	defer bc.RUnlock()
	if v, ok := bc.items[name]; ok {
		return !v.isExpire()
	}
	return false
}

// ClearAll will delete all cache in memory.
func (bc *MemoryCache) ClearAll() error {
	bc.Lock()
	defer bc.Unlock()
	bc.items = make(map[string]*MemoryItem)
	return nil
}

// StartAndGC start memory cache. it will check expiration in every clock time.
func (bc *MemoryCache) startAndGC() error {
	dur := time.Duration(DefaultEvery) * time.Second
	bc.Every = DefaultEvery
	bc.dur = dur
	if bc.isStartGc {
		return nil
	}
	bc.isStartGc = true
	go bc.vacuum()
	return nil
}

// check expiration.
func (bc *MemoryCache) vacuum() {
	bc.RLock()
	every := bc.Every
	bc.RUnlock()

	if every < 1 {
		return
	}
	for {
		time.Sleep(time.Second)
		<-time.After(bc.dur)
		if bc.items == nil {
			return
		}
		length := len(bc.items)
		if length == 0 {
			bc.Lock()
			bc.items = nil
			bc.Unlock()
		}
		if length < 10000 {
			bc.notExpiredKeys()
		} else {
			bc.expiredKeys()
		}
	}
}

// expiredKeys returns key list which are expired.
func (bc *MemoryCache) expiredKeys() {
	bc.RLock()
	keys := make([]string, 0)
	for key, itm := range bc.items {
		if itm.isExpire() {
			keys = append(keys, key)
		}
	}
	bc.RUnlock()
	if len(keys) == 0 {
		return
	}
	bc.clearItems(keys)
	keys = nil
	return
}

func (bc *MemoryCache) notExpiredKeys() {
	bc.RLock()
	keyMap := make(map[string]*MemoryItem)
	for _key, _itm := range bc.items {
		key := _key
		itm := _itm
		if !itm.isExpire() {
			keyMap[key] = itm
		}
	}
	bc.RUnlock()
	bc.Lock()
	bc.items = nil
	bc.items = keyMap
	bc.Unlock()
	return
}

// clearItems removes all the items which key in keys.
func (bc *MemoryCache) clearItems(keys []string) {
	bc.Lock()
	defer bc.Unlock()
	for _, key := range keys {
		bc.items[key] = nil
		delete(bc.items, key)
	}
}

func (bc *MemoryCache) TTL(key string) time.Duration {
	bc.RLock()
	defer bc.RUnlock()
	if _, ok := bc.items[key]; !ok {
		return 0
	}
	return bc.items[key].ttl()
}

func (bc *MemoryCache) Size() int {
	bc.RLock()
	defer bc.RUnlock()
	return len(bc.items)
}
