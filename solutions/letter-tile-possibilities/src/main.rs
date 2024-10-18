use std::collections::HashSet;

fn main() {
    println!("{}", Solution::num_tile_possibilities("AAB".to_string()));
    println!("{}", Solution::num_tile_possibilities("AAABBC".to_string()));
    println!("{}", Solution::num_tile_possibilities("V".to_string()));
}

struct Solution;

impl Solution {
    pub fn num_tile_possibilities(tiles: String) -> i32 {
        let mut results = HashSet::new();
        let n = tiles.len();

        for i in 0..n {
            for j in 0..n {
                Self::backtracking(
                    &tiles,
                    i + 1,
                    j,
                    &mut HashSet::new(),
                    &mut Vec::new(),
                    &mut results,
                );
            }
        }

        results.len() as i32
    }

    fn backtracking(
        tiles: &str,
        n: usize,
        i: usize,
        visited: &mut HashSet<usize>,
        result: &mut Vec<u8>,
        results: &mut HashSet<String>,
    ) {
        let tile = tiles.as_bytes()[i];
        result.push(tile);
        visited.insert(i);

        if n == 1 {
            results.insert(String::from_utf8(result.to_vec()).unwrap());
            return;
        }

        for idx in 0..tiles.len() {
            if !visited.contains(&idx) {
                Self::backtracking(tiles, n - 1, idx, visited, result, results);
                visited.remove(&idx);
                result.pop();
            }
        }
    }
}
