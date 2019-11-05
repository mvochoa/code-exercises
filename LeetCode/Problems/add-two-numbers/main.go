package main

// ListNode estructura para los nodos
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	o := &ListNode{Val: l1.Val + l2.Val}
	l := o
	e := 0

	if o.Val > 9 {
		o.Val = o.Val % 10
		e = 1
		if l1.Next == nil && l2.Next == nil {
			o.Next = &ListNode{Val: e}
			l = o
		}
	}

	for l1.Next != nil || l2.Next != nil {
		if l1.Next != nil && l2.Next == nil {
			l.Next = &ListNode{
				Val: l1.Next.Val + e,
			}
			l1 = l1.Next
		} else if l1.Next == nil && l2.Next != nil {
			l.Next = &ListNode{
				Val: l2.Next.Val + e,
			}
			l2 = l2.Next
		} else {
			l.Next = &ListNode{
				Val: l1.Next.Val + l2.Next.Val + e,
			}
			l1 = l1.Next
			l2 = l2.Next
		}

		if l.Next.Val > 9 {
			e = 1
			l.Next.Val = l.Next.Val % 10
		} else {
			e = 0
		}

		l = l.Next
	}

	if e == 1 {
		l.Next = &ListNode{Val: e}
	}

	return o
}

func main() {}
