fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    pub fn longest_ones(nums: Vec<i32>, k: i32) -> i32 {
        let mut window_start = 0;
        let mut longest = 0;
        let mut zero_counter = 0;

        for (window_end, &num) in nums.iter().enumerate() {
            if num == 0 {
                zero_counter += 1;
            }

            while zero_counter == k + 1 {
                if nums[window_start] == 0 {
                    zero_counter -= 1;
                }
                window_start += 1;
            }
            longest = std::cmp::max(longest, (window_end - window_start) as i32 + 1);
        }

        longest
    }
}
