fn main() {
    println!("{:?}", Solution::two_sum(vec![2, 7, 11, 15], 9));
    println!("{:?}", Solution::two_sum(vec![2, 3, 4], 6));
    println!("{:?}", Solution::two_sum(vec![-1, 0], -1));
}

struct Solution;

impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut i = 0;
        let mut j = numbers.len() - 1;

        loop {
            let num_i = numbers[i];
            let num_j = numbers[j];
            let sum = num_i + num_j;

            match sum.cmp(&target) {
                std::cmp::Ordering::Greater => j -= 1,
                std::cmp::Ordering::Less => i += 1,
                std::cmp::Ordering::Equal => return vec![i as i32 + 1, j as i32 + 1],
            }
        }
    }
}
