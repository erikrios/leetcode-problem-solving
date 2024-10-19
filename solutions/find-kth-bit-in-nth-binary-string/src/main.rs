fn main() {
    println!("{}", Solution::find_kth_bit(3, 1));
    println!("{}", Solution::find_kth_bit(4, 11));
}

struct Solution;

impl Solution {
    pub fn find_kth_bit(n: i32, k: i32) -> char {
        let mut s = "0".to_string();

        for _ in 2..=n {
            let si = s
                .chars()
                .map(|x| if x == '0' { '1' } else { '0' })
                .rev()
                .collect::<String>();

            s.push('1');
            s.push_str(&si);
        }

        s.chars().nth(k as usize - 1).unwrap()
    }
}
