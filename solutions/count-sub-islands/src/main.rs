use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::count_sub_islands(
            vec![
                vec![1, 1, 1, 0, 0],
                vec![0, 1, 1, 1, 1],
                vec![0, 0, 0, 0, 0],
                vec![1, 0, 0, 0, 0],
                vec![1, 1, 0, 1, 1]
            ],
            vec![
                vec![1, 1, 1, 0, 0],
                vec![0, 0, 1, 1, 1],
                vec![0, 1, 0, 0, 0],
                vec![1, 0, 1, 1, 0],
                vec![0, 1, 0, 1, 0]
            ]
        )
    );
    println!(
        "{}",
        Solution::count_sub_islands(
            vec![
                vec![1, 0, 1, 0, 1],
                vec![1, 1, 1, 1, 1],
                vec![0, 0, 0, 0, 0],
                vec![1, 1, 1, 1, 1],
                vec![1, 0, 1, 0, 1]
            ],
            vec![
                vec![0, 0, 0, 0, 0],
                vec![1, 1, 1, 1, 1],
                vec![0, 1, 0, 1, 0],
                vec![0, 1, 0, 1, 0],
                vec![1, 0, 0, 0, 1]
            ]
        )
    );
}

struct Solution;

impl Solution {
    pub fn count_sub_islands(grid1: Vec<Vec<i32>>, grid2: Vec<Vec<i32>>) -> i32 {
        let m = grid2.len();
        let n = grid2[0].len();

        let mut counter = 0;
        let mut visited = HashSet::new();
        for i in 0..m {
            for j in 0..n {
                if grid2[i][j] == 1
                    && !visited.contains(&(i, j))
                    && Self::visit(&grid1, &grid2, &mut visited, i, j)
                {
                    counter += 1;
                }
            }
        }

        counter
    }

    fn visit(
        grid1: &Vec<Vec<i32>>,
        grid2: &Vec<Vec<i32>>,
        visited: &mut HashSet<(usize, usize)>,
        i: usize,
        j: usize,
    ) -> bool {
        let m = grid1.len();
        let n = grid1[0].len();

        visited.insert((i, j));

        let mut res = true;
        let num = grid2[i][j];
        res &= num == grid1[i][j];

        if i > 0 && grid2[i - 1][j] != 0 && !visited.contains(&(i - 1, j)) {
            res &= Self::visit(grid1, grid2, visited, i - 1, j);
        }

        if j > 0 && grid2[i][j - 1] != 0 && !visited.contains(&(i, j - 1)) {
            res &= Self::visit(grid1, grid2, visited, i, j - 1);
        }

        if i + 1 < m && grid2[i + 1][j] != 0 && !visited.contains(&(i + 1, j)) {
            res &= Self::visit(grid1, grid2, visited, i + 1, j);
        }

        if j + 1 < n && grid2[i][j + 1] != 0 && !visited.contains(&(i, j + 1)) {
            res &= Self::visit(grid1, grid2, visited, i, j + 1);
        }

        res
    }
}
