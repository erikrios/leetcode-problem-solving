use std::collections::HashMap;

fn main() {
    println!("{:?}", Solution::letter_combinations("23".to_string()));
    println!("{:?}", Solution::letter_combinations("".to_string()));
    println!("{:?}", Solution::letter_combinations("2".to_string()));
    println!("{:?}", Solution::letter_combinations("2359".to_string()));
}

struct Solution;

impl Solution {
    pub fn letter_combinations(digits: String) -> Vec<String> {
        if digits.is_empty() {
            return vec![];
        }

        let mut digits_mapper: HashMap<u8, Vec<u8>> = HashMap::new();
        digits_mapper.insert(b'2', vec![b'a', b'b', b'c']);
        digits_mapper.insert(b'3', vec![b'd', b'e', b'f']);
        digits_mapper.insert(b'4', vec![b'g', b'h', b'i']);
        digits_mapper.insert(b'5', vec![b'j', b'k', b'l']);
        digits_mapper.insert(b'6', vec![b'm', b'n', b'o']);
        digits_mapper.insert(b'7', vec![b'p', b'q', b'r', b's']);
        digits_mapper.insert(b'8', vec![b't', b'u', b'v']);
        digits_mapper.insert(b'9', vec![b'w', b'x', b'y', b'z']);

        let mut results = Vec::new();

        Self::dfs(&digits, &digits_mapper, 0, &mut Vec::new(), &mut results);

        results
    }

    fn dfs(
        digits: &String,
        digits_mapper: &HashMap<u8, Vec<u8>>,
        i: usize,
        result: &mut Vec<u8>,
        results: &mut Vec<String>,
    ) {
        if i == digits.len() {
            results.push(String::from_utf8(result.to_vec()).unwrap());
            return;
        }

        let digit = digits.as_bytes()[i];
        let chars = digits_mapper.get(&digit).unwrap();

        for &ch in chars {
            result.push(ch);
            Self::dfs(digits, digits_mapper, i + 1, result, results);
            result.pop();
        }
    }
}
