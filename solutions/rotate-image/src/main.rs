fn main() {
    let mut matrix = vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]];
    Solution::rotate(&mut matrix);
    println!("{:?}", matrix);
    let mut matrix = vec![
        vec![5, 1, 9, 11],
        vec![2, 4, 8, 10],
        vec![13, 3, 6, 7],
        vec![15, 14, 12, 16],
    ];
    Solution::rotate(&mut matrix);
    println!("{:?}", matrix);
}

struct Solution;

impl Solution {
    pub fn rotate(matrix: &mut [Vec<i32>]) {
        let n = matrix.len();
        let mut i_start = 0;
        let mut j_start = 0;
        let mut i_end = n - 1;
        let mut j_end = n - 1;

        while i_start < i_end {
            Self::rotation(matrix, i_start, j_start, i_end, j_end);
            i_start += 1;
            j_start += 1;
            i_end -= 1;
            j_end -= 1;
        }
    }

    fn rotation(
        matrix: &mut [Vec<i32>],
        i_start: usize,
        j_start: usize,
        i_end: usize,
        j_end: usize,
    ) {
        let times = i_end - i_start;

        for _ in 0..times {
            let tmp = matrix[i_start][j_start];

            for left in i_start..i_end {
                matrix[left][j_start] = matrix[left + 1][j_start];
            }

            for bottom in j_start..j_end {
                matrix[i_end][bottom] = matrix[i_end][bottom + 1];
            }

            for right in (i_start..i_end).rev() {
                matrix[right + 1][j_end] = matrix[right][j_end];
            }

            for top in (j_start..j_end).rev() {
                matrix[i_start][top + 1] = matrix[i_start][top];
            }

            matrix[i_start][j_start + 1] = tmp
        }
    }
}
