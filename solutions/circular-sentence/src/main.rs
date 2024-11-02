fn main() {
    println!(
        "{}",
        Solution::is_circular_sentence("leetcode exercises sound delightful".to_string())
    );
    println!("{}", Solution::is_circular_sentence("eetcode".to_string()));
    println!(
        "{}",
        Solution::is_circular_sentence("Leetcode is cool".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn is_circular_sentence(sentence: String) -> bool {
        let sentence = sentence.into_bytes();
        let n = sentence.len();

        for (i, &ch) in sentence.iter().enumerate() {
            if i == n - 1 {
                return ch == sentence[0];
            }

            let next_ch = sentence[i + 1];
            if next_ch == b' ' {
                let next_ch = sentence[i + 2];
                if ch != next_ch {
                    return false;
                }
            }
        }

        true
    }
}
