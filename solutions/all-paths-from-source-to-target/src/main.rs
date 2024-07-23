fn main() {
    println!(
        "{:?}",
        Solution::all_paths_source_target(vec![vec![1, 2], vec![3], vec![3], vec![]])
    );
    println!(
        "{:?}",
        Solution::all_paths_source_target(vec![
            vec![4, 3, 1],
            vec![3, 2, 4],
            vec![3],
            vec![4],
            vec![]
        ])
    );
}

struct Solution;

impl Solution {
    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let path_nums = Vec::new();
        Self::paths(0, &graph, path_nums)
    }

    fn paths(i: usize, graph: &Vec<Vec<i32>>, path_nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut path_nums = path_nums.clone();
        path_nums.push(i as i32);
        if i == graph.len() - 1 {
            return vec![path_nums];
        }

        let mut results = vec![];
        for &dest in &graph[i] {
            let result = Self::paths(dest as usize, graph, path_nums.clone());
            results.extend(result.clone());
        }

        results
    }
}
