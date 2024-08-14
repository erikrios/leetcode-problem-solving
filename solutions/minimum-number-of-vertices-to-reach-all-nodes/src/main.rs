fn main() {
    println!(
        "{:?}",
        Solution::find_smallest_set_of_vertices(
            6,
            vec![vec![0, 1], vec![0, 2], vec![2, 5], vec![3, 4], vec![4, 2]]
        )
    );
    println!(
        "{:?}",
        Solution::find_smallest_set_of_vertices(
            5,
            vec![vec![0, 1], vec![2, 1], vec![3, 1], vec![1, 4], vec![2, 4]]
        )
    );
}

struct Solution;

impl Solution {
    pub fn find_smallest_set_of_vertices(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let mut vertices = vec![false; n as usize];

        for edge in edges.iter() {
            vertices[edge[1] as usize] = true;
        }

        (0..n as usize)
            .filter(|&i| !vertices[i])
            .map(|i| i as i32)
            .collect()
    }
}
