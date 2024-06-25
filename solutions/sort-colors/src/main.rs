fn main() {
    let mut nums = vec![2, 0, 2, 1, 1, 0];
    Solution::sort_colors(&mut nums);
    println!("{:?}", &nums);
    let mut nums = vec![2, 0, 1];
    Solution::sort_colors(&mut nums);
    println!("{:?}", &nums);
}

struct Solution;

impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut colors_counter = [0; 3];

        for &num in nums.iter() {
            colors_counter[num as usize] += 1;
        }

        let mut index = 0;
        for (color, &count) in colors_counter.iter().enumerate() {
            for _ in 0..count {
                nums[index] = color as i32;
                index += 1;
            }
        }
    }
}
