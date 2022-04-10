package main

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	for cur, count := this.Next, 0; cur != nil; cur, count = cur.Next, count+1 {
		if count == index {
			return cur.Val
		}
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Next == nil {
		this.Next = &MyLinkedList{
			Val: val,
		}
		return
	}

	temp := this.Next
	this.Next = &MyLinkedList{
		Val:  val,
		Next: temp,
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Next == nil {
		this.Next = &MyLinkedList{
			Val: val,
		}
		return
	}
	for cur := this.Next; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			cur.Next = &MyLinkedList{
				Val: val,
			}
			return
		}
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	count := 0
	pre := this
	for cur := this.Next; cur != nil; pre, cur, count = cur, cur.Next, count+1 {
		if count == index {
			pre.Next = &MyLinkedList{
				Val:  val,
				Next: cur,
			}
			return
		}
	}

	if count == index {
		pre.Next = &MyLinkedList{
			Val: val,
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	pre := this
	for cur, count := this.Next, 0; cur != nil; pre, cur, count = cur, cur.Next, count+1 {

		if count == index {
			pre.Next = cur.Next
			break
		}
	}
}
