package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type CircularDoublyLinkedList struct {
	head *Node
}

func (c *CircularDoublyLinkedList) Insert(value int) {
	newNode := &Node{value: value}

	if c.head == nil {
		newNode.next = newNode
		newNode.prev = newNode
		c.head = newNode
		return
	}

	tail := c.head.prev

	tail.next = newNode
	newNode.prev = tail
	newNode.next = c.head
	c.head.prev = newNode
}

// TraverseN moves n steps from the given starting node
// n > 0 moves forward; n < 0 moves backward
func (c *CircularDoublyLinkedList) TraverseN(start *Node, n int) *Node {
	if start == nil || c.head == nil {
		return nil
	}

	curr := start

	if n > 0 {
		for i := 0; i < n; i++ {
			curr = curr.next
		}
	} else {
		for i := 0; i < -n; i++ {
			curr = curr.prev
		}
	}

	return curr
}

func (c *CircularDoublyLinkedList) Print() {
	if c.head == nil {
		fmt.Println("empty")
		return
	}

	fmt.Print(c.head.value, " ")
	curr := c.head.next

	for curr != c.head {
		fmt.Print(curr.value, " ")
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println("Error closing file", err)
			return
		}
	}()

	dial := &CircularDoublyLinkedList{}
	for i := range 100 {
		dial.Insert(i)
	}
	dial.head = dial.TraverseN(dial.head, 50)

	password := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("invalid amount: ", err)
		}
		switch dir {
		case 'L':
			amount = -amount
			dial.head = dial.TraverseN(dial.head, amount)
			if dial.head.value == 0 {
				password++
			}
		case 'R':
			dial.head = dial.TraverseN(dial.head, amount)
			if dial.head.value == 0 {
				password++
			}
		default:
			fmt.Println("invalid direction")
			continue
		}
	}
	fmt.Println("Password:", password)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}
