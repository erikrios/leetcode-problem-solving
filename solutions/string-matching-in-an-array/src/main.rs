use std::collections::HashSet;

fn main() {
    println!("{:?}", Solution::string_matching(vec!["mass".to_string(),"as".to_string(),"hero".to_string(),"superhero".to_string()]));
    println!("{:?}", Solution::string_matching(vec!["leetcode".to_string(),"et".to_string(),"code".to_string()]));
    println!("{:?}", Solution::string_matching(vec!["blue".to_string(),"green".to_string(),"bu".to_string()]));
    println!("{:?}", Solution::string_matching(vec!["leetcoder".to_string(),"leetcode".to_string(),"od".to_string(),"hamlet".to_string(),"am".to_string()]));
}

struct Solution;

impl Solution {
    pub fn string_matching(words: Vec<String>) -> Vec<String> {
        let mut results: HashSet<String> = HashSet::new();

        for i in 0..words.len() {
            let word1 = &words[i];
            for j in 0..words.len() {
                if i == j {
                    continue;
                }
                let word2 = &words[j];
                if word2.contains(word1.as_str()) {
                    results.insert(word1.clone());
                }
            }
        }

        results.into_iter().collect()
    }
}
