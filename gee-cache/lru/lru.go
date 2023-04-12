package lru

import "container/list"

// Value use len to count how many bytes it takes
type Value interface {
	Len() int
}

// Cache is LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes  int64 //max memory
	bytes     int64 // use memory
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// New is the constructor of cache
func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (v Value, ok bool) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}

// RemoveOldElement removes the old element item
func (c *Cache) RemoveOldElement() {
	element := c.ll.Back()
	if element != nil {
		c.ll.Remove(element) //remove element
		kv := element.Value.(*entry)
		delete(c.cache, kv.key)
		c.bytes -= int64(len(kv.key)) + int64(kv.value.Len()) //calc use memory
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value) //delete after run func
		}
	}
}

// Add value to the cache
func (c *Cache) Add(key string, value Value) {
	//check exists
	if element, ok := c.cache[key]; ok { //edit
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		c.bytes += int64(len(kv.key)) - int64(kv.value.Len()) //calc memory
		kv.value = value
	} else {
		//new entry
		element = c.ll.PushFront(&entry{key, value})
		c.cache[key] = element
		c.bytes += int64(len(key)) + int64(value.Len()) //calc memory
	}
	//The cache is now more than upper limit Invoked automatically clean up
	for c.maxBytes != 0 && c.maxBytes < c.bytes {
		c.RemoveOldElement()
	}
}

// Len get cache entry number
func (c *Cache) Len() int {
	return c.ll.Len()
}
