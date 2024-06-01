fn main() {
    println!(
        "{}",
        Solution::number_of_pairs(vec![1, 3, 4], vec![1, 3, 4], 1)
    );
    println!(
        "{}",
        Solution::number_of_pairs(vec![1, 2, 4, 12], vec![2, 4], 3)
    );
}

struct Solution;

impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i32 {
        let n = nums1.len();
        let m = nums2.len();

        let mut total = 0;
        for i in 0..n {
            let num1 = nums1[i];
            for j in 0..m {
                let num2 = nums2[j];
                if num1 % (num2 * k) == 0 {
                    total += 1;
                }
            }
        }

        total
    }
}
