fn main() {
    println!(
        "{}",
        Solution::min_flips(vec![vec![1, 0, 0], vec![0, 0, 0], vec![0, 0, 1]])
    );
    println!(
        "{}",
        Solution::min_flips(vec![vec![0, 1], vec![0, 1], vec![0, 0]])
    );
    println!("{}", Solution::min_flips(vec![vec![1], vec![0]]));
}

struct Solution;

impl Solution {
    pub fn min_flips(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();

        let mut row_flip = 0;
        for nums in grid.iter() {
            for j in 0..n / 2 {
                if nums[j] != nums[n - 1 - j] {
                    row_flip += 1;
                }
            }
        }

        let mut col_flip = 0;
        for j in 0..n {
            for i in 0..m / 2 {
                if grid[i][j] != grid[m - 1 - i][j] {
                    col_flip += 1;
                }
            }
        }

        if col_flip < row_flip {
            return col_flip;
        }

        row_flip
    }
}
