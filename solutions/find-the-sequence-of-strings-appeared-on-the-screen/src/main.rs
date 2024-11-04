fn main() {
    println!("{:?}", Solution::string_sequence("abc".to_string()));
    println!("{:?}", Solution::string_sequence("he".to_string()));
}

struct Solution;

impl Solution {
    pub fn string_sequence(target: String) -> Vec<String> {
        let target = target.into_bytes();
        let mut typed = Vec::new();
        let mut results = Vec::new();

        for ch in target {
            typed.push(b'a');
            while typed[typed.len() - 1] != ch {
                results.push(String::from_utf8(typed.clone()).unwrap());
                let n = typed.len();
                typed[n - 1] = b'a' + ((typed[typed.len() - 1] - b'a' + 1) % 26);
            }

            results.push(String::from_utf8(typed.clone()).unwrap());
        }

        results
    }
}
