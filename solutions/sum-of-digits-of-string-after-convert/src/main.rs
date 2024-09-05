fn main() {
    println!("{}", Solution::get_lucky("iiii".to_string(), 1));
    println!("{}", Solution::get_lucky("leetcode".to_string(), 2));
    println!("{}", Solution::get_lucky("zbax".to_string(), 2));
    println!("{}", Solution::get_lucky("fleyctuuajsr".to_string(), 5));
}

struct Solution;

impl Solution {
    pub fn get_lucky(s: String, k: i32) -> i32 {
        let mut converted = String::new();

        for s in s.chars() {
            let num = (s as u8) - b'a' + 1;
            converted.push_str(&num.to_string());
        }

        let mut sum = 0;

        for _ in 0..k {
            sum = 0;

            for ch in converted.chars() {
                sum += (ch as u8 - b'0') as i32;
            }

            converted = sum.to_string();
        }

        sum
    }
}
