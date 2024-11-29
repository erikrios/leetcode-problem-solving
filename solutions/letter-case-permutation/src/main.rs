fn main() {
    println!(
        "{:?}",
        Solution::letter_case_permutation("a1b2".to_string())
    );
    println!("{:?}", Solution::letter_case_permutation("3z4".to_string()));
    println!("{:?}", Solution::letter_case_permutation("C".to_string()));
}

struct Solution;

impl Solution {
    pub fn letter_case_permutation(s: String) -> Vec<String> {
        let mut results = Vec::new();
        let mut s = s.into_bytes();

        Self::backtracking(&mut s, 0, &mut results);

        results
    }

    fn backtracking(s: &mut [u8], i: usize, results: &mut Vec<String>) {
        if i == s.len() {
            results.push(String::from_utf8(s.to_vec()).unwrap());
            return;
        }

        let ch = s[i];

        if !ch.is_ascii_lowercase() && !ch.is_ascii_uppercase() {
            Self::backtracking(s, i + 1, results);
        } else {
            if ch.is_ascii_lowercase() {
                s[i] = b'A' + (ch - b'a');
            } else {
                s[i] = b'a' + (ch - b'A');
            }

            Self::backtracking(s, i + 1, results);
            s[i] = ch;
            Self::backtracking(s, i + 1, results);
        }
    }
}
