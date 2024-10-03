package main

import (
	"fmt"
	"strings"
)

func main() {
	if false {
		a := Constructor()

		a.Inc("a")
		a.Inc("b")
		a.Inc("c")

		fmt.Println(a, "\n")
		a.Inc("a")
		a.Inc("a")
		a.Inc("a")

		fmt.Println(a)

		fmt.Println(a.GetMinKey())
		fmt.Println(a.GetMaxKey())
	}
	if false {
		a := Constructor()

		a.Inc("hello")
		a.Inc("hello")

		fmt.Println(a.GetMinKey())
		fmt.Println(a.GetMaxKey())

		a.Inc("leet")
		fmt.Println(a.GetMinKey())
		fmt.Println(a.GetMaxKey())
	}

	if false {

		a := Constructor()

		a.Inc("hello")
		a.Inc("world")
		a.Inc("leet")
		a.Inc("code")
		a.Inc("ds")
		//fmt.Println(a, "\n")
		a.Inc("leet")

		//fmt.Println(a)

		fmt.Println(a.GetMaxKey())
	}

	if true {
		a := Constructor()

		a.Inc("hello")
		a.Inc("l")
		a.Inc("l")
		a.Inc("l")

		a.Inc("k")
		a.Inc("k")
		a.Inc("k")

		a.Inc("j")
		a.Inc("j")
		a.Inc("j")

		a.Dec("j")
		a.Dec("k")

		fmt.Println(a.GetMaxKey())
	}
}

type node struct {
	strs       map[string]struct{}
	next, prev *node
	ct         int
}

type AllOne struct {
	sentinel *node
	m        map[string]*node
}

func Constructor() AllOne {
	var sentinel node
	sentinel.next = &sentinel
	sentinel.prev = &sentinel
	return AllOne{
		sentinel: &sentinel,
		m:        make(map[string]*node, 5*10e3),
	}
}

func (a *AllOne) Inc(key string) {
	cur, found := a.m[key]
	if found {

		newCt := cur.ct + 1

		next := cur.next

		if next.ct == newCt {
			// Our next node is the correct one for key. Add key to it.
			next.strs[key] = struct{}{}
			a.m[key] = next

		} else {
			// The next node does not have the right count, so add a new node between cur and next.
			a.addNewNodeAfter(cur, newCt, key)
		}

		removeKeyFromNode(cur, key)
	} else {
		a.addNewKey(key)
	}

}

func removeKeyFromNode(n *node, key string) {
	delete(n.strs, key)
	// If the node we moved key from is now empty, take it out of the list.
	if len(n.strs) == 0 {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
}

func (a *AllOne) addNewKey(key string) {
	// Our key has not been seen, so it belongs in a bucket with frequency 1, which can only be the head.
	// But a bucket with freq 1 is not always present.
	candidate := a.sentinel.next

	// If our candidate does not fit the bill, make a new node and make it the head.
	if candidate == a.sentinel || candidate.ct != 1 {
		newNode := node{
			strs: map[string]struct{}{
				key: {},
			},
			prev: a.sentinel,
			next: a.sentinel.next,
			ct:   1,
		}
		a.sentinel.next.prev = &newNode
		a.sentinel.next = &newNode
		a.m[key] = &newNode
	} else {
		// Otherwise, our candidate works.
		candidate.strs[key] = struct{}{}
		a.m[key] = candidate
	}
}

func (a *AllOne) addNewNodeAfter(prev *node, ct int, firstKey string) *node {
	newNode := node{
		ct:   ct,
		prev: prev,
		next: prev.next,
		strs: map[string]struct{}{
			firstKey: {},
		},
	}

	prev.next.prev = &newNode
	prev.next = &newNode

	a.m[firstKey] = &newNode

	return &newNode
}

func (a *AllOne) Dec(key string) {
	cur := a.m[key]
	prev := cur.prev

	newCt := cur.ct - 1

	removeKeyFromNode(cur, key)

	if newCt == 0 {
		delete(a.m, key)

	} else {
		if prev == a.sentinel || prev.ct != newCt {
			a.addNewNodeAfter(prev, newCt, key)

		} else if prev.ct == newCt {
			prev.strs[key] = struct{}{}
			a.m[key] = prev
		}
	}
}

func (a *AllOne) GetMaxKey() string {
	tail := a.sentinel.prev
	if tail == a.sentinel {
		return ""
	}
	for k := range tail.strs {
		return k
	}
	panic("node has no strs")
}

func (a *AllOne) GetMinKey() string {
	head := a.sentinel.next
	if head == a.sentinel {
		return ""
	}
	for k := range head.strs {
		return k
	}
	panic("node has no strs")
}

func (a AllOne) String() string {
	var b strings.Builder

	b.WriteString("forward:  ")
	b.WriteString("sentinel ")
	for cur := a.sentinel.next; cur != a.sentinel; cur = cur.next {
		b.WriteString(fmt.Sprintf("<-> {%d [", cur.ct))

		for k := range cur.strs {
			b.WriteString(k)
		}

		b.WriteString("]} ")
	}
	b.WriteString("<-> sentinel")

	b.WriteRune('\n')

	b.WriteString("backward: ")
	b.WriteString("sentinel ")
	for cur := a.sentinel.prev; cur != a.sentinel; cur = cur.prev {
		b.WriteString(fmt.Sprintf("<-> {%d [", cur.ct))
		for k := range cur.strs {
			b.WriteString(k)
		}

		b.WriteString("]} ")
	}
	b.WriteString("<-> sentinel")

	return b.String()
}
