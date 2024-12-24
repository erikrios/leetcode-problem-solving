fn main() {
    println!(
        "{}",
        Solution::unique_paths_iii(vec![vec![1, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 2, -1]])
    );
    println!(
        "{}",
        Solution::unique_paths_iii(vec![vec![1, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 2]])
    );
    println!(
        "{}",
        Solution::unique_paths_iii(vec![vec![0, 1], vec![2, 0]])
    );
}

struct Solution;

impl Solution {
    pub fn unique_paths_iii(grid: Vec<Vec<i32>>) -> i32 {
        let (mut i_start, mut j_start) = (0, 0);
        let mut total_obstacles = 0;

        for (i, nums) in grid.iter().enumerate() {
            for (j, &num) in nums.iter().enumerate() {
                if num == 1 {
                    i_start = i;
                    j_start = j;
                } else if num == -1 {
                    total_obstacles += 1;
                }
            }
        }

        let mut grid = grid;

        Self::dfs(
            &mut grid,
            total_obstacles,
            1,
            i_start as i32,
            j_start as i32,
        )
    }

    fn dfs(grid: &mut [Vec<i32>], obstacle_count: i32, total_visited: i32, i: i32, j: i32) -> i32 {
        let m = grid.len();
        let n = grid[0].len();

        if i < 0 || i >= m as i32 || j < 0 || j >= n as i32 {
            return 0;
        }

        let num = grid[i as usize][j as usize];

        if num == -1 || num == -2 {
            return 0;
        }

        if num == 2 {
            if total_visited == (m * n) as i32 - obstacle_count {
                return 1;
            } else {
                return 0;
            }
        }

        grid[i as usize][j as usize] = -2;

        let mut count = 0;

        let directions = [[0, -1], [1, 0], [0, 1], [-1, 0]];

        for direction in directions {
            count += Self::dfs(
                grid,
                obstacle_count,
                total_visited + 1,
                i + direction[0],
                j + direction[1],
            );
        }

        grid[i as usize][j as usize] = num;

        count
    }
}
