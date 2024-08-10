class Solution {
    fun reverseBetween(head: ListNode?, m: Int, n: Int): ListNode? {
        if(m == n) return head

        var iteration = 1 // 2
        var first:ListNode? = null
        var currentNode: ListNode? = head
        while(iteration < m){ // n
            first = currentNode
            currentNode = currentNode?.next
            iteration++
        }


        var newList: ListNode? = null
        var tail = currentNode

        while (iteration in m .. n){
            val next = currentNode?.next
            currentNode?.next = newList
            newList = currentNode
            currentNode = next
            iteration++
        }

        first?.next = newList
        tail?.next = currentNode

        return if (m > 1) head else newList
    }
}
