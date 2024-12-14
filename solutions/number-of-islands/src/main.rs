fn main() {
    println!(
        "{}",
        Solution::num_islands(vec![
            vec!['1', '1', '1', '1', '0'],
            vec!['1', '1', '0', '1', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '0', '0', '0'],
        ],)
    );
    println!(
        "{}",
        Solution::num_islands(vec![
            vec!['1', '1', '0', '0', '0'],
            vec!['1', '1', '0', '0', '0'],
            vec!['0', '0', '1', '0', '0'],
            vec!['0', '0', '0', '1', '1']
        ],)
    );
    println!(
        "{}",
        Solution::num_islands(vec![
            vec!['1', '1', '1'],
            vec!['0', '1', '0'],
            vec!['1', '1', '1'],
        ],)
    );
    println!(
        "{}",
        Solution::num_islands(vec![
            vec!['1', '0', '1', '1', '1'],
            vec!['1', '0', '1', '0', '1'],
            vec!['1', '1', '1', '0', '1'],
        ],)
    );
}

struct Solution;

impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let mut grid = grid;
        let m = grid.len();
        let n = grid[0].len();

        let mut counter = 0;
        for i in 0..m {
            for j in 0..n {
                if grid[i][j] == '1' {
                    Self::dfs(&mut grid, i, j);
                    counter += 1
                }
            }
        }

        counter
    }

    fn is_inside(m: usize, n: usize, i: i32, j: i32) -> bool {
        i >= 0 && i < m as i32 && j >= 0 && j < n as i32
    }

    fn dfs(grid: &mut Vec<Vec<char>>, i: usize, j: usize) {
        let m = grid.len();
        let n = grid[0].len();

        if !Self::is_inside(m, n, i as i32, j as i32) || grid[i][j] != '1' {
            return;
        }

        grid[i][j] = '-';

        if Self::is_inside(m, n, i as i32 - 1, j as i32) {
            Self::dfs(grid, i - 1, j);
        }
        if Self::is_inside(m, n, i as i32, j as i32 + 1) {
            Self::dfs(grid, i, j + 1);
        }
        if Self::is_inside(m, n, i as i32 + 1, j as i32) {
            Self::dfs(grid, i + 1, j);
        }
        if Self::is_inside(m, n, i as i32, j as i32 - 1) {
            Self::dfs(grid, i, j - 1);
        }
    }
}
