package main

import "fmt"

/* Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by
continuously following the next pointer. Internally, pos is used to denote the index
of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.*/

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCyCle(head *ListNode) bool {
	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		if (fast == slow) {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	// Test case 1: Linked list không có chu kỳ
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = nil  
	// Kết quả dự kiến: true
	fmt.Println("Test case 1: Không có chu kỳ:", hasCyCle(node1)) 

	// Test case 2: Linked list có chu kỳ
	nodeA := &ListNode{Val: 1}
	nodeB := &ListNode{Val: 2}
	nodeC := &ListNode{Val: 3}
	nodeD := &ListNode{Val: 4}

	nodeA.Next = nodeB
	nodeB.Next = nodeC
	nodeC.Next = nodeD
	nodeD.Next = nodeB  
	// Kết quả dự kiến: false
	fmt.Println("Test case 2: Có chu kỳ:", hasCyCle(nodeA))


	fmt.Printf("Study Pointer")
}