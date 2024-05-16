fn main() {
    println!(
        "{}",
        Solution::count_prefix_suffix_pairs(vec![
            "a".to_string(),
            "aba".to_string(),
            "ababa".to_string(),
            "aa".to_string()
        ])
    );
    println!(
        "{}",
        Solution::count_prefix_suffix_pairs(vec![
            "pa".to_string(),
            "papa".to_string(),
            "ma".to_string(),
            "mama".to_string()
        ])
    );
    println!(
        "{}",
        Solution::count_prefix_suffix_pairs(vec!["abab".to_string(), "ab".to_string()])
    );
}

struct Solution;

impl Solution {
    pub fn count_prefix_suffix_pairs(words: Vec<String>) -> i32 {
        let n = words.len();
        let mut counter = 0;

        for i in 0..n - 1 {
            for j in i + 1..n {
                if Solution::is_prefix_and_suffix(&words[i], &words[j]) {
                    counter += 1;
                }
            }
        }

        counter
    }

    fn is_prefix_and_suffix(word1: &String, word2: &String) -> bool {
        let n = word1.len();
        let m = word2.len();

        if n > m {
            return false;
        }

        for i in 0..n {
            let character_one_front = word1.as_bytes()[i];
            let character_two_front = word2.as_bytes()[i];
            if character_one_front != character_two_front {
                return false;
            }

            let character_one_back = word1.as_bytes()[n - 1 - i];
            let character_two_back = word2.as_bytes()[m - 1 - i];
            if character_one_back != character_two_back {
                return false;
            }
        }

        true
    }
}
