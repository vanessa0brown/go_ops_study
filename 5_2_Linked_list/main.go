package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(sl []int) *ListNode {
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
	//map
	//TEST
}

func main() {
	list := CreateList([]int{1, 2, 3})
	fmt.Println(list.Val)
	fmt.Println(list.Next.Val)
	fmt.Println(list.Next.Next.Val)
}
