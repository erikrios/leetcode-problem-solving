fn main() {
    println!("{}", Solution::string_hash("abcd".to_string(), 2));
    println!("{}", Solution::string_hash("mxz".to_string(), 3));
    println!("{}", Solution::string_hash("abcdef".to_string(), 3));
}

struct Solution;

impl Solution {
    pub fn string_hash(s: String, k: i32) -> String {
        let n = s.len();
        let mut hashed = String::with_capacity(n / k as usize);
        let mut sum = 0;

        for (i, char) in s.chars().enumerate() {
            sum += (char as u8 - b'a') as u16;
            if i as i32 % k == k - 1 {
                let hashed_char = (sum % 26) as u8 + b'a';
                hashed.push(hashed_char as char);
                sum = 0;
            }
        }

        hashed
    }
}
