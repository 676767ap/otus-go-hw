package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value      interface{}
	Next, Prev *ListItem
}

type list struct {
	head, tail *ListItem
	len        int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.head
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		l.head.Prev = newItem
		newItem.Next = l.head
		l.head = newItem
	}
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	switch {
	case l.head == nil:
		l.head = newItem
		l.tail = newItem
	case l.tail == l.head:
		newItem.Prev = l.head
		l.head.Next = newItem
		l.tail = newItem
	default:
		l.tail.Next = newItem
		newItem.Prev = l.tail
		l.tail = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.head {
		if i != l.tail {
			i.Next.Prev = i.Prev
			i.Prev.Next = i.Next
		} else {
			i.Prev.Next = nil
			l.tail = i.Prev
		}
		l.head.Prev = i
		i.Next = l.head
		i.Prev = nil
		l.head = i
	}
}

func NewList() List {
	return new(list)
}
