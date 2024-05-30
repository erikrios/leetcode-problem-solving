use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::count_good_rectangles(vec![vec![5, 8], vec![3, 9], vec![5, 12], vec![16, 5]])
    );
    println!(
        "{}",
        Solution::count_good_rectangles(vec![vec![2, 3], vec![3, 7], vec![4, 3], vec![3, 7]])
    );
}

struct Solution;

impl Solution {
    pub fn count_good_rectangles(rectangles: Vec<Vec<i32>>) -> i32 {
        let mut max: i32 = 0;
        let mut map_counter = HashMap::<i32, i32>::new();

        for rect in rectangles {
            let k = if rect[0] < rect[1] { rect[0] } else { rect[1] };

            if k > max {
                max = k;
            }

            let cur = match map_counter.get(&k) {
                Some(val) => *val,
                None => 0,
            };

            map_counter.insert(k, cur + 1);
        }

        *map_counter.get(&max).unwrap()
    }
}
