use std::collections::HashMap;

fn main() {
    println!(
        "{:?}",
        Solution::query_results(4, vec![vec![1, 4], vec![2, 5], vec![1, 3], vec![3, 4]])
    );
    println!(
        "{:?}",
        Solution::query_results(
            4,
            vec![vec![0, 1], vec![1, 2], vec![2, 2], vec![3, 4], vec![4, 5]]
        )
    );
    println!(
        "{:?}",
        Solution::query_results(
            1,
            vec![vec![0, 1], vec![1, 4], vec![1, 1], vec![1, 4], vec![1, 1]]
        )
    );
}

struct Solution;

impl Solution {
    pub fn query_results(_limit: i32, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let n = queries.len();
        let mut ball_mapper = HashMap::new();
        let mut color_mapper = HashMap::new();
        let mut results = Vec::with_capacity(n);

        for query in queries {
            let ball = query[0];
            let color = query[1];

            if let Some(col) = ball_mapper.get_mut(&ball) {
                if *col != color {
                    let &recent_count = color_mapper.get(col).unwrap();
                    if recent_count == 1 {
                        color_mapper.remove(col);
                    } else {
                        color_mapper.insert(*col, recent_count - 1);
                    }
                    color_mapper
                        .entry(color)
                        .and_modify(|x| *x += 1)
                        .or_insert(1);
                    *col = color;
                }
            } else {
                ball_mapper.insert(ball, color);
                color_mapper
                    .entry(color)
                    .and_modify(|x| *x += 1)
                    .or_insert(1);
            }

            results.push(color_mapper.len() as i32);
        }

        results
    }
}
