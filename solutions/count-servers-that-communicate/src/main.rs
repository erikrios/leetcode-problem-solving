use std::collections::HashMap;

fn main() {
    println!("{}", Solution::count_servers(vec![vec![1, 0], vec![0, 1]]));
    println!("{}", Solution::count_servers(vec![vec![1, 0], vec![1, 1]]));
    println!(
        "{}",
        Solution::count_servers(vec![
            vec![1, 1, 0, 0],
            vec![0, 0, 1, 0],
            vec![0, 0, 1, 0],
            vec![0, 0, 0, 1]
        ])
    );
}

struct Solution;

impl Solution {
    pub fn count_servers(grid: Vec<Vec<i32>>) -> i32 {
        let mut vertical_map = HashMap::new();
        let mut horizontal_map = HashMap::new();

        for (i, nums) in grid.iter().enumerate() {
            for (j, &num) in nums.iter().enumerate() {
                if num == 1 {
                    vertical_map.entry(i).and_modify(|e| *e += 1).or_insert(1);
                    horizontal_map.entry(j).and_modify(|e| *e += 1).or_insert(1);
                }
            }
        }

        let mut count = 0;
        for (i, nums) in grid.iter().enumerate() {
            for (j, &num) in nums.iter().enumerate() {
                if num == 1 {
                    if let Some(&ver_count) = vertical_map.get(&i) {
                        if ver_count > 1 {
                            count += 1;
                            continue;
                        }
                    }

                    if let Some(&hor_count) = horizontal_map.get(&j) {
                        if hor_count > 1 {
                            count += 1;
                        }
                    }
                }
            }
        }

        count
    }
}
