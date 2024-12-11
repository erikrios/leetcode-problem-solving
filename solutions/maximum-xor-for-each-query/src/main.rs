fn main() {
    println!("{:?}", Solution::get_maximum_xor(vec![0, 1, 1, 3], 2));
    println!("{:?}", Solution::get_maximum_xor(vec![2, 3, 4, 7], 3));
    println!("{:?}", Solution::get_maximum_xor(vec![0, 1, 2, 2, 5, 7], 3));
}

struct Solution;

impl Solution {
    pub fn get_maximum_xor(nums: Vec<i32>, maximum_bit: i32) -> Vec<i32> {
        let n = nums.len();

        let mut xor_sum = nums[0];
        for num in nums.iter().take(n).skip(1) {
            xor_sum ^= num;
        }

        let mut results = Vec::with_capacity(n);
        let max_value = 2_i32.pow(maximum_bit as u32) - 1;

        for num in nums.into_iter().rev() {
            let res = xor_sum ^ max_value;
            results.push(res);

            xor_sum ^= num;
        }

        results
    }
}
