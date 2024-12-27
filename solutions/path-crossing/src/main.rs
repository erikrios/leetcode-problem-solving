use std::collections::HashSet;

fn main() {
    println!("{}", Solution::is_path_crossing("NES".to_string()));
    println!("{}", Solution::is_path_crossing("NESWW".to_string()));
    println!(
        "{}",
        Solution::is_path_crossing("NNSWWEWSSESSWENNW".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn is_path_crossing(path: String) -> bool {
        let mut point = (0, 0);
        let mut visited = HashSet::new();
        visited.insert(point);

        for ch in path.into_bytes() {
            match ch {
                b'N' => point.1 -= 1,
                b'S' => point.1 += 1,
                b'E' => point.0 -= 1,
                b'W' => point.0 += 1,
                _ => {}
            }

            if visited.contains(&point) {
                return true;
            }

            visited.insert(point);
        }

        false
    }
}
