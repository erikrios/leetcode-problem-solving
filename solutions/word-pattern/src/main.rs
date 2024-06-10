use std::collections::HashMap;
use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::word_pattern("abba".to_string(), "dog cat cat dog".to_string())
    );
    println!(
        "{}",
        Solution::word_pattern("abba".to_string(), "dog cat cat fish".to_string())
    );
    println!(
        "{}",
        Solution::word_pattern("aaaa".to_string(), "dog cat cat dog".to_string())
    );
    println!(
        "{}",
        Solution::word_pattern("abba".to_string(), "dog dog dog dog".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn word_pattern(pattern: String, s: String) -> bool {
        let pattern: Vec<char> = pattern.chars().collect();
        let s: Vec<&str> = s.split_whitespace().collect();

        if pattern.len() != s.len() {
            return false;
        }

        let mut hash_map = HashMap::<char, &str>::new();
        let mut char_set = HashSet::<char>::new();
        let mut str_set = HashSet::<&str>::new();

        for i in 0..pattern.len() {
            let char = pattern[i];
            let word = s[i];
            match hash_map.get(&char) {
                Some(&val) => {
                    if val != word {
                        return false;
                    }
                }
                None => {
                    hash_map.insert(char, word);
                }
            }
            char_set.insert(char);
            str_set.insert(word);
        }

        if char_set.len() != str_set.len() {
            return false;
        }

        true
    }
}
