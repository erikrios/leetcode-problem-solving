fn main() {
    println!("{}", Solution::title_to_number("A".to_string()));
    println!("{}", Solution::title_to_number("AB".to_string()));
    println!("{}", Solution::title_to_number("ZY".to_string()));
    println!("{}", Solution::title_to_number("AAA".to_string()));
}

struct Solution;

impl Solution {
    pub fn title_to_number(column_title: String) -> i32 {
        const BASE: i32 = 26;

        let mut res = 0;
        for (i, ch) in column_title.chars().rev().enumerate() {
            let multiplier = (ch as u8) - b'A' + 1;
            res += BASE.pow(i as u32) * multiplier as i32
        }

        res
    }
}
