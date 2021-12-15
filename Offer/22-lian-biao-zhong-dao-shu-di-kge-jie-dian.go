package Offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for ; k > 0 && fast != nil; k-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
