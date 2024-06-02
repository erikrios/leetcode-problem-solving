fn main() {
    println!(
        "{}",
        Solution::are_almost_equal(String::from("bank"), String::from("kanb"))
    );
    println!(
        "{}",
        Solution::are_almost_equal(String::from("attack"), String::from("defend"))
    );
    println!(
        "{}",
        Solution::are_almost_equal(String::from("kelb"), String::from("kelb"))
    );
}

struct Solution;

impl Solution {
    pub fn are_almost_equal(s1: String, s2: String) -> bool {
        let n = s1.len();
        let mut diff_count = 0;
        let mut diffs = [0; 2];

        for i in 0..n {
            let character_a = s1.chars().nth(i).unwrap();
            let character_b = s2.chars().nth(i).unwrap();

            if character_a != character_b {
                diff_count += 1;
                if diff_count > 2 {
                    return false;
                }
                diffs[diff_count - 1] = i;
            }
        }

        if diff_count == 0 {
            return true;
        }

        let mut chars: Vec<char> = s2.chars().collect();
        chars.swap(diffs[0], diffs[1]);
        let s2: String = chars.into_iter().collect();

        if s1 != s2 {
            return false;
        }

        true
    }
}
