fn main() {
    println!("{}", Solution::min_operations(3));
    println!("{}", Solution::min_operations(6));
}

struct Solution;

impl Solution {
    pub fn min_operations(n: i32) -> i32 {
        n * n / 4
    }
}
