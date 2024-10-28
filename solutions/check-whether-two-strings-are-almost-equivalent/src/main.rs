fn main() {
    println!(
        "{}",
        Solution::check_almost_equivalent("aaaa".to_string(), "bccb".to_string())
    );
    println!(
        "{}",
        Solution::check_almost_equivalent("abcdeef".to_string(), "abaaacc".to_string())
    );
    println!(
        "{}",
        Solution::check_almost_equivalent("cccddabba".to_string(), "babababab".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn check_almost_equivalent(word1: String, word2: String) -> bool {
        let n = word1.len();
        let word1 = word1.into_bytes();
        let word2 = word2.into_bytes();
        let mut freqs = [0; 26];

        for i in 0..n {
            let ch_1 = word1[i];
            let ch_2 = word2[i];

            freqs[(ch_1 - b'a') as usize] += 1;
            freqs[(ch_2 - b'a') as usize] -= 1;
        }

        for freq in freqs {
            if !(-3..=3).contains(&freq) {
                return false;
            }
        }

        true
    }
}
