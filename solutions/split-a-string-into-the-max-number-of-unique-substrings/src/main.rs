use std::collections::HashSet;

fn main() {
    println!("{}", Solution::max_unique_split("abad".to_string()));
    println!("{}", Solution::max_unique_split("ababccc".to_string()));
    println!("{}", Solution::max_unique_split("aba".to_string()));
    println!("{}", Solution::max_unique_split("aa".to_string()));
    println!(
        "{}",
        Solution::max_unique_split("wwwzfvedwfvhsww".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn max_unique_split(s: String) -> i32 {
        Self::backtrack(s.as_str(), &mut HashSet::new()) as i32
    }

    fn backtrack<'a>(s: &'a str, hash_set: &mut HashSet<&'a str>) -> usize {
        if s.is_empty() {
            return hash_set.len();
        }

        let n = s.len();

        let mut max_len = 0;

        for i in 0..n {
            let pre = &s[..i + 1];

            if !hash_set.contains(pre) {
                hash_set.insert(pre);
                let post = &s[i + 1..];
                let len = Self::backtrack(post, hash_set);
                max_len = max_len.max(len);
                hash_set.remove(pre);
            }
        }

        max_len
    }
}
