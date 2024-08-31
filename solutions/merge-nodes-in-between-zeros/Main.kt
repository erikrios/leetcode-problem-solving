 class ListNode(var `val`: Int) {
     var next: ListNode? = null
 }

class Solution {
    fun mergeNodes(head: ListNode?): ListNode? {
        var root: ListNode? = null
        var tail: ListNode? = null
        var temp : Int = 0
        var currentNode: ListNode? = head

        while(currentNode != null){
            if(currentNode.`val` == 0 && temp > 0){
                var newNode = ListNode(temp)
                if(root == null){
                    root = newNode
                    tail = newNode
                } else {
                    tail?.next = newNode
                    tail = tail?.next
                }
                temp = 0
            } else {
                temp += currentNode.`val`
            }
            currentNode = currentNode.next
        }

        return root
    }
}

