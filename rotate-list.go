package main

import "fmt"

/**
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

输入：head = [0,1,2], k = 4
输出：[2,0,1]

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// head := &ListNode{Val: 0}
	// head.Next = &ListNode{Val: 1}
	// head.Next.Next = &ListNode{Val: 2}

	// res := head
	// cnt := 2
	// for cnt > 0 {
	// 	res = rotateOne(res)
	// 	cnt--
	// }
	res := rotateRight(head, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	// 先遍历一遍获取链表大小n
	// k%n=i,只需要执行i遍操作即可（执行n遍回到原点）

	if head == nil || k == 0 {
		return head
	}

	cnt := 1
	index := head
	for index.Next != nil {
		index = index.Next
		cnt++
	}

	cnt = k % cnt

	res := head
	for cnt > 0 {
		res = rotateOne(res)
		cnt--
	}

	return res
}

/*
* 执行一次右移，则整个链表需要操作一次，即把最后一个放在第一位
 */
func rotateOne(head *ListNode) *ListNode {
	// 找出倒数第二位
	index := head
	for index.Next.Next != nil {
		index = index.Next
	}
	tail := index.Next

	tail.Next = head
	index.Next = nil

	return tail
}
