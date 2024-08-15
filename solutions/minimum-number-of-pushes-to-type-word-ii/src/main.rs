use std::collections::HashMap;

fn main() {
    println!("{}", Solution::minimum_pushes("abcde".to_string()));
    println!("{}", Solution::minimum_pushes("xyzxyzxyzxyz".to_string()));
    println!(
        "{}",
        Solution::minimum_pushes("aabbccddeeffgghhiiiiii".to_string())
    );
}

struct Solution;

#[derive(Clone, Copy, Debug)]
struct CharCount {
    ch: u8,
    total: usize,
}

impl Solution {
    pub fn minimum_pushes(word: String) -> i32 {
        let mut alphabet_freq_counters = [CharCount { ch: b'a', total: 0 }; 26];
        for (i, char_counter) in alphabet_freq_counters.iter_mut().enumerate() {
            char_counter.ch = b'a' + (i as u8);
        }

        for ch in word.chars() {
            let index = ((ch as u8) - (b'a')) as usize;
            alphabet_freq_counters[index].total += 1;
        }

        alphabet_freq_counters.sort_by(|a, b| b.total.cmp(&a.total));

        let mut alphabet_index_mapper = HashMap::new();

        for (i, char_counter) in alphabet_freq_counters.iter().enumerate() {
            alphabet_index_mapper.insert(char_counter.ch, i);
        }

        let mut total = 0;
        const TOTAL_BUTTON: usize = 8;
        for ch in word.chars() {
            let index = alphabet_index_mapper.get(&(ch as u8)).unwrap();
            total += (index / TOTAL_BUTTON) + 1;
        }

        total as i32
    }
}
