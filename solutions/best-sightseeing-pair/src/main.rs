fn main() {
    println!(
        "{}",
        Solution::max_score_sightseeing_pair(vec![8, 1, 5, 2, 6])
    );
    println!("{}", Solution::max_score_sightseeing_pair(vec![1, 2]));
}

struct Solution;

impl Solution {
    pub fn max_score_sightseeing_pair(values: Vec<i32>) -> i32 {
        let n = values.len();
        let mut max_score = 0;

        for i in 0..n - 1 {
            for j in i + 1..n {
                max_score = max_score.max(values[i] + values[j] + i as i32 - j as i32);
            }
        }

        max_score
    }
}
