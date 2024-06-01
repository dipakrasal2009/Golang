package main

import "fmt"

type node struct{
	data int
	next *node
}

type linkedlist struct{
	head *node
}

func (list *linkedlist) InsertFirst(data int){
	newn := &node{
		data : data,
		next : list.head,
	}
	list.head = newn
}

func (list *linkedlist) InsertLast(data int) {
	newNode := &node{
		data: data,
		next: nil,
	}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *linkedlist) Display(){
	current := list.head

	for current != nil{
		fmt.Printf("|%d|=>",current.data)
		current = current.next
	}
	fmt.Println("NULL")
}

func (list *linkedlist) DeleteFirst(){
	if list.head == nil{
		fmt.Println("nothing to delete linked list is already empty..!")
		return
	}
	list.head = list.head.next
}

func (list *linkedlist) DeleteLast(){

	if list.head == nil {
		fmt.Println("nothing to delete")
		return
	}
	list.head = list.head.next

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = nil
}

func (list *linkedlist) InsertAtPos(data int,pos int){


func main(){
	list := linkedlist{}

	list.InsertFirst(11)
	list.InsertFirst(12)
	fmt.Println("Insert First Linked List :")
	list.Display()

	list.InsertLast(13)
	list.InsertLast(14)
	fmt.Println("Insert Last Linked List : ")
	list.Display()

	list.DeleteFirst()
	fmt.Println("Delete First Linked List :")
	list.Display()

	list.DeleteLast()
	fmt.Println("Delete Last inked List : ")
	list.Display()
}
