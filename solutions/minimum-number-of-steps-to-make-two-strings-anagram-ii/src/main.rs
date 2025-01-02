fn main() {
    println!(
        "{}",
        Solution::min_steps("leetcode".to_string(), "coats".to_string())
    );
    println!(
        "{}",
        Solution::min_steps("night".to_string(), "thing".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn min_steps(s: String, t: String) -> i32 {
        let mut s =
            s.into_bytes()
                .into_iter()
                .fold(std::collections::HashMap::new(), |mut map, b| {
                    map.entry(b).and_modify(|count| *count += 1).or_insert(1);
                    map
                });

        let mut t =
            t.into_bytes()
                .into_iter()
                .fold(std::collections::HashMap::new(), |mut map, b| {
                    map.entry(b).and_modify(|count| *count += 1).or_insert(1);
                    map
                });

        let s_entries: Vec<(u8, usize)> = s.iter().map(|(&k, &v)| (k, v)).collect();

        for (ch_s, count_s) in s_entries {
            if let Some(&count_t) = t.get(&ch_s) {
                match count_s.cmp(&count_t) {
                    std::cmp::Ordering::Equal => {
                        s.remove(&ch_s);
                        t.remove(&ch_s);
                    }
                    std::cmp::Ordering::Less => {
                        s.remove(&ch_s);
                        t.insert(ch_s, count_t - count_s);
                    }
                    std::cmp::Ordering::Greater => {
                        t.remove(&ch_s);
                        s.insert(ch_s, count_s - count_t);
                    }
                }
            }
        }

        (t.values().sum::<usize>() + s.values().sum::<usize>()) as i32
    }
}
