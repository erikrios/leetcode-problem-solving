fn main() {
    println!("{:?}", Solution::stable_mountains(vec![1, 2, 3, 4, 5], 2));
    println!(
        "{:?}",
        Solution::stable_mountains(vec![10, 1, 10, 1, 10], 3)
    );
    println!(
        "{:?}",
        Solution::stable_mountains(vec![10, 1, 10, 1, 10], 10)
    );
}

struct Solution;

impl Solution {
    pub fn stable_mountains(height: Vec<i32>, threshold: i32) -> Vec<i32> {
        let n = height.len();
        let mut results = Vec::new();

        for i in 1..n {
            if height[i - 1] > threshold {
                results.push(i as i32);
            }
        }

        results
    }
}
