use std::collections::HashSet;

fn main() {
    println!(
        "{:?}",
        Solution::queens_attackthe_king(
            vec![
                vec![0, 1],
                vec![1, 0],
                vec![4, 0],
                vec![0, 4],
                vec![3, 3],
                vec![2, 4]
            ],
            vec![0, 0]
        )
    );

    println!(
        "{:?}",
        Solution::queens_attackthe_king(
            vec![
                vec![0, 0],
                vec![1, 1],
                vec![2, 2],
                vec![3, 4],
                vec![3, 5],
                vec![4, 4],
                vec![4, 5]
            ],
            vec![3, 3]
        )
    );
}

struct Solution;

const DIRECTIONS: [(i32, i32); 8] = [
    (-1, 0),
    (-1, 1),
    (0, 1),
    (1, 1),
    (1, 0),
    (1, -1),
    (0, -1),
    (-1, -1),
];

impl Solution {
    pub fn queens_attackthe_king(queens: Vec<Vec<i32>>, king: Vec<i32>) -> Vec<Vec<i32>> {
        let mut queen_coordinates = HashSet::new();

        for queen in queens {
            queen_coordinates.insert((queen[0], queen[1]));
        }

        const N: usize = 8;

        let king = (king[0], king[1]);
        let mut results = Vec::new();

        for dir in DIRECTIONS {
            if let Some(coordinate) = Self::dfs(N, &queen_coordinates, king, dir) {
                results.push(vec![coordinate.0, coordinate.1]);
            }
        }

        results
    }

    fn dfs(
        n: usize,
        queen_coordinates: &HashSet<(i32, i32)>,
        king: (i32, i32),
        dir: (i32, i32),
    ) -> Option<(i32, i32)> {
        if king.0 < 0 || king.0 >= n as i32 || king.1 < 0 || king.1 >= n as i32 {
            return None;
        }

        if queen_coordinates.contains(&king) {
            return Some(king);
        }

        Self::dfs(n, queen_coordinates, (king.0 + dir.0, king.1 + dir.1), dir)
    }
}
