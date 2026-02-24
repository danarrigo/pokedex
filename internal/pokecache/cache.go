package pokecache

import ("time"
		"sync")

type cacheEntry struct{
 	createdAt time.Time
    val       []byte
}

type Cache struct{
	CacheMap map[string]cacheEntry 
	interval time.Duration
	mu sync.Mutex
}

func NewCache(interval time.Duration)*Cache {
	c:= &Cache{
		CacheMap:make(map[string]cacheEntry),
		interval:interval,
	}
	go c.reapLoop()
	return c 
}

func (cache *Cache)Add(key string,val []byte){
	cache.mu.Lock()	
 	defer cache.mu.Unlock()
	new_entry := cacheEntry{
		createdAt: time.Now(),
		val : val,
	}
	cache.CacheMap[key]=new_entry
}

func (cache *Cache)Get(key string)([]byte,bool){
	cache.mu.Lock()	
	defer cache.mu.Unlock()
	val,ok:=cache.CacheMap[key]
	if ok {
		return val.val,true
	}
	return nil,false
}

func (cache *Cache)reapLoop (){
	ticker := time.NewTicker(cache.interval)
	for range ticker.C{
		cache.mu.Lock()
		for key,value:=range cache.CacheMap{
			if time.Since(value.createdAt)>cache.interval{
				delete (cache.CacheMap,key)
			}
		}
		cache.mu.Unlock()	
	}
	
}
