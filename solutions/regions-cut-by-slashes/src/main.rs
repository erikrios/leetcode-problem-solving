use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::regions_by_slashes(vec![" /".to_string(), "/ ".to_string()])
    );
    println!(
        "{}",
        Solution::regions_by_slashes(vec![" /".to_string(), "  ".to_string()])
    );
    println!(
        "{}",
        Solution::regions_by_slashes(vec!["/\\".to_string(), "\\/".to_string()])
    );
}

struct Solution;

impl Solution {
    pub fn regions_by_slashes(grid: Vec<String>) -> i32 {
        let n = grid.len();

        let mut new_grid = vec![vec![0; n * 3]; n * 3];

        for (i, squares) in grid.iter().enumerate() {
            for (j, square) in squares.chars().enumerate() {
                if square == '/' {
                    new_grid[i * 3][j * 3 + 2] = 1;
                    new_grid[i * 3 + 1][j * 3 + 1] = 1;
                    new_grid[i * 3 + 2][j * 3] = 1;
                } else if square == '\\' {
                    new_grid[i * 3][j * 3] = 1;
                    new_grid[i * 3 + 1][j * 3 + 1] = 1;
                    new_grid[i * 3 + 2][j * 3 + 2] = 1;
                }
            }
        }

        let mut visited = HashSet::new();
        let mut num_of_regions = 0;
        let n = new_grid.len();

        for i in 0..n {
            for j in 0..n {
                if visited.contains(&(i, j)) || new_grid[i][j] == 1 {
                    continue;
                }
                Self::dfs(&new_grid, i, j, &mut visited);
                num_of_regions += 1;
            }
        }

        num_of_regions
    }

    fn dfs(grid: &[Vec<i32>], i: usize, j: usize, visited: &mut HashSet<(usize, usize)>) {
        let n = grid.len();
        if i >= n || j >= n || grid[i][j] == 1 || visited.contains(&(i, j)) {
            return;
        }

        visited.insert((i, j));

        let mut neigbors = vec![(i, j + 1), (i + 1, j)];
        if i > 0 {
            neigbors.push((i - 1, j));
        }
        if j > 0 {
            neigbors.push((i, j - 1));
        }

        for neighbor in neigbors {
            Self::dfs(grid, neighbor.0, neighbor.1, visited);
        }
    }
}
