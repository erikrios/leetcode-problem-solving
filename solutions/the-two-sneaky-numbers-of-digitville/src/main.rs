fn main() {
    println!("{:?}", Solution::get_sneaky_numbers(vec![0, 1, 1, 0]));
    println!("{:?}", Solution::get_sneaky_numbers(vec![0, 3, 2, 1, 3, 2]));
    println!(
        "{:?}",
        Solution::get_sneaky_numbers(vec![7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2])
    );
}

struct Solution;

impl Solution {
    pub fn get_sneaky_numbers(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len() - 2;
        let mut appear_checkers = vec![false; n];
        let mut sneaky_numbers = Vec::with_capacity(2);

        let mut i = 0;
        while sneaky_numbers.len() < 2 {
            let num = nums[i] as usize;
            if appear_checkers[num] {
                sneaky_numbers.push(num as i32);
            } else {
                appear_checkers[num] = !appear_checkers[num];
            }
            i += 1;
        }

        sneaky_numbers
    }
}
