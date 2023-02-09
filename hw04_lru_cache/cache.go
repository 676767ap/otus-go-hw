package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type Item struct {
	Key   Key
	Value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	item := &Item{
		Key:   key,
		Value: value,
	}
	var flag bool
	if _, found := l.items[key]; found {
		l.items[key].Value.(*Item).Value = value
		l.queue.MoveToFront(l.items[key])
		flag = true
	} else {
		l.queue.PushFront(item)
		l.items[key] = l.queue.Front()
		flag = false
	}
	if l.queue.Len() == l.capacity {
		l.queue.Remove(l.queue.Back())
		delete(l.items, item.Key)
	}
	return flag
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if _, found := l.items[key]; !found {
		return nil, false
	}
	l.queue.PushFront(l.items[key])
	return l.items[key].Value.(*Item).Value, true
}

func (l *lruCache) Clear() {
	l.capacity = 0
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, 0)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
