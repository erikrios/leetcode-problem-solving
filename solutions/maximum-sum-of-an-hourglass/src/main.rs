fn main() {
    println!(
        "{}",
        Solution::max_sum(vec![
            vec![6, 2, 1, 3],
            vec![4, 2, 1, 5],
            vec![9, 2, 8, 7],
            vec![4, 1, 2, 9]
        ])
    );
    println!(
        "{}",
        Solution::max_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]])
    );
}

struct Solution;

const DIRECTIONS: [(i8, i8); 7] = [(0, 0), (-1, -1), (-1, 0), (-1, 1), (1, -1), (1, 0), (1, 1)];

impl Solution {
    pub fn max_sum(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut max_sum = 0;

        for i in 1..m - 1 {
            for j in 1..n - 1 {
                let mut sum = 0;

                for dir in DIRECTIONS {
                    sum += grid[(i as i32 + dir.0 as i32) as usize]
                        [(j as i32 + dir.1 as i32) as usize];
                }

                max_sum = max_sum.max(sum);
            }
        }

        max_sum
    }
}
