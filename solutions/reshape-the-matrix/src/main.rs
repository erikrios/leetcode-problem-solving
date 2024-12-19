fn main() {
    println!(
        "{:?}",
        Solution::matrix_reshape(vec![vec![1, 2], vec![3, 4]], 1, 4)
    );
    println!(
        "{:?}",
        Solution::matrix_reshape(vec![vec![1, 2], vec![3, 4]], 2, 4)
    );
}

struct Solution;

impl Solution {
    pub fn matrix_reshape(mat: Vec<Vec<i32>>, r: i32, c: i32) -> Vec<Vec<i32>> {
        let r = r as usize;
        let c = c as usize;
        let m = mat.len();
        let n = mat[0].len();

        if m * n != r * c {
            return mat;
        }

        let mut results = vec![vec![0; c]; r];
        let mut i = 0;
        let mut j = 0;

        for nums in mat {
            for num in nums {
                results[i][j] = num;
                j += 1;
                if j == c {
                    j = 0;
                    i += 1;
                }
            }
        }

        results
    }
}
