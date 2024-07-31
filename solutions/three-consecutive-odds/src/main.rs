fn main() {
    println!("{}", Solution::three_consecutive_odds(vec![2, 6, 4, 1]));
    println!(
        "{}",
        Solution::three_consecutive_odds(vec![1, 2, 34, 3, 4, 5, 7, 23, 12])
    );
}

struct Solution;

impl Solution {
    pub fn three_consecutive_odds(arr: Vec<i32>) -> bool {
        let mut odd_counter = 0;

        for num in arr {
            if num % 2 == 0 {
                odd_counter = 0;
            } else {
                odd_counter += 1;
                if odd_counter == 3 {
                    return true;
                }
            }
        }

        false
    }
}
