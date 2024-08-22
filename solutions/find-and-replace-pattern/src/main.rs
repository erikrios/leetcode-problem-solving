fn main() {
    println!(
        "{:?}",
        Solution::find_and_replace_pattern(
            vec![
                "abc".to_string(),
                "deq".to_string(),
                "mee".to_string(),
                "aqq".to_string(),
                "dkd".to_string(),
                "ccc".to_string(),
            ],
            "abb".to_string()
        )
    );
    println!(
        "{:?}",
        Solution::find_and_replace_pattern(
            vec!["a".to_string(), "b".to_string(), "c".to_string()],
            "a".to_string()
        )
    );
}

struct Solution;

impl Solution {
    pub fn find_and_replace_pattern(words: Vec<String>, pattern: String) -> Vec<String> {
        let pattern: Vec<char> = pattern.chars().collect();
        let mut results = Vec::new();

        'outer: for word in words {
            let mut mapper_one = std::collections::HashMap::new();
            let mut mapper_two = std::collections::HashMap::new();
            for (j, ch) in word.chars().enumerate() {
                let val = pattern[j];
                mapper_one.insert(ch, val);
                mapper_two.insert(val, ch);
            }

            for (j, ch) in word.chars().enumerate() {
                let &val_one = mapper_one.get(&ch).unwrap();
                let expected_one = pattern[j];

                let &val_two = mapper_two.get(&pattern[j]).unwrap();
                let expected_two = ch;
                if val_one != expected_one || val_two != expected_two {
                    continue 'outer;
                }
            }

            results.push(word);
        }

        results
    }
}
