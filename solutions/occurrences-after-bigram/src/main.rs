fn main() {
    println!(
        "{:?}",
        Solution::find_ocurrences(
            "alice is a good girl she is a good student".to_string(),
            "a".to_string(),
            "good".to_string()
        )
    );
    println!(
        "{:?}",
        Solution::find_ocurrences(
            "we will we will rock you".to_string(),
            "we".to_string(),
            "will".to_string()
        )
    );
    println!(
        "{:?}",
        Solution::find_ocurrences(
            "alice is aa good girl she is a good student".to_string(),
            "a".to_string(),
            "good".to_string()
        )
    );
    println!(
        "{:?}",
        Solution::find_ocurrences(
            "jkypmsxd jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa jkypmsxd kcyxdfnoa kcyxdfnoa jkypmsxd kcyxdfnoa".to_string(),
            "kcyxdfnoa".to_string(),
            "jkypmsxd".to_string()
        )
    );
    println!(
        "{:?}",
        Solution::find_ocurrences(
            "we we we we will rock you".to_string(),
            "we".to_string(),
            "we".to_string()
        )
    );
}

struct Solution;

impl Solution {
    pub fn find_ocurrences(text: String, first: String, second: String) -> Vec<String> {
        let mut results = Vec::new();
        let mut prevs: Vec<&str> = text.split(' ').take(2).collect();
        for val in text.split(' ').skip(2) {
            if first == prevs[0] && second == prevs[1] {
                results.push(val);
            }
            prevs[0] = prevs[1];
            prevs[1] = val;
        }

        results.into_iter().map(|word| word.to_string()).collect()
    }
}
