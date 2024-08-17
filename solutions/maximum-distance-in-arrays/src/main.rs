fn main() {
    println!(
        "{}",
        Solution::max_distance(vec![vec![1, 2, 3], vec![4, 5], vec![1, 2, 3]])
    );
    println!("{}", Solution::max_distance(vec![vec![1], vec![1],]));
    println!(
        "{}",
        Solution::max_distance(vec![vec![1, 2, 3, 4, 5, 6], vec![4, 5], vec![2, 3]])
    );
    println!("{}", Solution::max_distance(vec![vec![1, 5], vec![3, 4],]));
}

struct Solution;

impl Solution {
    pub fn max_distance(arrays: Vec<Vec<i32>>) -> i32 {
        let mut res = i32::MIN;
        let mut min = arrays[0][0];
        let mut max = arrays[0][arrays[0].len() - 1];

        for array in arrays.iter().skip(1) {
            let cur_min = array[0];
            let cur_max = array[array.len() - 1];

            res = std::cmp::max(res, (cur_max - min).abs());
            res = std::cmp::max(res, (max - cur_min).abs());

            min = std::cmp::min(min, cur_min);
            max = std::cmp::max(max, cur_max);
        }

        res
    }
}
