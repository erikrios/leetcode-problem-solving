use std::cmp::max;
use std::collections::HashMap;
use std::vec;

fn main() {
    println!("{}", Solution::total_fruit(vec![1, 2, 1]));
    println!("{}", Solution::total_fruit(vec![0, 1, 2, 2]));
    println!("{}", Solution::total_fruit(vec![1, 2, 3, 2, 2]));
}

struct Solution;

impl Solution {
    pub fn total_fruit(fruits: Vec<i32>) -> i32 {
        let mut mapper = HashMap::new();
        let mut window_start = 0;
        let mut maximum = 0i32;

        for (window_end, &fruit) in fruits.iter().enumerate() {
            mapper.insert(fruit, mapper.get(&fruit).unwrap_or(&0) + 1);

            while mapper.len() > 2 {
                if let Some(value) = mapper.get(&fruits[window_start]) {
                    if *value == 1 {
                        mapper.remove(&fruits[window_start]);
                    } else {
                        mapper.insert(fruits[window_start], value - 1);
                    }
                }
                window_start += 1;
            }

            maximum = max(maximum, (window_end - window_start) as i32 + 1);
        }

        maximum
    }
}
