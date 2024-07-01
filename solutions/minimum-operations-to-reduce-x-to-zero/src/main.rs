fn main() {
    println!("{:?}", Solution::min_operations(vec![1, 1, 4, 2, 3], 5));
    println!("{:?}", Solution::min_operations(vec![5, 6, 7, 8, 9], 4));
    println!(
        "{:?}",
        Solution::min_operations(vec![3, 2, 20, 1, 1, 3], 10)
    );
    println!("{:?}", Solution::min_operations(vec![1, 1], 3));
    println!(
        "{:?}",
        Solution::min_operations(
            vec![
                8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904,
                8819, 1231, 6309
            ],
            134365
        )
    );
}

struct Solution;

impl Solution {
    pub fn min_operations(nums: Vec<i32>, x: i32) -> i32 {
        let target = nums.iter().sum::<i32>() - x;
        if target < 0 {
            return -1;
        } else if target == 0 {
            return nums.len() as i32;
        }

        let mut window_start = 0;
        let mut window_end = 0;
        let mut max = -1;
        let mut sum = 0;

        while window_end < nums.len() {
            let num = nums[window_end];
            sum += num;

            while sum > target {
                sum -= nums[window_start];
                window_start += 1;
            }

            if sum == target {
                let range = (window_end - window_start) as i32 + 1;
                max = std::cmp::max(max, range)
            }

            window_end += 1;
        }

        if max == -1 {
            return max;
        }

        nums.len() as i32 - max
    }
}
