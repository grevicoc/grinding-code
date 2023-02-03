/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var arrNode []*ListNode
    i := 0
    temp := head
    for temp != nil{
        i++
        arrNode = append(arrNode,temp)
        temp = temp.Next
    }

    idxToRemove := len(arrNode) - n
    if idxToRemove == 0{
        head = arrNode[idxToRemove].Next
        return head
    }
    arrNode[idxToRemove-1].Next = arrNode[idxToRemove].Next
    return head
}