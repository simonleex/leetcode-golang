package leetcode_golang


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	var vals []int
	val := 0
	for l1 != nil || l2 != nil || val != 0{
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		vals = append(vals,val%10)
		val = val / 10
	}

	head := &ListNode{vals[0], nil}
	tail := head

	for i := 1; i < len(vals); i ++ {
		Node := &ListNode{vals[i], nil}
		tail.Next = Node
		tail = Node
	}

	return head
}