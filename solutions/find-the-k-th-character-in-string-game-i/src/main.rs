fn main() {
    println!("{}", Solution::kth_character(5));
    println!("{}", Solution::kth_character(10));
}

struct Solution;

impl Solution {
    pub fn kth_character(k: i32) -> char {
        let mut word = String::from("a");

        while word.len() < k as usize {
            let mut new_word = String::with_capacity(word.len() * 2);
            let word_bytes = word.clone();
            for &ch in word_bytes.as_bytes() {
                let ch = if ch == b'z' { b'a' } else { ch + 1 };
                new_word.push(ch as char);
            }
            word.push_str(&new_word);
        }

        word.chars().nth(k as usize - 1).unwrap()
    }
}
