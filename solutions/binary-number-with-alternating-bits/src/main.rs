fn main() {
    println!("{}", Solution::has_alternating_bits(5));
    println!("{}", Solution::has_alternating_bits(7));
    println!("{}", Solution::has_alternating_bits(11));
    println!("{}", Solution::has_alternating_bits(2));
}

struct Solution;

impl Solution {
    pub fn has_alternating_bits(n: i32) -> bool {
        let mut n = n;
        let mut flag = n % 2 == 0;

        while n > 0 {
            let digit = n % 2;
            let flag_checker = digit != 0;
            if flag_checker == flag {
                return false;
            }

            flag = !flag;

            n /= 2;
        }

        true
    }
}
