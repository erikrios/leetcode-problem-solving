use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::can_visit_all_rooms(vec![vec![1], vec![2], vec![3], vec![]])
    );
    println!(
        "{}",
        Solution::can_visit_all_rooms(vec![vec![1, 3], vec![3, 0, 1], vec![2], vec![0]])
    );
    println!(
        "{}",
        Solution::can_visit_all_rooms(vec![vec![2], vec![], vec![1]])
    );
}

struct Solution;

impl Solution {
    pub fn can_visit_all_rooms(rooms: Vec<Vec<i32>>) -> bool {
        let n = rooms.len();
        let mut rooms_id = HashSet::new();
        let mut keys = HashSet::new();

        for i in 0..n {
            rooms_id.insert(i);
        }

        while !rooms_id.is_empty() {
            let ids: Vec<usize> = rooms_id.iter().copied().collect();
            let mut opened_count = 0;
            for room in ids {
                if room == 0 || keys.contains(&room) {
                    for &founded_key in &rooms[room] {
                        keys.insert(founded_key as usize);
                    }
                    rooms_id.remove(&room);
                    opened_count += 1;
                }
            }

            if opened_count == 0 {
                return false;
            }
        }

        true
    }
}