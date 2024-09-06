use std::collections::{BTreeMap, BTreeSet};

fn main() {
    println!("{:?}", Solution::get_final_state(vec![2, 1, 3, 5, 6], 5, 2));
    println!("{:?}", Solution::get_final_state(vec![1, 2], 3, 4));
}

struct Solution;

impl Solution {
    pub fn get_final_state(nums: Vec<i32>, k: i32, multiplier: i32) -> Vec<i32> {
        let mut mapper = BTreeMap::new();

        for (i, &num) in nums.iter().enumerate() {
            mapper
                .entry(num)
                .and_modify(|set: &mut BTreeSet<usize>| {
                    set.insert(i);
                })
                .or_insert_with(|| {
                    let mut new_set = BTreeSet::new();
                    new_set.insert(i);
                    new_set
                });
        }

        let mut nums = nums;
        for _ in 0..k {
            let (key, val) = mapper.first_key_value().unwrap();
            let (key, val) = (*key, val.clone());

            let new_num = key * multiplier;
            let index = val.first().unwrap();

            if val.len() == 1 {
                mapper.remove(&key);
            } else {
                let set = mapper.get_mut(&key).unwrap();
                set.remove(index);
            }

            mapper
                .entry(new_num)
                .and_modify(|set| {
                    set.insert(*index);
                })
                .or_insert_with(|| {
                    let mut new_set = BTreeSet::new();
                    new_set.insert(*index);
                    new_set
                });

            nums[*index] = new_num;
        }

        nums
    }
}
