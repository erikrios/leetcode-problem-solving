fn main() {
    println!(
        "{}",
        Solution::add_spaces("LeetcodeHelpsMeLearn".to_string(), vec![8, 13, 15])
    );
    println!(
        "{}",
        Solution::add_spaces("icodeinpython".to_string(), vec![1, 5, 7, 9])
    );
    println!(
        "{}",
        Solution::add_spaces("spacing".to_string(), vec![0, 1, 2, 3, 4, 5, 6])
    );
}

struct Solution;

impl Solution {
    pub fn add_spaces(s: String, spaces: Vec<i32>) -> String {
        let mut results = String::with_capacity(s.len() + spaces.len());
        let s = s.into_bytes();
        let mut j = 0;

        for (i, ch) in s.into_iter().enumerate() {
            if j < spaces.len() && spaces[j] as usize == i {
                results.push(' ');
                j += 1;
            }
            results.push(ch as char);
        }

        results
    }
}
