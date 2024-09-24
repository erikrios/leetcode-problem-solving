fn main() {
    println!(
        "{}",
        Solution::num_submatrix_sum_target(vec![vec![0, 1, 0], vec![1, 1, 1], vec![0, 1, 0]], 0)
    );
    println!(
        "{}",
        Solution::num_submatrix_sum_target(vec![vec![1, -1], vec![-1, 1]], 0)
    );
    println!("{}", Solution::num_submatrix_sum_target(vec![vec![904]], 0));
}

struct Solution;

impl Solution {
    pub fn num_submatrix_sum_target(matrix: Vec<Vec<i32>>, target: i32) -> i32 {
        let m = matrix.len();
        let n = matrix[0].len();

        let mut matrix = matrix.clone();
        let mut total = 0;

        for i in 0..m {
            for j in 1..n {
                matrix[i][j] += matrix[i][j - 1];
            }
        }

        for left in 0..n {
            for right in left..n {
                let mut sum_counts = std::collections::HashMap::new();
                sum_counts.insert(0, 1);
                let mut current_sum = 0;

                for i in 0..m {
                    current_sum += if left > 0 {
                        matrix[i][right] - matrix[i][left - 1]
                    } else {
                        matrix[i][right]
                    };

                    if let Some(count) = sum_counts.get(&(current_sum - target)) {
                        total += count;
                    }

                    *sum_counts.entry(current_sum).or_insert(0) += 1;
                }
            }
        }

        total
    }
}
