package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	hair := &ListNode{}
	tail := hair
	p, q := l1, l2
	for p != nil || q != nil {
		sum := carry
		if p != nil {
			sum += p.Val
		}
		if q != nil {
			sum += q.Val
		}
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}
	return hair.Next
}

func main() {
	// coding
}
