package main

import "testing"

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		vs []int
		toRemove int
		want []int
	} {
		{
			vs: []int{1, 2, 6, 3, 4, 5, 6},
			toRemove: 6,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			vs: nil,
			toRemove: 1,
			want: nil,
		},
		{
			vs: []int{7, 7, 7, 7},
			toRemove: 7,
			want: nil,
		},
		{
			vs: []int{1, 2, 2, 1},
			toRemove: 2,
			want: []int{1, 1},
		},
	}
	
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			ll := sliceToList(test.vs)
			got := removeElements(ll, test.toRemove)
			want := sliceToList(test.want)
			if !listsEqual(got, want) {
				t.Errorf("got: %q want: %q", got.String(), want.String())
			}
		})
	}
}

// sliceToList returns a linked list representing the given slice.
// The input slice must have at least one element.
func sliceToList(vs []int) *ListNode {
	if len(vs) == 0 {
		return nil
	}
	
	head := ListNode {
		Val: vs[0],
	}
	
	cur := &head
	for i := 1; i < len(vs); i++ {
		cur.Next = &ListNode {
			Val: vs[i],
		}
		cur = cur.Next
	}
	
	return &head
}

func TestSliceToList(t *testing.T) {
	tests := [][]int {
		{1},
		{1, 2},
		{1, 2, 3},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := sliceToList(test)
			if !sliceAndListEqual(test, got) {
				t.Errorf("input slice: %v, output list: %v", test, got.String())
			}
		})	
	}
}


func sliceAndListEqual(s []int, h *ListNode) bool {
	for _, v := range s {
		if h == nil { // If the linked list is shorter than the slice.
			return false
		}
		
		if v != h.Val {
			return false
		}
		
		h = h.Next
	}
	
	// Ensure the linked list is not longer than the slice.
	return h == nil 
}


func TestListNodeString(t *testing.T) {
	tests := []struct {
		s []int
		want string
	} {
		{
			s: []int{1},
			want: "1 -> nil",
		},
		{
			s: []int{1, 2, 3},
			want: "1 -> 2 -> 3 -> nil",
		},
		{
			s: []int{1, 2, 3, 4},
			want: "1 -> 2 -> 3 -> 4 -> nil",
		},
	}
	
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			ll := sliceToList(test.s)
			got := ll.String()
			if got != test.want {
				t.Fatalf("got: %q want: %q", got, test.want)
			}
		})
	}
}

func listsEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	
	return (l1 == nil) == (l2 == nil)
}