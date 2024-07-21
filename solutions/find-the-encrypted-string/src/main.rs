fn main() {
    println!("{}", Solution::get_encrypted_string("dart".to_string(), 3));
    println!("{}", Solution::get_encrypted_string("aaa".to_string(), 1));
}

struct Solution;

impl Solution {
    pub fn get_encrypted_string(s: String, k: i32) -> String {
        let s: Vec<char> = s.chars().collect();
        let mut chars = Vec::with_capacity(s.len());

        for i in 0..s.len() {
            let j = (k as usize + i) % s.len();
            chars.push(s[j]);
        }

        chars.into_iter().collect()
    }
}
