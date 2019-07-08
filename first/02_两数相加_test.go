package first

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return head.Next
}
func print(n *ListNode) {
	for n != nil {
		fmt.Print(n.Val, "->")
		n = n.Next
	}
	fmt.Print("nil\n")
}
func TestAddTwoNumbers(t *testing.T) {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	result1 := addTwoNumbers(node1, node2)
	print(result1)

	node3 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	node4 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result2 := addTwoNumbers(node3, node4)
	print(result2)
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
