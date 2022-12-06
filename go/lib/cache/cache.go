package cache

type Cache[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Has(key K) (ok bool)
	Set(key K, val V)
	Delete(key K)
	Values() []V
	Len() int
}

type NaiveCache[K comparable, V any] struct {
	items map[K]*item[V]
}

type item[V any] struct {
	value V
}

// Cache constructor
func New[K comparable, V any]() NaiveCache[K, V] {
	return NaiveCache[K, V]{
		items: make(map[K]*item[V], 0),
	}
}

func (n NaiveCache[K, V]) Set(key K, value V) {
	n.items[key] = &item[V]{value: value}
}

func (n NaiveCache[K, V]) Get(key K) (value V, ok bool) {
	item, ok := n.items[key]
	if !ok {
		return
	}
	return item.value, ok
}

func (n NaiveCache[K, V]) GetOrSet(key K, defaultValue V) (value V, ok bool) {
	if val, ok := n.Get(key); ok {
		return val, ok
	}

	n.Set(key, defaultValue)
	return n.Get(key)
}

func (n NaiveCache[K, V]) Has(key K) bool {
	_, ok := n.items[key]
	return ok
}

func (n NaiveCache[K, V]) Delete(key K) {
	delete(n.items, key)
}

func (n NaiveCache[K, V]) Len() int {
	return len(n.items)
}

func (n NaiveCache[K, V]) Values() []V {
	result := make([]V, 0, len(n.items))
	for _, value := range n.items {
		result = append(result, value.value)
	}
	return result
}
