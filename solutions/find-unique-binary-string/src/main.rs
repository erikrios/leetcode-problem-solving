use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::find_different_binary_string(vec!["01".to_string(), "10".to_string()])
    );
    println!(
        "{}",
        Solution::find_different_binary_string(vec!["00".to_string(), "01".to_string()])
    );
    println!(
        "{}",
        Solution::find_different_binary_string(vec![
            "111".to_string(),
            "011".to_string(),
            "001".to_string()
        ])
    );
    println!(
        "{}",
        Solution::find_different_binary_string(vec!["0".to_string()])
    );
}

struct Solution;

impl Solution {
    pub fn find_different_binary_string(nums: Vec<String>) -> String {
        let n = nums[0].len();
        let hash_set = HashSet::from_iter(nums);

        let (not_found, _) = Self::generate_binary_string(n, n, &hash_set);

        not_found
    }

    fn generate_binary_string(
        n: usize,
        n_target: usize,
        hash_set: &HashSet<String>,
    ) -> (String, Vec<String>) {
        if n == 1 {
            if n_target == n {
                if !hash_set.contains("1") {
                    return (String::from("1"), vec![]);
                } else {
                    return (String::from("0"), vec![]);
                }
            }
            return (String::new(), vec!["0".to_string(), "1".to_string()]);
        }

        let (mut not_found, res) = Self::generate_binary_string(n - 1, n_target, hash_set);

        let mut results = Vec::with_capacity(2 * res.len());

        for v in res {
            let mut a = String::from("0");
            a.push_str(&v);
            if n == n_target && !hash_set.contains(&a) {
                not_found.clear();
                not_found.push_str(a.as_str());
                break;
            }
            results.push(a);
            let mut b = String::from("1");
            b.push_str(&v);
            if n == n_target && !hash_set.contains(&b) {
                not_found.clear();
                not_found.push_str(b.as_str());
                break;
            }
            results.push(b);
        }

        (not_found, results)
    }
}
