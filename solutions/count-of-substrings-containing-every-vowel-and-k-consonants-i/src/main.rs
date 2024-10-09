use std::collections::HashSet;

fn main() {
    println!("{}", Solution::count_of_substrings("aeioqq".to_string(), 1));
    println!("{}", Solution::count_of_substrings("aeiou".to_string(), 0));
    println!(
        "{}",
        Solution::count_of_substrings("ieaouqqieaouqq".to_string(), 1)
    );
    println!(
        "{}",
        Solution::count_of_substrings("iqeaouqi".to_string(), 2)
    );
    println!(
        "{}",
        Solution::count_of_substrings("auieoui".to_string(), 0)
    );
}

struct Solution;

impl Solution {
    pub fn count_of_substrings(word: String, k: i32) -> i32 {
        let mut counter = 0;
        let min_word_len = 5 + k as usize;

        let mut vowels = HashSet::new();
        vowels.insert(b'a');
        vowels.insert(b'i');
        vowels.insert(b'u');
        vowels.insert(b'e');
        vowels.insert(b'o');

        for i in 0..word.len() - min_word_len + 1 {
            let mut vowels_tracker = HashSet::new();
            let mut consonants_counter = 0;
            for j in i..word.len() {
                let ch = word.as_bytes()[j];
                if vowels.contains(&ch) {
                    vowels_tracker.insert(ch);
                } else {
                    consonants_counter += 1;
                }

                if j - i >= min_word_len - 1 {
                    if vowels_tracker.len() == 5 && consonants_counter == k {
                        counter += 1;
                    } else if consonants_counter > k {
                        break;
                    }
                }
            }
        }

        counter
    }
}
