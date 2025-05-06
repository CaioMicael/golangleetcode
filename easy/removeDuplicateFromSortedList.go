package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	aux := head
	for aux != nil && aux.Next != nil {
		if aux.Val == aux.Next.Val {
			aux.Next = aux.Next.Next
		} else {
			aux = aux.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	deleteDuplicates(head)
}
