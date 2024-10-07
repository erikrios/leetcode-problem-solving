fn main() {
    println!("{}", Solution::min_changes(13, 4));
    println!("{}", Solution::min_changes(21, 21));
    println!("{}", Solution::min_changes(14, 13));
    println!("{}", Solution::min_changes(1, 135));
}

struct Solution;

impl Solution {
    pub fn min_changes(n: i32, k: i32) -> i32 {
        let mut n = n;
        let mut k = k;

        let mut counter = 0;
        while n > 0 || k > 0 {
            let n_digit = n % 2;
            let k_digit = k % 2;

            n /= 2;
            k /= 2;

            if n_digit == k_digit {
                continue;
            }

            if k_digit == 1 {
                return -1;
            } else {
                counter += 1;
            }
        }

        counter
    }
}
