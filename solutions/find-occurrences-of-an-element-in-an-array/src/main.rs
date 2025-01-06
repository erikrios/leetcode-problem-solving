fn main() {
    println!(
        "{:?}",
        Solution::occurrences_of_element(vec![1, 3, 1, 7], vec![1, 3, 2, 4], 1)
    );
    println!(
        "{:?}",
        Solution::occurrences_of_element(vec![1, 2, 3], vec![10], 5)
    );
}

struct Solution;

impl Solution {
    pub fn occurrences_of_element(nums: Vec<i32>, queries: Vec<i32>, x: i32) -> Vec<i32> {
        let mut num_occurences = Vec::new();

        for (i, num) in nums.into_iter().enumerate() {
            if num == x {
                num_occurences.push(i);
            }
        }

        let mut results = vec![-1; queries.len()];
        if num_occurences.is_empty() {
            return results;
        }

        for (i, query) in queries.into_iter().enumerate() {
            if (query as usize) <= num_occurences.len() {
                results[i] = num_occurences[query as usize - 1] as i32;
            }
        }

        results
    }
}
