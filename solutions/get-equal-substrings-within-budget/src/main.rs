fn main() {
    println!(
        "{}",
        Solution::equal_substring(String::from("abcd"), String::from("bcdf"), 3)
    );
    println!(
        "{}",
        Solution::equal_substring(String::from("abcd"), String::from("cdef"), 3)
    );
    println!(
        "{}",
        Solution::equal_substring(String::from("abcd"), String::from("acde"), 0)
    );
    println!(
        "{}",
        Solution::equal_substring(String::from("krrgw"), String::from("zjxss"), 19)
    );
    println!(
        "{}",
        Solution::equal_substring(String::from("pxezla"), String::from("loewbi"), 25)
    );
}

struct Solution;

impl Solution {
    pub fn equal_substring(s: String, t: String, max_cost: i32) -> i32 {
        let n = s.len();

        let s_bytes = s.as_bytes();
        let t_bytes = t.as_bytes();

        let mut max_cost = max_cost;
        let mut i = 0;

        for j in 0..n {
            let s_char = s_bytes[j];
            let t_char = t_bytes[j];

            let abs = (s_char as i32 - t_char as i32).abs();

            max_cost -= abs;

            if max_cost < 0 {
                max_cost += (s_bytes[i] as i32 - t_bytes[i] as i32).abs();
                i += 1;
            }
        }

        (n - i) as i32
    }
}
