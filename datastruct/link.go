package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var cur *LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}
func main() {
	// 字符串切片
	s := "hello world"
	src := []rune(s)
	fmt.Println(src)
	head := &LNode{}
	fmt.Println("就地逆序")
	CreateNode(head, 8)
	PrintNode("逆序前：", head)
	RevLink(head)
	PrintNode("逆序后：", head)
}

//1->2->3->4
//2->1->4->3
func RevLink(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head.Next //1
	pre := head      //nil
	var next *LNode

	for cur != nil && cur.Next != nil {
		//2
		next = cur.Next.Next
		//1
		pre.Next = cur.Next
		//3
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
	return
}
