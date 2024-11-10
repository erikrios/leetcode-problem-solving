fn main() {
    println!("{:?}", Solution::min_bitwise_array(vec![2, 3, 5, 7]));
    println!("{:?}", Solution::min_bitwise_array(vec![11, 13, 31]));
}

struct Solution;

impl Solution {
    pub fn min_bitwise_array(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();

        let mut results = Vec::with_capacity(n);

        'outer: for num in nums {
            for i in 0..num {
                if i | (i + 1) == num {
                    results.push(i);
                    continue 'outer;
                }
            }
            results.push(-1);
        }

        results
    }
}
