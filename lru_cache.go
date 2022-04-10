package main

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	length   int
	linkList *list.List
	hashList []*list.Element
}

type hashValue struct {
	key   int
	value int
	next  *list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		linkList: list.New(),
		hashList: make([]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	hashKey := key % this.capacity
	pre := this.hashList[hashKey]

	if pre == nil {
		return -1
	}

	cur := pre.Value.(*hashValue).next

	for cur != nil {

		temp := cur.Value.(*hashValue)

		if temp.key == key {
			this.linkList.MoveToBack(cur)
			return temp.value
		}

		pre, cur = cur, temp.next
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	hashKey := key % this.capacity

	if this.hashList[hashKey] == nil {
		this.hashList[hashKey] = &list.Element{
			Value: &hashValue{},
		}
	}

	pre := this.hashList[hashKey]
	cur := pre.Value.(*hashValue).next

	for cur != nil {

		temp := cur.Value.(*hashValue)

		if temp.key == key {
			temp.value = value
			this.linkList.MoveToBack(cur)
			return
		}

		pre, cur = cur, temp.next
	}

	preValue := pre.Value.(*hashValue)
	preValue.next = this.linkList.PushBack(&hashValue{
		key:   key,
		value: value,
	})

	this.length++
	if this.length > this.capacity {
		this.remove()
	}
}

func (this *LRUCache) remove() {
	front := this.linkList.Front()

	key := front.Value.(*hashValue).key
	hashKey := key % this.capacity

	pre := this.hashList[hashKey]
	cur := pre.Value.(*hashValue).next

	for cur != nil {

		temp := cur.Value.(*hashValue)

		if temp.key == key {
			pre.Value.(*hashValue).next = temp.next
			break
		}

		pre, cur = cur, temp.next
	}

	this.linkList.Remove(front)
	this.length--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
