class Solution {
    fun searchRange(nums: IntArray, target: Int): IntArray {
        if(nums.isEmpty()) return intArrayOf(-1,-1)

        var left = 0
        var right = nums.size-1
        var finalMid = -1
        while(left <= right){
            val mid = (right + left)/2

            when {
                nums[mid] < target -> left = mid + 1
                nums[mid] > target -> right = mid -1
                nums[mid] == target -> {
                    finalMid = mid
                    break
                }
            }
        }

        if(finalMid == -1 ) return intArrayOf(-1,-1)

        var firstIdx = finalMid
        var endIdx = finalMid

        var leftFirstIdx = firstIdx-1
        var rightEndIdx = endIdx+1

        while(leftFirstIdx >= 0 || rightEndIdx < nums.size){
            var isShouldBreak = true
            if(leftFirstIdx >= 0 && nums[leftFirstIdx] == target){
                firstIdx = leftFirstIdx
                isShouldBreak = false
            }

            if(rightEndIdx < nums.size && nums[rightEndIdx] == target){
                endIdx = rightEndIdx
                isShouldBreak = false
            }

            if(isShouldBreak) break

            leftFirstIdx--
            rightEndIdx++
        }

        return intArrayOf(firstIdx, endIdx)
    }
}
