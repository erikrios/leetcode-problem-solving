fn main() {
    println!("{:?}", Solution::find_even_numbers(vec![2, 1, 3, 0]));
    println!("{:?}", Solution::find_even_numbers(vec![2, 2, 8, 8, 2]));
    println!("{:?}", Solution::find_even_numbers(vec![3, 7, 5]));
    println!("{:?}", Solution::find_even_numbers(vec![0, 2, 0, 0]));
}

struct Solution;

impl Solution {
    pub fn find_even_numbers(digits: Vec<i32>) -> Vec<i32> {
        let n = digits.len();
        let mut digits = digits;
        digits.sort();
        let mut results = std::collections::BTreeSet::new();

        for i in 0..n {
            if digits[i] == 0 {
                continue;
            }
            for j in 0..n {
                if j == i {
                    continue;
                }
                for k in 0..n {
                    if k == i || k == j || digits[k] & 1 == 1 {
                        continue;
                    }

                    let res = digits[i] * 100 + digits[j] * 10 + digits[k];

                    results.insert(res);
                }
            }
        }

        results.into_iter().collect()
    }
}
