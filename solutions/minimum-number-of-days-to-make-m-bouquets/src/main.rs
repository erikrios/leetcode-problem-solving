fn main() {
    println!("{}", Solution::min_days(vec![1, 10, 3, 10, 2], 3, 1));
    println!("{}", Solution::min_days(vec![1, 10, 3, 10, 2], 3, 2));
    println!("{}", Solution::min_days(vec![7, 7, 7, 7, 12, 7, 7], 2, 3));
}

struct Solution;

impl Solution {
    pub fn min_days(bloom_day: Vec<i32>, m: i32, k: i32) -> i32 {
        let mut left = i32::MAX;
        let mut right = i32::MIN;

        for &day in &bloom_day {
            if day < left {
                left = day;
            }

            if day > right {
                right = day;
            }
        }

        let mut min_days = -1;
        while left <= right {
            let mid = (left + right) / 2;

            let count = Solution::count_days(&bloom_day, mid, k);

            if count >= m {
                right = mid - 1;
                min_days = mid;
            } else {
                left = mid + 1;
            }
        }

        min_days
    }

    fn count_days(bloom_day: &Vec<i32>, mid: i32, k: i32) -> i32 {
        let mut counter = 0;
        let mut days = 0;

        for &day in bloom_day {
            if mid >= day {
                counter += 1;
            } else {
                counter = 0;
            }

            if counter == k {
                days += 1;
                counter = 0;
            }
        }

        days
    }
}
