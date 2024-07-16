fn main() {
    println!(
        "{}",
        Solution::maximum_gain("cdbcbbaaabab".to_string(), 4, 5)
    );
    println!(
        "{}",
        Solution::maximum_gain("aabbaaxybbaabb".to_string(), 5, 4)
    );
}

struct Solution;

impl Solution {
    pub fn maximum_gain(s: String, x: i32, y: i32) -> i32 {
        const X_SUBSTR: &str = "ab";
        const Y_SUBSTR: &str = "ba";

        let (max, max_str, min, min_str): (i32, Vec<char>, i32, Vec<char>) = if x > y {
            (x, X_SUBSTR.chars().collect(), y, Y_SUBSTR.chars().collect())
        } else {
            (y, Y_SUBSTR.chars().collect(), x, X_SUBSTR.chars().collect())
        };

        let mut chars: Vec<char> = s.chars().collect();

        let mut score = 0;

        let mut i = 0;
        while i < chars.len() as i32 - 1 {
            let cur = chars[i as usize];
            let next = chars[i as usize + 1];

            if cur == max_str[0] && next == max_str[1] {
                chars.drain(i as usize..=i as usize + 1);
                score += max;
                i = i.saturating_sub(1).max(0);
            } else {
                i += 1;
            }
        }

        i = 0;
        while i < chars.len() as i32 - 1 {
            let cur = chars[i as usize];
            let next = chars[i as usize + 1];

            if cur == min_str[0] && next == min_str[1] {
                chars.drain(i as usize..=i as usize + 1);
                score += min;
                i = i.saturating_sub(1).max(0);
            } else {
                i += 1;
            }
        }

        score
    }
}
