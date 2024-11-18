use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::max_moves(vec![
            vec![2, 4, 3, 5],
            vec![5, 4, 9, 3],
            vec![3, 4, 2, 11],
            vec![10, 9, 13, 15]
        ])
    );
    println!(
        "{}",
        Solution::max_moves(vec![vec![3, 2, 4], vec![2, 1, 9], vec![1, 1, 7]])
    );
}

struct Solution;

const DIRECTIONS: [(i32, i32); 3] = [(-1, 1), (0, 1), (1, 1)];

impl Solution {
    pub fn max_moves(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();

        let mut max = 0;
        let mut visited = HashMap::new();

        for i in 0..m {
            max = max.max(Self::dfs(i as i32, 0, 0, &grid, &mut visited));
        }

        max - 1
    }

    fn dfs(
        i: i32,
        j: i32,
        prev: i32,
        grid: &[Vec<i32>],
        visited: &mut HashMap<(i32, i32), i32>,
    ) -> i32 {
        if let Some(&val) = visited.get(&(i, j)) {
            return val;
        }

        let m = grid.len();
        let n = grid[0].len();

        if i >= m as i32 || i < 0 || j >= n as i32 {
            return 0;
        }

        let num = grid[i as usize][j as usize];
        if num <= prev {
            return 0;
        }

        let mut max = 0;

        for dir in DIRECTIONS {
            max = max.max(Self::dfs(i + dir.0, j + dir.1, num, grid, visited));
        }

        visited.insert((i, j), max + 1);
        max + 1
    }
}
