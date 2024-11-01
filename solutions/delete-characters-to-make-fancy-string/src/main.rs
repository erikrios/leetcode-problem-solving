fn main() {
    println!("{}", Solution::make_fancy_string("leeetcode".to_string()));
    println!("{}", Solution::make_fancy_string("aaabaaaa".to_string()));
    println!("{}", Solution::make_fancy_string("aab".to_string()));
}

struct Solution;

impl Solution {
    pub fn make_fancy_string(s: String) -> String {
        let mut results = String::new();
        let mut prev = b' ';
        let mut consecutive_counter = 0;

        for ch in s.into_bytes() {
            if ch == prev {
                consecutive_counter += 1;
            } else {
                consecutive_counter = 1;
            }

            if consecutive_counter <= 2 {
                results.push(ch as char);
            }

            prev = ch;
        }

        results
    }
}
