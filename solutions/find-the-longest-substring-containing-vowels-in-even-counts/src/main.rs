fn main() {
    println!(
        "{}",
        Solution::find_the_longest_substring("eleetminicoworoep".to_string())
    );
    println!(
        "{}",
        Solution::find_the_longest_substring("leetcodeisgreat".to_string())
    );
    println!(
        "{}",
        Solution::find_the_longest_substring("bcbcbc".to_string())
    );
    println!(
        "{}",
        Solution::find_the_longest_substring("yopumzgd".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn find_the_longest_substring(s: String) -> i32 {
        let mut max_len = 0;
        let mut bit_mask = 0;
        let mut map = std::collections::HashMap::new();
        map.insert(0, -1);

        for (i, ch) in s.bytes().enumerate() {
            match ch {
                b'a' => bit_mask ^= 1 << 0,
                b'e' => bit_mask ^= 1 << 1,
                b'i' => bit_mask ^= 1 << 2,
                b'o' => bit_mask ^= 1 << 3,
                b'u' => bit_mask ^= 1 << 4,
                _ => {}
            }

            if let Some(prev_index) = map.get(&bit_mask) {
                max_len = max_len.max(i as i32 - prev_index);
            } else {
                map.insert(bit_mask, i as i32);
            }
        }

        max_len
    }
}
