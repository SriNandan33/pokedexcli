package pokecache

type Cache struct {
	data map[string][]byte
}

func NewCache() *Cache {
	return &Cache{data: map[string][]byte{}}
}

func (c *Cache) Add(key string, val []byte) {
	c.data[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.data[key]
	return val, ok
}
