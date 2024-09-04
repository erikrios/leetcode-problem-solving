fn main() {
    println!("{}", Solution::chalk_replacer(vec![5, 1, 5], 22));
    println!("{}", Solution::chalk_replacer(vec![3, 4, 1, 2], 25));
    println!("{}", Solution::chalk_replacer(vec![1], 1000000000));
    println!(
        "{}",
        Solution::chalk_replacer(vec![22, 25, 39, 3, 45, 45, 12, 17, 32, 9], 835)
    );
}

struct Solution;

impl Solution {
    pub fn chalk_replacer(chalk: Vec<i32>, k: i32) -> i32 {
        let total: u64 = chalk.iter().map(|x| *x as u64).sum();
        let mut k = k as u64;

        k %= total;

        let mut k = k as i32;

        let mut i = 0;

        while chalk[i] <= k {
            k -= chalk[i];
            i += 1;
        }

        i as i32
    }
}
