fn main() {
    println!("{}", Solution::clear_digits("abc".to_string()));
    println!("{}", Solution::clear_digits("cb34".to_string()));
    println!("{}", Solution::clear_digits("g0".to_string()));
}

struct Solution;

impl Solution {
    pub fn clear_digits(s: String) -> String {
        let mut res = String::new();

        let s = s.chars();

        for character in s {
            if character as u8 >= b'0' && character as u8 <= b'9' {
                if !res.is_empty() {
                    res.remove(res.len() - 1);
                }
            } else {
                res.push(character);
            }
        }

        res
    }
}
