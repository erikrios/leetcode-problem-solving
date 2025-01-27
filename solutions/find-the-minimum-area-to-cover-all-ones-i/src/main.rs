fn main() {
    println!(
        "{}",
        Solution::minimum_area(vec![vec![0, 1, 0], vec![1, 0, 1],])
    );
    println!("{}", Solution::minimum_area(vec![vec![1, 0], vec![0, 0],]));
}

struct Solution;

impl Solution {
    pub fn minimum_area(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();

        let mut left_most_index = n - 1;
        let mut top_most_index = m - 1;
        let mut right_most_index = 0;
        let mut bottom_most_index = 0;

        for (i, nums) in grid.iter().enumerate().take(m) {
            for (j, &num) in nums.iter().enumerate().take(n) {
                if num == 1 {
                    left_most_index = left_most_index.min(j);
                    top_most_index = top_most_index.min(i);
                    right_most_index = right_most_index.max(j);
                    bottom_most_index = bottom_most_index.max(i);
                }
            }
        }

        let width = right_most_index - left_most_index + 1;
        let height = bottom_most_index - top_most_index + 1;

        (width * height) as i32
    }
}
