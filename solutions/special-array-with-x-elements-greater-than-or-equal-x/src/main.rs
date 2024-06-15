fn main() {
    println!("{}", Solution::special_array(vec![3, 5]));
    println!("{}", Solution::special_array(vec![0, 0]));
    println!("{}", Solution::special_array(vec![0, 4, 3, 0, 4]));
}

struct Solution;

impl Solution {
    pub fn special_array(nums: Vec<i32>) -> i32 {
        let n = nums.len();

        for i in 0..=n {
            let mut counter = 0;
            for j in 0..n {
                let num = nums[j];
                if num >= i as i32 {
                    counter += 1;
                }
            }
            if i as i32 == counter {
                return counter;
            }
        }

        -1
    }
}
