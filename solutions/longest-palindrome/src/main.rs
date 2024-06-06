use std::collections::HashMap;

fn main() {
    println!("{}", Solution::longest_palindrome("abccccdd".to_string()));
    println!("{}", Solution::longest_palindrome("a".to_string()));
}

struct Solution;

impl Solution {
    pub fn longest_palindrome(s: String) -> i32 {
        let s: Vec<char> = s.chars().collect();
        let mut tree_map = HashMap::<char, i32>::new();

        for character in s {
            match tree_map.get_mut(&character) {
                Some(count) => {
                    *count += 1;
                }
                None => {
                    tree_map.insert(character, 1);
                }
            }
        }

        let mut max_odd_key = ' ';

        for (k, count) in &tree_map {
            if count % 2 != 0 {
                if max_odd_key == ' ' {
                    max_odd_key = *k
                } else {
                    if count > tree_map.get(&max_odd_key).unwrap() {
                        max_odd_key = *k
                    }
                }
            }
        }

        let mut counter = 0;
        for (k, count) in tree_map {
            if count % 2 != 0 {
                if k == max_odd_key {
                    counter += count;
                } else {
                    counter += count - 1
                }
            } else {
                counter += count;
            }
        }

        counter
    }
}
