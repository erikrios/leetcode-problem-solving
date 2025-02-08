fn main() {
    println!(
        "{}",
        Solution::grid_game(vec![vec![2, 5, 4], vec![1, 5, 1]])
    );
    println!(
        "{}",
        Solution::grid_game(vec![vec![3, 3, 1], vec![8, 5, 2]])
    );
    println!(
        "{}",
        Solution::grid_game(vec![vec![1, 3, 1, 15], vec![1, 3, 3, 1]])
    );
    println!(
        "{}",
        Solution::grid_game(vec![
            vec![20, 3, 20, 17, 2, 12, 15, 17, 4, 15],
            vec![20, 10, 13, 14, 15, 5, 2, 3, 14, 3]
        ])
    );
}

struct Solution;

impl Solution {
    pub fn grid_game(grid: Vec<Vec<i32>>) -> i64 {
        let n = grid[0].len();
        let mut prefix_sum_row_one: Vec<i64> = grid[0].iter().map(|&x| x as i64).collect();
        let mut prefix_sum_row_two: Vec<i64> = grid[1].iter().map(|&x| x as i64).collect();

        for i in 1..n {
            prefix_sum_row_one[i] += prefix_sum_row_one[i - 1];
            prefix_sum_row_two[i] += prefix_sum_row_two[i - 1];
        }

        let mut res = i64::MAX;
        for i in 0..n {
            let row_one = prefix_sum_row_one[n - 1] - prefix_sum_row_one[i];
            let row_two = if i > 0 { prefix_sum_row_two[i - 1] } else { 0 };

            let robot_two_max_collected_point = row_one.max(row_two);
            res = res.min(robot_two_max_collected_point);
        }

        res
    }
}
