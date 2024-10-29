fn main() {
    println!("{}", Solution::largest_odd_number("52".to_string()));
    println!("{}", Solution::largest_odd_number("4206".to_string()));
    println!("{}", Solution::largest_odd_number("35427".to_string()));
    println!(
        "{}",
        Solution::largest_odd_number("239537672423884969653287101".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn largest_odd_number(num: String) -> String {
        let num = num.as_str();
        let n = num.len();

        for i in (0..n).rev() {
            if num.as_bytes()[i] & 1 == 1 {
                return num[0..=i].to_string();
            }
        }

        String::new()
    }
}
