use std::cmp::min;

fn main() {
    println!("{}", Solution::max_score(vec![1, 2, 3, 4, 5, 6, 1], 3));
    println!("{}", Solution::max_score(vec![2, 2, 2], 2));
    println!("{}", Solution::max_score(vec![9, 7, 7, 9, 7, 7, 9], 7));
    println!("{}", Solution::max_score(vec![9, 7, 7, 9, 7, 7, 10], 1));
}

struct Solution;

impl Solution {
    pub fn max_score(card_points: Vec<i32>, k: i32) -> i32 {
        let n = card_points.len();
        let window_range = n - k as usize;
        let mut window_start = 0;
        let mut min_sum = i32::MAX;
        let mut current_sum = 0;
        let mut sum = 0;

        for (window_end, card_point) in card_points.iter().enumerate() {
            current_sum += card_point;
            sum += card_point;

            if window_range == 0 {
                continue;
            }

            if window_end >= window_range - 1 {
                min_sum = min(min_sum, current_sum);
                current_sum -= card_points[window_start];
                window_start += 1;
            }
        }

        if min_sum == i32::MAX {
            min_sum = 0;
        }

        sum - min_sum
    }
}
