package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := newNode(1)
	addNewNode(&l1, 2)
	l2 := newNode(1)
	addNewNode(&l2, 4)

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

func addNewNode(l **ListNode, item int) {
	lst := *l
	if lst != nil {
		for lst.Next != nil {
			lst = lst.Next
		}
		lst.Next = newNode(item)
	} else {
		node := newNode(item)
		*l = node
	}
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	l1copy := l1
	l2copy := l2
	if l1copy == nil && l2copy == nil {
		return nil
	}
	var mergedList *ListNode

	for l1copy != nil && l2copy != nil {
		if l1copy.Val < l2copy.Val {
			addNewNode(&mergedList, l1copy.Val)
			if l1copy.Next != nil {
				l1copy = l1copy.Next
			} else {
				l1copy = nil
			}
		} else {
			addNewNode(&mergedList, l2copy.Val)
			if l2copy.Next != nil {
				l2copy = l2copy.Next
			} else {
				l2copy = nil
			}
		}
	}
	if l1copy != nil {
		for l1copy != nil {
			addNewNode(&mergedList, l1copy.Val)
			if l1copy.Next != nil {
				l1copy = l1copy.Next
			} else {
				l1copy = nil
			}
		}
	}
	if l2copy != nil {
		for l2copy != nil {
			addNewNode(&mergedList, l2copy.Val)
			if l2copy.Next != nil {
				l2copy = l2copy.Next
			} else {
				l2copy = nil
			}
		}
	}
	return mergedList
}
