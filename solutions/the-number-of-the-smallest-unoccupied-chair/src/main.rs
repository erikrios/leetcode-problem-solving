use std::collections::{BTreeSet, HashMap};

fn main() {
    println!(
        "{}",
        Solution::smallest_chair(vec![vec![1, 4], vec![2, 3], vec![4, 6]], 1)
    );
    println!(
        "{}",
        Solution::smallest_chair(vec![vec![3, 10], vec![1, 5], vec![2, 6]], 0)
    );
}

struct Solution;

struct Occupied;

impl Solution {
    pub fn smallest_chair(times: Vec<Vec<i32>>, target_friend: i32) -> i32 {
        let mut list_time_b_tree_set = BTreeSet::new();
        let mut arrival_mapper = HashMap::new();
        let mut leaving_mapper = HashMap::new();

        for (i, time) in times.into_iter().enumerate() {
            let arrival = time[0];
            let leaving = time[1];

            list_time_b_tree_set.insert(arrival);
            list_time_b_tree_set.insert(leaving);

            arrival_mapper.insert(arrival, i);
            leaving_mapper
                .entry(leaving)
                .and_modify(|x: &mut Vec<usize>| x.push(i))
                .or_insert(vec![i]);
        }

        let mut unoccupied_chairs_b_tree_set = BTreeSet::new();
        let mut chairs_mapper = HashMap::new();
        let mut chairs = Vec::new();

        for time in list_time_b_tree_set {
            if let Some(user_ids) = leaving_mapper.remove(&time) {
                for user_id in user_ids {
                    let &chair_pos = chairs_mapper.get(&user_id).unwrap();
                    unoccupied_chairs_b_tree_set.insert(chair_pos);
                }
            }

            if let Some(&user_id) = arrival_mapper.get(&time) {
                let chair_pos;
                if unoccupied_chairs_b_tree_set.is_empty() {
                    chairs.push(Occupied);
                    chair_pos = chairs.len() - 1;
                } else {
                    chair_pos = *unoccupied_chairs_b_tree_set.iter().next().unwrap();
                    unoccupied_chairs_b_tree_set.remove(&chair_pos);
                }
                chairs_mapper.insert(user_id, chair_pos);
                if user_id == target_friend as usize {
                    return chair_pos as i32;
                }
            }
        }

        0
    }
}
