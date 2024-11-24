fn main() {
    println!("{:?}", Solution::missing_rolls(vec![3, 2, 4, 3], 4, 2));
    println!("{:?}", Solution::missing_rolls(vec![1, 5, 6], 3, 4));
    println!("{:?}", Solution::missing_rolls(vec![1, 2, 3, 4], 6, 4));
    println!(
        "{:?}",
        Solution::missing_rolls(vec![6, 3, 4, 3, 5, 3], 1, 6)
    );
    println!("{:?}", Solution::missing_rolls(vec![6, 1, 5, 2], 4, 4));
    println!(
        "{:?}",
        Solution::missing_rolls(vec![1, 3, 6, 4, 1, 2, 1, 5, 5, 4], 6, 10)
    );
    println!(
        "{:?}",
        Solution::missing_rolls(
            vec![
                4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3,
                5, 1, 5
            ],
            4,
            40
        )
    );
    println!(
        "{:?}",
        Solution::missing_rolls(
            vec![
                4, 2, 2, 5, 4, 5, 4, 5, 3, 3, 6, 1, 2, 4, 2, 1, 6, 5, 4, 2, 3, 4, 2, 3, 3, 5, 4, 1,
                4, 4, 5, 3, 6, 1, 5, 2, 3, 3, 6, 1, 6, 4, 1, 3
            ],
            2,
            53
        )
    );
}

struct Solution;

impl Solution {
    pub fn missing_rolls(rolls: Vec<i32>, mean: i32, n: i32) -> Vec<i32> {
        let rolls_sum: i32 = rolls.iter().sum();
        let total = (rolls.len() as i32 + n) * mean;
        let remaining = total - rolls_sum;

        if rolls_sum > total {
            return vec![];
        }

        let num = remaining / n;
        let sub = remaining - (num * n);
        if remaining < n || num > 6 || (num == 6 && sub > 0) {
            return vec![];
        }

        let mut results = vec![num; n as usize];
        for i in 0..sub {
            results[i as usize] += 1;
        }

        results
    }
}
