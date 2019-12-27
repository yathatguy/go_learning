package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := newNode(1)
	addNewNode(l1, 2)
	addNewNode(l1, 3)
	l2 := newNode(1)
	addNewNode(l2, 3)
	addNewNode(l2, 4)

	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mergedList := merge(l1, l2)
	return mergedList
}

func newNode(item int) *ListNode {
	return &ListNode{
		Val:  item,
		Next: nil,
	}
}

func addNewNode(l *ListNode, item int) {
	for l.Next != nil {
		l = l.Next
	}
	node := newNode(item)
	l.Next = node
}

func IsEmpty(l *ListNode) bool {
	return reflect.DeepEqual(&l, ListNode{})
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	l1copy := l1
	l2copy := l2
	var mergedList ListNode

	if l1copy.Val < l2copy.Val {
		mergedList = ListNode{
			Val:  l1copy.Val,
			Next: nil,
		}
		l1copy = l1copy.Next
	} else {
		mergedList = ListNode{
			Val:  l2copy.Val,
			Next: nil,
		}
		l2copy = l2copy.Next
	}

	for !(IsEmpty(l1copy) || IsEmpty(l2copy)) {
		if l1copy.Val < l2copy.Val {
			addNewNode(&mergedList, l1copy.Val)
			if l1copy.Next != nil {
				l1copy = l1copy.Next
			} else {
				l1copy = &ListNode{}
			}
		} else {
			addNewNode(&mergedList, l2copy.Val)
			if l2copy.Next != nil {
				l2copy = l2copy.Next
			} else {
				l2copy = &ListNode{}
			}
		}
	}
	if IsEmpty(l1copy) {
		for !IsEmpty(l2copy) {
			addNewNode(&mergedList, l2copy.Val)
			if l2copy.Next != nil {
				l2copy = l2copy.Next
			} else {
				l2copy = &ListNode{}
			}
		}
	} else {
		for !IsEmpty(l1copy) {
			addNewNode(&mergedList, l1copy.Val)
			if l1copy.Next != nil {
				l1copy = l1copy.Next
			} else {
				l1copy = &ListNode{}
			}
		}
	}
	return &mergedList
}