package main

import (
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func removeElementsInPlace(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	
	if head == nil {
		return nil
	}
	
	prev := head
	cur := head.Next
	for cur != nil {		
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		
		cur = cur.Next
	}
	
	return head
}


func removeElements(head *ListNode, val int) *ListNode {
	var newHead *ListNode
	var cur *ListNode
	
	for head != nil {
		if head.Val != val {
			if newHead == nil {
				newHead = &ListNode {
					Val: head.Val,
				}
				cur = newHead
			} else {
				cur.Next = &ListNode {
					Val: head.Val,
				}
				cur = cur.Next
			}
		}
		
		head = head.Next
	}
	
	return newHead
}


func (l *ListNode) String() string {
	if l == nil {
		return "nil ListNode"
	}
	var b strings.Builder
	for l != nil {
		_, _ = b.WriteString(strconv.Itoa(l.Val))
		_, _ = b.WriteString(" -> ")
		l = l.Next
	}
	_, _ = b.WriteString("nil")
	return b.String()
}