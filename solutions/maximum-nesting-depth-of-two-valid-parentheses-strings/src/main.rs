fn main() {
    println!(
        "{:?}",
        Solution::max_depth_after_split("(()())".to_string())
    );
    println!(
        "{:?}",
        Solution::max_depth_after_split("()(())()".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn max_depth_after_split(seq: String) -> Vec<i32> {
        let n = seq.len();
        let mut depth = 0;
        let mut results = Vec::with_capacity(n);

        for ch in seq.into_bytes() {
            if ch == b'(' {
                depth += 1;
            }

            if depth & 1 == 1 {
                results.push(0);
            } else {
                results.push(1);
            }

            if ch == b')' {
                depth -= 1;
            }
        }

        results
    }
}
