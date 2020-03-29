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
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	_, found := c.items[key]
	if found {
		lItem := c.queue.Search(key)
		c.queue.MoveToFront(lItem)
	} else {
		c.queue.PushFront(key)
	}

	if c.queue.Len() > c.capacity {
		lastKey := c.queue.Back()

		c.queue.Remove(lastKey)
		delete(c.items, lastKey.Value.(Key))
	}

	c.items[key] = cacheItem{Value: value}
	return found
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	value, found := c.items[key]
	if found {
		lItem := c.queue.Search(key)
		c.queue.MoveToFront(lItem)
	}

	return value.Value, found
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
