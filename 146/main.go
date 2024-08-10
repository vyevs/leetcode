package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	c := Constructor(10)
	
	type instr struct {
		typ string
		v1, v2 int
	}
	
	// [10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]
	instrs := []instr{
		{ "put", 10, 13 },
		{ "put", 3, 17 },
		{ "put", 6, 11 },
		{ "put", 10, 5 },
		{ "put", 9, 10 },
		{ "get", 13, 0 },
		{ "put", 2, 19 },
		{ "get", 2, 0 },
		{ "get", 3, 0 },
		{ "put", 5, 25 },
		{ "get", 8, 0 },
		{ "put", 9, 22 },
		{ "put", 5, 5 },
		{ "put", 1, 30 },
		{ "get", 11, 0 },
		{ "put", 9, 12 },
		{ "get", 7, 0 },
		{ "get", 5, 0 },
		{ "get", 8, 0 },
		{ "get", 9, 0 },
		{ "put", 4, 30 },
		{ "put", 9, 3 },
		{ "get", 9, 0 },
		{ "get", 10, 0 },
		{ "get", 10, 0 },
		{ "put", 6, 16 },
		{ "put", 3, 1 },
		{ "get", 3, 0 },
		{ "put", 10, 11 },
		{ "get", 8, 0 },
		{ "put", 2, 14 },
	}
	
	for _, instr := range instrs {
		if instr.typ == "put" {
			c.Put(instr.v1, instr.v2)
		} else {
			v := c.Get(instr.v1)
			fmt.Printf("%d -> %d\n", instr.v1, v)
		}
		fmt.Println(c)
	}
}

type list struct {
	key int
	prev *list
	next *list
}

type entry struct {
	val int
	node *list
}


type LRUCache struct {
	entries map[int]entry
	sentinel *list
	size int
	capacity int
}


func Constructor(capacity int) LRUCache {
    c := LRUCache {
		entries: make(map[int]entry, capacity),
		capacity: capacity,
		sentinel: &list{},
	}
	c.sentinel.next = c.sentinel
	c.sentinel.prev = c.sentinel
	
	return c
}

func (c *LRUCache) Get(key int) int {
	fmt.Printf("Get(%d)\n", key)
    e, found := c.entries[key]
	if !found {
		return -1
	}
	
	disconnectNode(e.node)
	c.addHead(e.node)	
	
	return e.val
}

func (c *LRUCache) Put(key int, value int)  {
	fmt.Printf("Put(%d, %d)\n", key, value)
		
	e, have := c.entries[key]
	if have {
		disconnectNode(e.node)
		c.addHead(e.node)
		
		e.val = value
		c.entries[key] = e
		
	
	} else {
		node := &list{
			key: key,
		}
		
		e := entry {
			val: value,
			node: node,
		}
		
		if c.size == c.capacity {
			// We are full, so remove the LRU entry, which is the tail.
			delete(c.entries, c.sentinel.prev.key)
			disconnectNode(c.sentinel.prev)
		} else {
			c.size++
		}
		c.addHead(node)
		
		c.entries[key] = e
	}
}

func disconnectNode(node *list) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addHead(node *list) {
	node.prev = c.sentinel
	node.next = c.sentinel.next
	c.sentinel.next.prev = node
	c.sentinel.next = node
}
 
func (c LRUCache) String() string {
	if c.size == 0 {
		return "empty LRUCache"
	}
	
	var b strings.Builder	
	for cur := c.sentinel.next; cur != c.sentinel; cur = cur.next {
		b.WriteString(strconv.Itoa(cur.key))
		b.WriteRune(':')
		e := c.entries[cur.key]
		b.WriteString(strconv.Itoa(e.val))
		
		if cur.next != c.sentinel {
			b.WriteString(" <-> ")
		}
	}
	
	return b.String()
}
