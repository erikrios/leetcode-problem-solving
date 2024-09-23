fn main() {
    println!(
        "{:?}",
        Solution::xor_queries(
            vec![1, 3, 4, 8],
            vec![vec![0, 1], vec![1, 2], vec![0, 3], vec![3, 3]]
        )
    );
    println!(
        "{:?}",
        Solution::xor_queries(
            vec![4, 8, 2, 10],
            vec![vec![2, 3], vec![1, 3], vec![0, 0], vec![0, 3]]
        )
    );
}

struct Solution;

impl Solution {
    pub fn xor_queries(arr: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let n = arr.len();
        let mut prefix_xor = Vec::with_capacity(n);
        let mut total = 0;

        for (i, &num) in arr.iter().enumerate() {
            if i == 0 {
                total = num;
            } else {
                total ^= num;
            }

            prefix_xor.push(total);
        }

        let mut results = Vec::with_capacity(queries.len());

        for query in queries {
            let start_index = query[0] as usize;
            let end_index = query[1] as usize;

            let left_of_query = if start_index == 0 {
                0
            } else {
                prefix_xor[start_index - 1]
            };

            let rem_two_right = left_of_query ^ total;

            let till_last_query = prefix_xor[end_index];

            let rem_last_right = till_last_query ^ total;

            let res = rem_two_right ^ rem_last_right;

            results.push(res);
        }

        results
    }
}
