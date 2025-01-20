fn main() {
    let mut nums = vec![0, 1, 0, 3, 12];
    Solution::move_zeroes(&mut nums);
    println!("{:?}", nums);

    let mut nums = vec![0];
    Solution::move_zeroes(&mut nums);
    println!("{:?}", nums);
}

struct Solution;

impl Solution {
    pub fn move_zeroes(nums: &mut [i32]) {
        let n = nums.len();

        for i in (0..n).rev() {
            if nums[i] == 0 {
                for j in i..n - 1 {
                    let next = nums[j + 1];
                    if next == 0 {
                        break;
                    }
                    nums[j] = next;
                    nums[j + 1] = 0;
                }
            }
        }
    }
}
