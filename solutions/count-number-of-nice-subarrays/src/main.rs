fn main() {
    println!("{}", Solution::number_of_subarrays(vec![1, 1, 2, 1, 1], 3));
    println!("{}", Solution::number_of_subarrays(vec![2, 4, 6], 1));
    println!(
        "{}",
        Solution::number_of_subarrays(vec![2, 2, 2, 1, 2, 2, 1, 2, 2, 2], 2)
    );
}

struct Solution;

impl Solution {
    pub fn number_of_subarrays(nums: Vec<i32>, k: i32) -> i32 {
        let mut total = 0;
        let mut window_start = 0;
        let mut odd_counter = 0;
        let mut temp = 0;

        for &num in &nums {
            if num % 2 == 1 {
                odd_counter += 1;
                temp = 0;
            }

            while odd_counter == k {
                temp += 1;
                if nums[window_start] % 2 == 1 {
                    odd_counter -= 1;
                }
                window_start += 1;
            }
            total += temp;
        }

        total
    }
}
