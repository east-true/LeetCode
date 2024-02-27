/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	res := &ListNode{
		Val:  head.Val,
		Next: nil,
	}

	for head.Next != nil {
		head = head.Next
		if res.Val > head.Val {
			res = &ListNode{
				Val:  head.Val,
				Next: res,
			}
		} else {
			node := res
			for {
				if node.Next == nil {
					node.Next = &ListNode{
						Val:  head.Val,
						Next: nil,
					}
					break
				} else if node.Next.Val > head.Val {
					node.Next = &ListNode{
						Val:  head.Val,
						Next: node.Next,
					}
					break
				}

				node = node.Next
			}
		}
	}

	return res
}
