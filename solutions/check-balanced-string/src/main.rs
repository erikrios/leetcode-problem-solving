fn main() {
    println!("{}", Solution::is_balanced("1234".to_string()));
    println!("{}", Solution::is_balanced("24123".to_string()));
}

struct Solution;

impl Solution {
    pub fn is_balanced(num: String) -> bool {
        let mut even = 0;
        let mut odd = 0;

        for (i, ch) in num.into_bytes().into_iter().enumerate() {
            let ch = ch - b'0';
            if i & 1 == 0 {
                even += ch as usize;
            } else {
                odd += ch as usize;
            }
        }

        even == odd
    }
}
