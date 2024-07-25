fn main() {
    println!("{}", Solution::count_triplets(vec![2, 3, 1, 6, 7]));
    println!("{}", Solution::count_triplets(vec![1, 1, 1, 1, 1]));
}

struct Solution;

impl Solution {
    pub fn count_triplets(arr: Vec<i32>) -> i32 {
        let mut counter = 0;
        let n = arr.len();

        for i in 0..n - 1 {
            let mut xor = arr[i];
            for j in i + 1..n {
                xor ^= arr[j];
                if xor == 0 {
                    counter += (j-i) as i32;
                }
            }
        }

        counter
    }
}
