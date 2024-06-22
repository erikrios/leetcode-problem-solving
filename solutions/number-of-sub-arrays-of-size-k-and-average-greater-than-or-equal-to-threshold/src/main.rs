fn main() {
    println!(
        "{}",
        Solution::num_of_subarrays(vec![2, 2, 2, 2, 5, 5, 5, 8], 3, 4)
    );
    println!(
        "{}",
        Solution::num_of_subarrays(vec![11, 13, 17, 23, 29, 31, 7, 5, 2, 3], 3, 5)
    );
}

struct Solution;

impl Solution {
    pub fn num_of_subarrays(arr: Vec<i32>, k: i32, threshold: i32) -> i32 {
        let mut start_window = 0;
        let mut current_sum = 0;
        let mut counter = 0;

        for (end_window, num) in arr.iter().enumerate() {
            current_sum += num;
            if end_window >= k as usize - 1 {
                let average = current_sum / k;
                if average >= threshold {
                    counter += 1;
                }
                current_sum -= arr[start_window];
                start_window += 1;
            }
        }

        counter
    }
}
