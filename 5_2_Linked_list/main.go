package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(sl []int) *ListNode {
	if len(sl) == 0 {
		return nil
	}

	res := ListNode{}
	cur := &res
	for _, val := range sl {
		if cur.Next != nil {
			cur = cur.Next
		}
		cur.Val = val
		cur.Next = &ListNode{}
	}
	if cur.Next.Val == 0 {
		cur.Next = nil
	}
	return &res
}

func deleteDuplicates(l *ListNode) {
	if l != nil {

		m := make(map[int]bool)

		left, right := l, l

		for l.Next != nil {
			if _, ok := m[l.Val]; ok {
				left.Next = right
				l = l.Next
				right = l.Next
			} else {
				m[l.Val] = true
				left = l
				l = l.Next
				right = l.Next
			}
			if _, ok := m[l.Val]; ok {
				left.Next = nil
			}
		}
	}
}

func printList(l *ListNode) {
	if l != nil {
		for l.Next != nil {
			fmt.Println(l.Val)
			l = l.Next
		}
		fmt.Println(l.Val)
	}
}

func main() {
	list := CreateList([]int{})
	deleteDuplicates(list)
	printList(list)
}
