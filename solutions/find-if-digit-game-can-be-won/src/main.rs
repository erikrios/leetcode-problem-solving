fn main() {
    println!("{}", Solution::can_alice_win(vec![1, 2, 3, 4, 10]));
    println!("{}", Solution::can_alice_win(vec![1, 2, 3, 4, 5, 14]));
    println!("{}", Solution::can_alice_win(vec![5, 5, 5, 25]));
}

struct Solution;

impl Solution {
    pub fn can_alice_win(nums: Vec<i32>) -> bool {
        let mut sum = 0;

        for &num in nums.iter() {
            if num > 9 {
                sum += num
            } else {
                sum -= num
            }
        }

        return sum != 0;
    }
}
