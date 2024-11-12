fn main() {
    println!(
        "{}",
        Solution::largest_combination(vec![16, 17, 71, 62, 12, 24, 14])
    );
    println!("{}", Solution::largest_combination(vec![8, 8]));
}

struct Solution;

impl Solution {
    pub fn largest_combination(candidates: Vec<i32>) -> i32 {
        let mut counts = [0_usize; 32];

        for candidate in candidates {
            let mut candidate = candidate;
            let mut i = 0;

            while candidate > 0 {
                counts[i] += (candidate & 1) as usize;
                candidate >>= 1;
                i += 1;
            }
        }

        counts.into_iter().max().unwrap() as i32
    }
}
