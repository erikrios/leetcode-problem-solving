fn main() {
    println!("{:?}", Solution::restore_matrix(vec![3, 8], vec![4, 7]));
    println!(
        "{:?}",
        Solution::restore_matrix(vec![5, 7, 10], vec![8, 6, 8])
    );
    println!(
        "{:?}",
        Solution::restore_matrix(vec![7, 9, 12], vec![8, 6, 8, 8])
    );
}

struct Solution;

impl Solution {
    pub fn restore_matrix(row_sum: Vec<i32>, col_sum: Vec<i32>) -> Vec<Vec<i32>> {
        let mut row_sum = row_sum;
        let mut col_sum = col_sum;
        let mut results = Vec::with_capacity(row_sum.len());

        for i in 0..row_sum.len() {
            let mut result = Vec::with_capacity(col_sum.len());
            for j in 0..col_sum.len() {
                let smallest = if row_sum[i] < col_sum[j] {
                    row_sum[i]
                } else {
                    col_sum[j]
                };
                row_sum[i] -= smallest;
                col_sum[j] -= smallest;
                result.push(smallest);
            }
            results.push(result);
        }

        results
    }
}
