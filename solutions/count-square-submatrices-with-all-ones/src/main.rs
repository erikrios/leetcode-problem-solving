fn main() {
    println!(
        "{}",
        Solution::count_squares(vec![vec![0, 1, 1, 1], vec![1, 1, 1, 1], vec![0, 1, 1, 1],])
    );
    println!(
        "{}",
        Solution::count_squares(vec![vec![1, 0, 1], vec![1, 1, 0], vec![1, 1, 0],])
    );
}

struct Solution;

impl Solution {
    pub fn count_squares(matrix: Vec<Vec<i32>>) -> i32 {
        let mut result = 0;

        for (i, nums) in matrix.iter().enumerate() {
            for j in 0..nums.len() {
                result += Self::dfs(i, j, i, j, 1, &matrix);
            }
        }

        result
    }

    fn dfs(
        i_start: usize,
        j_start: usize,
        i_end: usize,
        j_end: usize,
        size: usize,
        matrix: &[Vec<i32>],
    ) -> i32 {
        let m = matrix.len();
        let n = matrix[0].len();

        if i_end >= m || j_end >= n {
            return 0;
        }

        let mut result = 0;

        if size == 1 && matrix[i_end][j_end] == 0 {
            return 0;
        }

        for nums in matrix.iter().take(i_end + 1).skip(i_start) {
            if nums[j_end] == 0 {
                return 0;
            }
        }

        for j in j_start..=j_end {
            if matrix[i_end][j] == 0 {
                return 0;
            }
        }

        result += Self::dfs(i_start, j_start, i_end + 1, j_end + 1, size + 1, matrix);

        result + 1
    }
}
