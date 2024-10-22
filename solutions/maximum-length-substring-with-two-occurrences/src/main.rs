use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::maximum_length_substring("bcbbbcba".to_string())
    );
    println!("{}", Solution::maximum_length_substring("aaaa".to_string()));
}

struct Solution;

impl Solution {
    pub fn maximum_length_substring(s: String) -> i32 {
        let s = s.into_bytes();
        let n = s.len();

        let mut i = 0;
        let mut mapper = HashMap::new();
        mapper.insert(s[i], 1);

        let mut max_len = 0;

        for (j, &ch) in s.iter().enumerate().take(n).skip(1) {
            mapper
                .entry(ch)
                .and_modify(|occ| *occ += 1)
                .or_insert(1_usize);

            loop {
                let mut is_contains_more_than_two_occurences = false;

                for &v in mapper.values() {
                    if v > 2 {
                        is_contains_more_than_two_occurences = true;
                        break;
                    }
                }

                if !is_contains_more_than_two_occurences {
                    break;
                }

                let ch = s[i];
                if let Some(occ) = mapper.get_mut(&ch) {
                    if *occ == 1 {
                        mapper.remove(&ch);
                    } else {
                        *occ -= 1;
                    }
                }
                i += 1;
            }

            max_len = max_len.max(j - i + 1);
        }

        max_len as i32
    }
}
