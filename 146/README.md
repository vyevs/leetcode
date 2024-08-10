LRU cache implementation that uses a map and a circular doubly linked list.
The linked list uses a sentinel node as the root, so we don't have to maintain a head and a tail.
See https://web.eecs.utk.edu/~bvanderz/teaching/cs140Fa09/notes/Dllists for a wonderful overview.