# Gotas
Implementing and learning data structures. All code is written in Go.

## Data Structures

- Lists
    - [x] [Singly Linked List](#singly-linked-list)
    - [ ] Doubly Linked List
    - [ ] Circular Linked List
- Stack
    - [ ] Array
    - [ ] Linked List
- Queue
    - [ ] Array
    - [ ] Linked List
- Tree
    - [ ] Binary Tree
    - [ ] Binary Search Tree
- Graph
    - [ ] Adjacency List
    - [ ] Adjacency Matrix
- Heap
    - [ ] Min Heap
    - [ ] Max Heap

## Singly Linked List

Also known simply as "linked list," it is a linear data structure where each element is a node containing a value and a pointer to the next node. The last node points to `nil`. It is noteworthy that the value contained in the node can be of any type and is defined at the time of list creation.

The beginning of the list is represented by the `Head` pointer, which points to the first node. The end of the list is represented by the `Tail` pointer, which points to the last node. Some implementations can use only Head pointer, but here we will use both. The list below represents a singly linked list with four nodes.

    Head                        Tail
      |                          |
      v                          v
    +---+    +---+    +---+    +---+
    | A | -> | B | -> | C | -> | D | -> nil
    +---+    +---+    +---+    +---+

The structure of the list is defined as follows:

```go
type List[T comparable] struct {
    head *node[T]
    tail *node[T]
    size int
}
```

Where `node` is the structure of the node.

The structure of the node is defined as follows:

```go
type node[T any] struct {
    data T
    next *node[T]
}
```
In the representation above:
- Each box represents a node.
- The number inside each box represents the value stored in that node.
- The arrow (->) represents the pointer from one node to the next.
- The arrow pointing to nil indicates the end of the list.

The basic operations implemented here are:
- `Get(index int)`: returns the value of the node at position `index`.
- `Prepend(value T)`: adds a new node at the beginning of the list.
- `Append(value T)`: adds a new node at the end of the list.
- `Add(index int, value T)`: adds a new node at position `index`.
- `Delete(index int)`: removes the node at position `index`.
- `Values() []T`: returns a slice with the values of the nodes in the list.
- `Empty() bool`: returns `true` if the list is empty, `false` otherwise.
