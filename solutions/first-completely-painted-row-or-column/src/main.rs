use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::first_complete_index(vec![1, 3, 4, 2], vec![vec![1, 4], vec![2, 3]])
    );
    println!(
        "{}",
        Solution::first_complete_index(
            vec![2, 8, 7, 4, 1, 3, 5, 6, 9],
            vec![vec![3, 2, 5], vec![1, 4, 6], vec![8, 7, 9]]
        )
    );
}

struct Solution;

impl Solution {
    pub fn first_complete_index(arr: Vec<i32>, mat: Vec<Vec<i32>>) -> i32 {
        let m = mat.len();
        let n = mat[0].len();

        let mut matrix_mapper = HashMap::new();

        for (i, nums) in mat.into_iter().enumerate() {
            for (j, num) in nums.into_iter().enumerate() {
                matrix_mapper.insert(num, (i, j));
            }
        }

        let mut row_counter = HashMap::with_capacity(n);
        let mut col_counter = HashMap::with_capacity(m);

        for (i, num) in arr.into_iter().enumerate() {
            let index = matrix_mapper.get(&num).unwrap();

            row_counter
                .entry(index.0)
                .and_modify(|e| *e += 1)
                .or_insert(1);

            if *row_counter.get(&index.0).unwrap() == n {
                return i as i32;
            }

            col_counter
                .entry(index.1)
                .and_modify(|e| *e += 1)
                .or_insert(1);

            if *col_counter.get(&index.1).unwrap() == m {
                return i as i32;
            }
        }

        0
    }
}
