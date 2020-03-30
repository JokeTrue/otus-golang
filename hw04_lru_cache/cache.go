package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]cacheItem
}

type cacheItem struct {
	Key   *listItem
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	cItem, found := c.items[key]
	var cacheKey *listItem

	if found {
		cacheKey = cItem.Key
		c.queue.MoveToFront(cacheKey)
	} else {
		cacheKey = c.queue.PushFront(key)
	}

	if c.queue.Len() > c.capacity {
		lastQItem := c.queue.Back()

		delete(c.items, lastQItem.Value.(Key))
		c.queue.Remove(lastQItem)
	}

	c.items[key] = cacheItem{Value: value, Key: cacheKey}
	return found
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	cacheItem, found := c.items[key]
	if found {
		c.queue.MoveToFront(cacheItem.Key)
	}

	return cacheItem.Value, found
}

func (c *lruCache) Clear() {
	c.queue = NewList()

	for key := range c.items {
		delete(c.items, key)
	}
}

func NewCache(capacity int) Cache {
	items := make(map[Key]cacheItem)
	queue := NewList()
	return &lruCache{capacity: capacity, queue: queue, items: items}
}
