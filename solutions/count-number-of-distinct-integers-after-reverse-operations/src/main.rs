use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::count_distinct_integers(vec![1, 13, 10, 12, 31])
    );
    println!("{}", Solution::count_distinct_integers(vec![2, 2, 2]));
}

struct Solution;

impl Solution {
    pub fn count_distinct_integers(nums: Vec<i32>) -> i32 {
        let mut set = HashSet::<i32>::new();

        for num in nums {
            set.insert(num);
            set.insert(Solution::reverse(num));
        }

        set.len() as i32
    }

    fn reverse(mut num: i32) -> i32 {
        let mut res = 0;

        while num > 0 {
            let digit = num % 10;
            res = res * 10 + digit;
            num /= 10;
        }

        res
    }
}
