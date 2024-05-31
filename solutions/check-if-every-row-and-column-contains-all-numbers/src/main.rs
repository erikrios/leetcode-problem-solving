use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::check_valid(vec![vec![1, 2, 3], vec![3, 1, 2], vec![2, 3, 1]])
    );
    println!(
        "{}",
        Solution::check_valid(vec![vec![1, 1, 1], vec![1, 2, 3], vec![1, 2, 3]])
    );
}

struct Solution;

impl Solution {
    pub fn check_valid(matrix: Vec<Vec<i32>>) -> bool {
        let n = matrix.len();

        for i in 0..n {
            let mut set_row = HashSet::<i32>::new();
            let mut set_col = HashSet::<i32>::new();
            for j in 0..n {
                set_row.insert(matrix[i][j]);
                set_col.insert(matrix[j][i]);
            }
            if set_row.len() != n {
                return false;
            }
            if set_col.len() != n {
                return false;
            }
        }

        true
    }
}
