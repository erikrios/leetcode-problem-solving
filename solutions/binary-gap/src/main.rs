fn main() {
    println!("{}", Solution::binary_gap(22));
    println!("{}", Solution::binary_gap(8));
    println!("{}", Solution::binary_gap(5));
    println!("{}", Solution::binary_gap(1));
}

struct Solution;

impl Solution {
    pub fn binary_gap(n: i32) -> i32 {
        let mut n = n;
        let mut should_counting = false;
        let mut longest_distance = 0;
        let mut distance = 0;

        while n > 0 {
            let binary_digit = n % 2;
            n /= 2;

            if should_counting {
                distance += 1;
            }

            if binary_digit == 1 {
                should_counting = true;
                longest_distance = longest_distance.max(distance);
                distance = 0;
            }
        }

        longest_distance
    }
}
