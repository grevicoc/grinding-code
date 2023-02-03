/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
    var arrNode []*ListNode
    i := 0
    temp := head
    for temp != nil{
        i++
        arrNode = append(arrNode,temp)
        temp = temp.Next
    }

    mid := i/2
    return arrNode[mid]
}