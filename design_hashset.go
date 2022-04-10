package main

const mod int = 10000

type MyHashSet struct {
	list [mod]*MyLinkedList
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) getHashNode(key int) *MyLinkedList {
	hash := key % mod
	node := this.list[hash]
	if node == nil {
		this.list[hash] = &MyLinkedList{}
		return this.list[hash]
	}

	return node
}

func (this *MyHashSet) Add(key int) {

	this.getHashNode(key).Put(key, key)
}

func (this *MyHashSet) Contains(key int) bool {

	if node := this.getHashNode(key).Get(key); node != nil {
		return true
	}

	return false
}

func (this *MyHashSet) Remove(key int) {

	this.getHashNode(key).Remove(key)
}

type MyLinkedList struct {
	Key, Val  int
	Next *MyLinkedList
}

func (this *MyLinkedList) Get(key int) *MyLinkedList {

	for cur := this.Next; cur != nil; cur = cur.Next {
		if cur.Key == key {
			return cur
		}
	}

	return nil
}

func (this *MyLinkedList) Put(key, value int) {

	if node := this.Get(key); node != nil {
		node.Val = value
		return
	}

	for cur := this; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			cur.Next = &MyLinkedList{
				Key: key,
				Val: value,
			}
			return
		}
	}
}

func (this *MyLinkedList) Remove(key int) {
	for pre, cur := this, this.Next; cur != nil; pre, cur = cur, cur.Next {
		if cur.Key == key {
			pre.Next = cur.Next
			return
		}
	}
}
