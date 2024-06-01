package main

import "fmt"

type node struct{
	data int
	next *node
}

type LinkedList struct{
	head *node
}

func (list *LinkedList)InsertAtBegining(data int){
	newNode := &node{
		data : data,
		next : list.head,
	}
	list.head = newNode
}

func (list *LinkedList) Display(){
	current := list.head
	for current != nil{
		fmt.Printf("| %d |->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main(){

	list := LinkedList{}

	list.InsertAtBegining(10)
	list.InsertAtBegining(20)
	list.InsertAtBegining(30)

	fmt.Println("linked List :")
	list.Display()
}

