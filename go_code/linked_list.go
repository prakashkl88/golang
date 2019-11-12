// Go program

/*
Linked List implementation in GO
insert and print functionalities 
*/

package main
import "fmt"

//define node structure
type Node struct{
	data int
	next *Node
}

//define linked-list structure
type List struct{
	head *Node
}

// write linked-list methods (l *List)
func (l *List) insert(val int) {
	node := &Node{data: val}
	if l.head == nil{
		l.head = node
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}
}

func (l *List) printIterative() {
	curr := l.head 
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func (l *List) printRecursiveUtil(curr *Node) {
	if curr == nil{
		return
	}
	fmt.Println(curr.data)
	curr = curr.next
	l.printRecursiveUtil(curr)
}

func (l *List) printRecursive() {
	curr := l.head
	l.printRecursiveUtil(curr)
}

//Driver code for testing
func main() {
	list := List{}
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.insert(5)
	list.insert(6)
	list.printIterative()
	list.printRecursive()
}
