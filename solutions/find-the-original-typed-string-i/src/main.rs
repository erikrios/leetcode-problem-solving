fn main() {
    println!("{}", Solution::possible_string_count("abbcccc".to_string()));
    println!("{}", Solution::possible_string_count("abcd".to_string()));
    println!("{}", Solution::possible_string_count("aaaa".to_string()));
}

struct Solution;

impl Solution {
    pub fn possible_string_count(word: String) -> i32 {
        let word = word.into_bytes();
        let n = word.len();

        let mut counter = 0;

        for i in 0..n {
            let ch = word[i];
            if i == n - 1 || word[i + 1] != ch {
                continue;
            }
            counter += 1;
        }

        counter + 1
    }
}
