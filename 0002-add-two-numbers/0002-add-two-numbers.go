/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := new(ListNode)
	ptr := res
	for {
		add := ptr.Val
		if l1 != nil {
			add += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			add += l2.Val
			l2 = l2.Next
		}

		ptr.Val = add % 10
		ups := add / 10
		if l1 == nil && l2 == nil && ups == 0 {
			break
		}

		ptr.Next = new(ListNode)
		ptr = ptr.Next
		ptr.Val = ups
	}

	return res
}