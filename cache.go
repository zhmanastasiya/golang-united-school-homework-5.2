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

	if time.Now().After(k.Deadline) {
		delete(c.Items, key)
		return "", false
	}
	return k.Value, true
}

func (c *Cache) Put(key, value string) {
	k := Item{Value: value, Deadline: time.Now().Add(1 * time.Minute)}
	c.Items[key] = k
}
func (c *Cache) Keys() []string {
	keys := []string{}
	now := time.Now()
	for k, p := range c.Items {
		if now.Before(p.Deadline) {
			keys = append(keys, k)
		} else {
			delete(c.Items, k)
		}

	}
	return keys
}
func (c *Cache) PutTill(key, value string, deadline time.Time) {
	k := Item{Value: value, Deadline: deadline}
	c.Items[key] = k
}
