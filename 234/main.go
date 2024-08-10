package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	s := []int{1}
	l := sliceToList(s)
	fmt.Println(l)
	l = reverse(l)
	fmt.Println(l)
	
	fmt.Println(isPalindrome(l))
}

func isPalindrome(head *ListNode) bool {
	mid := getMid(head)
	
	rev := reverse(mid.Next)
	
	mid.Next = nil
	
	for rev != nil {
		if head.Val != rev.Val {
			return false
		}
		head = head.Next
		rev = rev.Next
	}
	
	return true
}

func getMid(n *ListNode) *ListNode {
	slow := n
	fast := n.Next
	
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	return slow
}

func reverse(n *ListNode) *ListNode {
	var prev *ListNode
	cur := n
	for cur != nil {
		next := cur.Next
		
		cur.Next = prev
		prev = cur
		cur = next
	}
		
	return prev
}

func sliceToList(s []int) *ListNode {
	head := &ListNode{}
	
	cur := head
	for i, v := range s {
		cur.Val = v
		
		if i < len(s) - 1 {
			next := &ListNode{}
			cur.Next = next
			cur = next
		}
	}
	
	return head
}

func (n *ListNode) String() string {
	var b strings.Builder
	
	for c := n; c != nil; c = c.Next {
		
		b.WriteString(strconv.Itoa(c.Val))
		
		if c.Next != nil {
			b.WriteString(" -> ")
		}
	}
	
	return b.String()
}