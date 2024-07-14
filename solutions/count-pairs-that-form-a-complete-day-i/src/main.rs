fn main() {
    println!(
        "{}",
        Solution::count_complete_day_pairs(vec![12, 12, 30, 24, 24])
    );
    println!(
        "{}",
        Solution::count_complete_day_pairs(vec![72, 48, 24, 3])
    );
}

struct Solution;

impl Solution {
    pub fn count_complete_day_pairs(hours: Vec<i32>) -> i32 {
        let n = hours.len();
        let mut counter = 0;

        for i in 0..n - 1 {
            for j in i + 1..n {
                if (hours[i] + hours[j]) % 24 == 0 {
                    counter += 1;
                }
            }
        }

        counter
    }
}
