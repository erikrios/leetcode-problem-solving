fn main() {
    println!(
        "{:?}",
        Solution::word_subsets(
            vec![
                "amazon".to_string(),
                "apple".to_string(),
                "facebook".to_string(),
                "google".to_string(),
                "leetcode".to_string()
            ],
            vec!["e".to_string(), "o".to_string()]
        )
    );
    println!(
        "{:?}",
        Solution::word_subsets(
            vec![
                "amazon".to_string(),
                "apple".to_string(),
                "facebook".to_string(),
                "google".to_string(),
                "leetcode".to_string()
            ],
            vec!["l".to_string(), "e".to_string()]
        )
    );
    println!(
        "{:?}",
        Solution::word_subsets(
            vec![
                "amazon".to_string(),
                "apple".to_string(),
                "facebook".to_string(),
                "google".to_string(),
                "leetcode".to_string()
            ],
            vec!["e".to_string(), "oo".to_string()]
        )
    );
    println!(
        "{:?}",
        Solution::word_subsets(
            vec![
                "amazon".to_string(),
                "apple".to_string(),
                "facebook".to_string(),
                "google".to_string(),
                "leetcode".to_string()
            ],
            vec!["lo".to_string(), "eo".to_string()]
        )
    );
}

struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn word_subsets(words1: Vec<String>, words2: Vec<String>) -> Vec<String> {
        let mut universal_chars = HashMap::new();

        for word in words2 {
            let mut mapper = HashMap::new();
            for ch in word.into_bytes() {
                mapper.entry(ch).and_modify(|e| *e += 1).or_insert(1_usize);
            }
            for (ch_c, count) in mapper {
                if !universal_chars.contains_key(&ch_c)
                    || count > *universal_chars.get(&ch_c).unwrap()
                {
                    universal_chars.insert(ch_c, count);
                }
            }
        }

        let mut words = Vec::new();
        'outer: for word in words1 {
            let mut ch_counters = HashMap::new();
            for &ch in word.as_bytes() {
                ch_counters
                    .entry(ch)
                    .and_modify(|e| *e += 1)
                    .or_insert(1_usize);
            }

            for (ch, &count) in &universal_chars {
                if let Some(&ch_c) = ch_counters.get(ch) {
                    if ch_c < count {
                        continue 'outer;
                    }
                } else {
                    continue 'outer;
                }
            }

            words.push(word);
        }

        words
    }
}
