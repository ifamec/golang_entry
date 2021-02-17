# Chapter 20 Notes - Data Structure & Algorithms

## Sparse Array
A sparse array is an array of data in which many elements have a value of zero. This is in contrast to a dense array, where most of the elements have non-zero values or are “full” of numbers. A sparse array may be treated differently than a dense array in digital data handling.

[Sparse Array Example](01_sparsearray/main.go)

## Queue
- Ordered List (Array, Linked List)
- First In First Out (FIFO)

### Queue in Array
- declare: 
    - need a maxSize
    - need two pointers front rear to indicate
- addQueue
    - rear + 1, front = rear
    - rear == maxSize-1 -> Queue Full
- getQueue
- listQueue

[Array Queue](02_queue/array_queue/main.go)

**Issue:**
- Array Space Waste
- Use `Circle Queue`

### Circle Queue w Array
- tail `(tail+1)%maxSize`
- `(tail+1)%maxSize` == head -> full
- tail == head -> empty
- init: head = 0 tail = 0
- how many data in queue `(tail+maxSize-head)%maxSize`

[circle_queue](02_queue/circle_queue/main.go)

## Linked List

### Singly Linked List

head -> nil, next -> val: X, next -> val: Y, next -> nil

[singly_linked_list](03_linkedlist/singly_linked_list/main.go)

### Doubly Linked List

head -> nil, next <-> prev, val: X, next <-> prev, val: Y, next -> nil

[doubly_linked_list](03_linkedlist/doubly_linked_list/main.go)

### Circular Linked List (Singly)

last element .next -> first element

[circular_linked_list](03_linkedlist/circular_linked_list/main.go)

### Josephus Problem

[josephus_problem](03_linkedlist/josephus_problem/main.go)

## Sort
### Bubble sort
Check [Chapter 8 Bubble Sort](../chapter08/sort01_bubble/main.go)
### Selection sort
[Selection Sort](04_sort/selection_sort/main.go)
### Insertion sort
[Insertion Sort](04_sort/insertion_sort/main.go)
### Quick sort
[Quick Sort](04_sort/quick_sort/main.go)

## Stack
- First In Last Out (FILO)
- stackTop stackBottom -> -1 init
- Use
    - call subProcess
    - iteration
    - statement transfer and calculation
    - binary tree traverse
    - graph depth-first search
  
[stack](05_stack/stack01/main.go)

[stack_calc](05_stack/stack02_calc/main.go)
- NumberStack, OperatorStack
- index := 0
- Exp - string
- if is a number, push into numberStack 
- if is an operator
    - if operatorStack is empty, push
    - if operatorStackTop prior than the one pend to push, 
      pop operator and two numbers, calculate and push into numberStack, 
      then push the new operator into operatorStack
      else operator push into operatorStack
    - everything is done, pop operator and two numbers, calculate and push into numberStack.
      until operatorStack empty
      
## Recursion
- Each function in the recursion will create a new individual space in ram
- Function variables are independent
- Recursion needs to be exited tended
- When the function done / meet return the function will return, follows value returns to the caller. Then GC

[recursion](06_recursion/recursion01/main.go)

[recursion_eg](06_recursion/recursion02/main.go)

## Hash Table
Hash Table is a data structure which stores data in an associative manner. In a hash table, data is stored in an array format, where each data value has its own unique index value. Access of data becomes very fast if we know the index of the desired data.

Example
- Use linked list without head
- Realize CRUD

```
hashtable->
linkArr[0]->Emp0->Emp16->Emp32
linkArr[1]->
linkArr[2]->
linkArr[3]->Emp3->Emp11
linkArr[4]->
linkArr[5]->
linkArr[6]->Emp22
linkArr[7]->
```

```go
type LinkArr struct {
	LinkArr [7]EmpLink
}
type EmpLink struct {
	Head *Emp
}
type Emp struct {
	Id int
	Name string
	Next
}
methods
```

## Binary Tree
A binary tree is a tree data structure in which each node has at most two children, which are referred to as the left child and the right child.

Depth First Traversals:
- Preorder (Root, Left, Right)
- Inorder (Left, Root, Right)
- Postorder (Left, Right, Root)