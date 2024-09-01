fn main() {
    println!("{}", Solution::generate_key(1, 10, 1000));
    println!("{}", Solution::generate_key(987, 879, 789));
    println!("{}", Solution::generate_key(1, 2, 3));
}

struct Solution;

impl Solution {
    pub fn generate_key(num1: i32, num2: i32, num3: i32) -> i32 {
        let mut num1 = num1;
        let mut num2 = num2;
        let mut num3 = num3;

        let mut multiplier = 1;
        let mut key = 0;

        while num1 > 0 || num2 > 0 || num3 > 0 {
            let digit = if num1 > 0 {
                let res = num1 % 10;
                num1 /= 10;
                res
            } else {
                0
            };

            let digit2 = if num2 > 0 {
                let res = num2 % 10;
                num2 /= 10;
                res
            } else {
                0
            };

            let digit = std::cmp::min(digit, digit2);

            let digit2 = if num3 > 0 {
                let res = num3 % 10;
                num3 /= 10;
                res
            } else {
                0
            };

            let digit = std::cmp::min(digit, digit2) * multiplier;
            multiplier *= 10;
            key += digit;
        }

        key
    }
}
