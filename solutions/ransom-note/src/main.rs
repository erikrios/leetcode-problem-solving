fn main() {
    println!(
        "{}",
        Solution::can_construct("a".to_string(), "b".to_string())
    );
    println!(
        "{}",
        Solution::can_construct("aa".to_string(), "ab".to_string())
    );
    println!(
        "{}",
        Solution::can_construct("aa".to_string(), "aab".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut letters = [0_usize; 26];

        for ch in magazine.into_bytes() {
            letters[(ch - b'a') as usize] += 1;
        }

        for ch in ransom_note.into_bytes() {
            let letter_counter = letters[(ch - b'a') as usize];
            if letter_counter == 0 {
                return false;
            }

            letters[(ch - b'a') as usize] -= 1;
        }

        true
    }
}
