fn main() {
    println!(
        "{:?}",
        Solution::matrix_block_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]], 1)
    );
    println!(
        "{:?}",
        Solution::matrix_block_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]], 2)
    );
}

struct Solution;

impl Solution {
    pub fn matrix_block_sum(mat: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let m = mat.len();
        let n = mat[0].len();
        let k = k as usize;

        let mut prefix_sums = Vec::with_capacity(m);
        for nums in &mat {
            let mut prefix_sum = Vec::with_capacity(n);
            for (j, &num) in nums.iter().enumerate() {
                if j == 0 {
                    prefix_sum.push(num);
                } else {
                    prefix_sum.push(prefix_sum[j - 1] + num);
                }
            }
            prefix_sums.push(prefix_sum);
        }

        let mut results = Vec::with_capacity(m);

        for i in 0..m {
            let i_start = if (i as i32) - (k as i32) < 0 {
                0
            } else {
                i - k
            };
            let i_end = if i + k >= m { m - 1 } else { i + k };
            let mut result = Vec::with_capacity(n);
            for j in 0..n {
                let j_start = if (j as i32) - (k as i32) < 0 {
                    0
                } else {
                    j - k
                };
                let j_end = if j + k >= n { n - 1 } else { j + k };

                let mut sum = 0;
                for prefix_sum in prefix_sums.iter().take(i_end + 1).skip(i_start) {
                    sum += if j_start == 0 {
                        prefix_sum[j_end]
                    } else {
                        prefix_sum[j_end] - prefix_sum[j_start - 1]
                    };
                }

                result.push(sum)
            }
            results.push(result);
        }

        results
    }
}
