use std::cmp::max;

fn main() {
    println!("{}", Solution::minimum_chairs("EEEEEEE".to_string()));
    println!("{}", Solution::minimum_chairs("ELELEEL".to_string()));
    println!("{}", Solution::minimum_chairs("ELEELEELLL".to_string()));
}

struct Solution;

impl Solution {
    pub fn minimum_chairs(s: String) -> i32 {
        let mut res = 0;
        let mut counter = 0;

        for character in s.chars() {
            if character == 'E' {
                counter += 1;
            } else {
                counter -= 1;
            }

            res = max(res, counter);
        }

        res
    }
}
