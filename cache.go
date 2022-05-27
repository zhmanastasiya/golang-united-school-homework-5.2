package cache

import "time"

type Cache struct {
	Items map[string]Item
}
type Item struct {
	Value    string
	Deadline time.Time
}

func NewCache() Cache {
	i := make(map[string]Item)
	return Cache{Items: i}
}

func (c *Cache) Get(key string) (string, bool) {
	k, ok := c.Items[key]

	if !ok {
		return "", false
	}

	if !k.Deadline.Before(time.Now()) {
		delete(c.Items, key)
		return "", false
	}

	return k.Value, ok
}

func (c *Cache) Put(key, value string) {
	c.Items[key] = Item{
		Value:    value,
		Deadline: time.Now().Add(10 * time.Minute)}
}

func (c *Cache) Keys() []string {
	keys := []string{}
	now := time.Now()
	for k, p := range c.Items {
		if p.Deadline.Before(now) {
			keys = append(keys, k)
		} else {
			delete(c.Items, k)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Items[key] = Item{
		Value:    value,
		Deadline: deadline,
	}

}
