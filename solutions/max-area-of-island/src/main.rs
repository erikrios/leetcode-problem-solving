fn main() {
    println!(
        "{}",
        Solution::max_area_of_island(vec![
            vec![0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0]
        ])
    );
    println!(
        "{}",
        Solution::max_area_of_island(vec![vec![0, 0, 0, 0, 0, 0, 0, 0]])
    );
    println!(
        "{}",
        Solution::max_area_of_island(vec![
            vec![1, 1, 0, 0, 0],
            vec![1, 1, 0, 0, 0],
            vec![0, 0, 0, 1, 1],
            vec![0, 0, 0, 1, 1]
        ])
    );
}

struct Solution;

impl Solution {
    pub fn max_area_of_island(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut grid = grid;
        let mut res = 0;

        for i in 0..m {
            for j in 0..n {
                if grid[i][j] == 1 {
                    res = res.max(Self::mark_as_visited(i, j, &mut grid));
                }
            }
        }

        res as i32
    }

    fn mark_as_visited(i: usize, j: usize, grid: &mut [Vec<i32>]) -> usize {
        let m = grid.len();
        let n = grid[0].len();

        if i >= m || j >= n || grid[i][j] == 0 {
            return 0;
        }

        grid[i][j] = 0;
        let mut area = 1;

        if i > 0 {
            area += Self::mark_as_visited(i - 1, j, grid);
        }
        if j > 0 {
            area += Self::mark_as_visited(i, j - 1, grid);
        }
        area += Self::mark_as_visited(i + 1, j, grid);
        area += Self::mark_as_visited(i, j + 1, grid);

        area
    }
}
