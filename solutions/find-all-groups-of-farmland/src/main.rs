fn main() {
    println!(
        "{:?}",
        Solution::find_farmland(vec![vec![1, 0, 0], vec![0, 1, 1], vec![0, 1, 1]])
    );
    println!(
        "{:?}",
        Solution::find_farmland(vec![vec![1, 1], vec![1, 1]])
    );
    println!("{:?}", Solution::find_farmland(vec![vec![0]]));
}

struct Solution;

impl Solution {
    pub fn find_farmland(land: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let m = land.len();
        let n = land[0].len();
        let mut land = land;
        let mut farmland_group_coordinates = Vec::new();

        for i in 0..m {
            for j in 0..n {
                if land[i][j] == 0 {
                    continue;
                }

                let bottom_right = Self::dfs(&mut land, i, j);
                farmland_group_coordinates.push(vec![
                    i as i32,
                    j as i32,
                    bottom_right[0] as i32,
                    bottom_right[1] as i32,
                ]);
            }
        }

        farmland_group_coordinates
    }

    fn dfs(land: &mut [Vec<i32>], i: usize, j: usize) -> [usize; 2] {
        let m = land.len();
        let n = land[0].len();

        land[i][j] = 0;

        let mut bottom_right = [i, j];

        if i + 1 < m && land[i + 1][j] == 1 {
            let results = Self::dfs(land, i + 1, j);
            if results[0] > bottom_right[0] || results[1] > bottom_right[1] {
                bottom_right = results;
            }
        }

        if j + 1 < n && land[i][j + 1] == 1 {
            let results = Self::dfs(land, i, j + 1);
            if results[0] > bottom_right[0] || results[1] > bottom_right[1] {
                bottom_right = results;
            }
        }

        bottom_right
    }
}
