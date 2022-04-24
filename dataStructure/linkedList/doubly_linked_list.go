package linkedList

import "fmt"

/*NOTE:

A Doubly Linked List contains three fields in its node. 

Data field 
One Next pointer points to the next node in the list
One Previous pointer which points to the previous node in the list
Here “Data” and “Next” fields are the same as in the case of a singly linked list. “Previous” pointer field is new which makes a linked 
list a doubly linked list.

Refer to this article to learn about single linked list

https://golangbyexample.com/singly-linked-list-in-golang/

Below is an example of a doubly linked list. Prev pointer of head (start) node points to Null. Similarly, the Next pointer of the 
last node points to Null.

```double linked list img link: https://i2.wp.com/golangbyexample.com/wp-content/uploads/2021/06/dll.png?resize=768%2C120&ssl=1```

Program:
To implement a doubly linked list in Go language, create a node struct with data, prev pointer, and next pointer, methods to add nodes 
in the doubly linked list (from the front end from the back end both), and methods to traverse forward and reverse.


*/

type node struct {
	data string
	prev *node
	next *node
}

type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (d *doublyLinkedList) AddFrontNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *doublyLinkedList) AddEndNodeDLL(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}
func (d *doublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkedList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkedList) Size() int {
	return d.len
}
func DoublyLinkedListExample() {
	doublyList := initDoublyList()
	fmt.Printf("Add Front Node: C\n")
	doublyList.AddFrontNodeDLL("C")
	fmt.Printf("Add Front Node: B\n")
	doublyList.AddFrontNodeDLL("B")
	fmt.Printf("Add Front Node: A\n")
	doublyList.AddFrontNodeDLL("A")
	fmt.Printf("Add End Node: D\n")
	doublyList.AddEndNodeDLL("D")
	fmt.Printf("Add End Node: E\n")
	doublyList.AddEndNodeDLL("E")

	fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())

	err := doublyList.TraverseForward()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = doublyList.TraverseReverse()
	if err != nil {
		fmt.Println(err.Error())
	}
}
