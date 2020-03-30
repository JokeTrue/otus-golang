package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Next  *listItem
	Prev  *listItem
	Value interface{}
}

type list struct {
	firstItem *listItem
	lastItem  *listItem
	length    int
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *listItem {
	return l.firstItem
}

func (l list) Back() *listItem {
	return l.lastItem
}

func (l *list) insertAfter(item *listItem, newItem *listItem) {
	newItem.Prev = item

	if item.Next == nil {
		newItem.Next = nil
		l.lastItem = newItem
	} else {
		newItem.Next = item.Next
		item.Next.Prev = newItem
	}

	item.Next = newItem
}

func (l *list) insertBefore(item *listItem, newItem *listItem) {
	newItem.Next = item

	if item.Prev == nil {
		newItem.Prev = nil
		l.firstItem = newItem
	} else {
		newItem.Prev = item.Prev
		item.Prev.Next = newItem
	}

	item.Prev = newItem
}

func (l *list) PushFront(v interface{}) *listItem {
	newItem := &listItem{Value: v}

	if l.firstItem == nil {
		l.firstItem = newItem
		l.lastItem = newItem
	} else {
		l.insertBefore(l.firstItem, newItem)
	}

	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *listItem {
	var newItem *listItem

	if l.lastItem == nil {
		newItem = l.PushFront(v)
	} else {
		newItem = &listItem{Value: v}
		l.insertAfter(l.lastItem, newItem)
	}

	l.length++
	return newItem
}

func (l *list) Remove(i *listItem) {
	if i.Prev == nil {
		l.firstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.lastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	if l.firstItem == i {
		return
	}

	if l.lastItem == i {
		l.lastItem = i.Prev
	}

	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	tmpFirstItem := l.firstItem
	l.firstItem = i

	i.Prev = nil
	i.Next = tmpFirstItem

	tmpFirstItem.Prev = i
}

func NewList() List {
	return &list{}
}
