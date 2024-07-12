fn main() {
    println!("{}", Solution::judge_square_sum(5));
    println!("{}", Solution::judge_square_sum(3));
    println!("{}", Solution::judge_square_sum(4));
    println!("{}", Solution::judge_square_sum(50));
    println!("{}", Solution::judge_square_sum(0));
    println!("{}", Solution::judge_square_sum(11));
}

struct Solution;

impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        let mut i = 0;
        loop {
            let remaining = c - i * i;
            if remaining < 0 {
                break;
            }

            if !Self::has_fractional_part(remaining as f64) {
                return true;
            }

            i += 1;
        }

        false
    }

    fn has_fractional_part(n: f64) -> bool {
        let sqrt_n = n.sqrt();
        sqrt_n.fract() != 0.0
    }
}
