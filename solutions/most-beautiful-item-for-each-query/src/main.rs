use std::vec;

fn main() {
    println!(
        "{:?}",
        Solution::maximum_beauty(
            vec![vec![1, 2], vec![3, 2], vec![2, 4], vec![5, 6], vec![3, 5]],
            vec![1, 2, 3, 4, 5, 6]
        )
    );
    println!(
        "{:?}",
        Solution::maximum_beauty(
            vec![vec![1, 2], vec![1, 2], vec![1, 3], vec![1, 4]],
            vec![1]
        )
    );
    println!(
        "{:?}",
        Solution::maximum_beauty(vec![vec![10, 1000]], vec![5])
    );
}

struct Solution;

impl Solution {
    pub fn maximum_beauty(items: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        let mut items = items;
        items.sort_by_key(|elements| elements[0]);
        let n = items.len();
        let mut max = items[0][1];
        for item in items.iter_mut().take(n) {
            max = max.max(item[1]);
            item[1] = max;
        }

        let mut results = Vec::with_capacity(queries.len());

        for query in queries {
            let mut start = 0_i32;
            let mut end = n as i32 - 1;

            let mut max_beauty = 0;
            while start <= end {
                let mid = (start + end) / 2;
                let found = items[mid as usize][0];

                if found > query {
                    end = mid - 1;
                } else {
                    max_beauty = max_beauty.max(items[mid as usize][1]);
                    start = mid + 1;
                }
            }
            results.push(max_beauty);
        }

        results
    }
}
