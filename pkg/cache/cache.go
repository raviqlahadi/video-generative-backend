package cache

import (
	"sync"
)

type Cache struct {
	data  map[string]interface{}
	mutex sync.Mutex
}

var videoCache = Cache{data: make(map[string]interface{})}

func Set(key string, value interface{}) {
	videoCache.mutex.Lock()
	defer videoCache.mutex.Unlock()
	videoCache.data[key] = value
}

func Get(key string) (interface{}, bool) {
	videoCache.mutex.Lock()
	defer videoCache.mutex.Unlock()
	value, exists := videoCache.data[key]
	return value, exists
}

func Clear() {
	videoCache.mutex.Lock()
	defer videoCache.mutex.Unlock()
	videoCache.data = make(map[string]interface{})
}
