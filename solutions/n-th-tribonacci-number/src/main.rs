use std::collections::HashMap;

fn main() {
    println!("{}", Solution::tribonacci(4));
    println!("{}", Solution::tribonacci(25));
    println!("{}", Solution::tribonacci(37));
}

struct Solution;

impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        let mut memo = HashMap::new();
        Self::tribonacci_with_memo(n, &mut memo)
    }

    fn tribonacci_with_memo(n: i32, memo: &mut HashMap<i32, i32>) -> i32 {
        if n == 0 {
            return 0;
        } else if n == 1 || n == 2 {
            return 1;
        }

        if let Some(&val) = memo.get(&n) {
            return val;
        }

        let res = Self::tribonacci_with_memo(n - 1, memo)
            + Self::tribonacci_with_memo(n - 2, memo)
            + Self::tribonacci_with_memo(n - 3, memo);
        memo.insert(n, res);
        res
    }
}
