# Gotas
Implementing and learning data structures. All code is written in Go.

## Data Structures

- Linked List
    - [x] [Singly Linked List](#singly-linked-list)
    - [x] [Doubly Linked List](#doubly-linked-list)
    - [x] [Circular Linked List](#circular-linked-list)
- Stack
    - [x] [Array](#array-stack)
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


## Linked List

Linked List is as data structure used for storing collections of data, that has the following properties:

- Successive elements are connected by pointers.
- The last element points to `nil`.
- Can grow or shrink in size during execution of the program.
- Can be made just as long as required (until memory exhaustion).
- Does not waste memory space (but takes up a bit more space for the pointers).

The basic operations in all kinds of linked lists implemented here are:

- `Get(index int)`: returns the value of the node at position `index`.
- `Prepend(value T)`: adds a new node at the beginning of the list.
- `Append(value T)`: adds a new node at the end of the list.
- `Add(index int, value T)`: adds a new node at position `index`.
- `Delete(index int)`: removes the node at position `index`.
- `Values() []T`: returns a slice with the values of the nodes in the list.
- `Empty() bool`: returns `true` if the list is empty, `false` otherwise.

The main advantage of linked lists, is that can be expanded in constant time, which is not the case with arrays, since we need to copy all elements to a new array. Thus, we can start with an empty list and add elements as needed. In the other hand, the main disadvantage is that we need to traverse the list to find the `n-th` element, which is not the case with arrays, since we can access the element in constant time. Furthermore, arrays are cache friendly, since all elements are stored in contiguous memory locations, which is not the case with linked lists.

All kinds of linked lists implemented here have the following structure:

```go
type List[T comparable] struct {
    head *node[T]
    tail *node[T]
    size int
}
```

Where `node` is the structure of the node, that can be different in each case, so it is defined in each section.

### Singly Linked List

Also known simply as "linked list," it is a linear data structure where each element is a node containing a value and a pointer to the next node. The last node points to `nil`. It is noteworthy that the value contained in the node can be of any type and is defined at the time of list creation.

The beginning of the list is represented by the `Head` pointer, which points to the first node. The end of the list is represented by the `Tail` pointer, which points to the last node. Some implementations can use only Head pointer, but here we will use both. The list below represents a singly linked list with four nodes.

    Head                        Tail
      |                          |
      v                          v
    +---+    +---+    +---+    +---+
    | A | -> | B | -> | C | -> | D | -> nil
    +---+    +---+    +---+    +---+

The structure of the node is defined as follows:

```go
type node[T any] struct {
    data T
    next *node[T]
}
```

### Doubly Linked List

It is a linear data structure where each element is a node containing a value, a pointer to the next node and a pointer to the previous node. The last node points to `nil`. The beginning of the list is represented by the `Head` pointer, which points to the first node. The end of the list is represented by the `Tail` pointer, which points to the last node. The list below represents a doubly linked list with four nodes.

    Head                           Tail
      |                             |
      v                             v
    +---+     +---+     +---+     +---+
    | A | <-> | B | <-> | C | <-> | D | -> nil
    +---+     +---+     +---+     +---+

The structure of the node is defined as follows:

```go
type node[T any] struct {
    data T
    next *node[T]
    prev *node[T]
}
```

### Circular Linked List

Circular Linked List is similar to a singly linked list, but the last node points to the first node.

The beginning of the list is represented by the `Head` pointer, which points to the first node. The end of the list is represented by the `Tail` pointer, which points to the last node. The list below represents a circular linked list with four nodes.

    Head                        Tail
      |                          |
      v                          v
    +---+    +---+    +---+    +---+
    | A | -> | B | -> | C | -> | D | ---
    +---+    +---+    +---+    +---+   |
      ^                                |
      |--------------------------------+

The structure of the node is defined as follows:

```go
type node[T any] struct {
    data T
    next *node[T]
}
```

## Stack

A stack is a simple data structure used for store elements in a last-in-first-out (LIFO) order. The LIFO order means that the last element added to the stack is the first one to be removed, similar to a pile of plates. When we add a new plate to the pile, we place it on top of the others. When we remove a plate from the pile, we always take the one on top. The stack works the same way. There are special names for the operations on a stack. When we add an element to the stack, we say we `push` it. When we remove an element from the stack, we say we `pop` it. So, basically, we have these two operations and some other auxiliary operations in all kinds of stacks implemented here:

- `Push(value T)`: adds a new element to the top of the stack.
- `Pop() T`: removes the element from the top of the stack and returns it.
- `Peek() T`: returns the element from the top of the stack without removing it.
- `Size() int`: returns the number of elements in the stack.
- `Empty() bool`: returns `true` if the stack is empty, `false` otherwise.
- `Full() bool`: returns `true` if the stack is full, `false` otherwise.
- `Values() []T`: returns a slice with the values of the elements in the stack.

### Array Stack

The array stack is a stack implemented using an array. The array has a fixed size, which is defined at the time of stack creation, we add elements from left to right and we use a variable to keep track of the top of the stack. The stack below represents an array stack with size four and two elements.

         Top
          |
          v
    +---+---+---+---+
    | A | B |   |   |
    +---+---+---+---+

The structure of the stack is defined as follows:

```go
type ArrayStack[T any] struct {
    data     []T
	top      int
    capacity int
}
```

### References

- Data Structures and Algorithms Made Easy: Data Structures and Algorithmic Puzzles by Narasimha Karumanchi
