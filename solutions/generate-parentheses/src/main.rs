fn main() {
    println!("{:?}", Solution::generate_parenthesis(3));
    println!("{:?}", Solution::generate_parenthesis(1));
    println!("{:?}", Solution::generate_parenthesis(2));
    println!("{:?}", Solution::generate_parenthesis(4));
}

struct Solution;

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut res = Vec::new();

        let mut char_arr = vec![0; n as usize + n as usize];
        Self::generate(n, n, &mut char_arr, 0, &mut res);

        res
    }

    fn generate(open: i32, close: i32, arr: &mut Vec<u8>, n: i32, res: &mut Vec<String>) {
        if open == 0 && close == 0 {
            res.push(String::from_utf8(arr.to_vec()).unwrap());
            return;
        }

        if open > 0 {
            arr[n as usize] = b'(';
            Self::generate(open - 1, close, arr, n + 1, res);
        }

        if close > open {
            arr[n as usize] = b')';
            Self::generate(open, close - 1, arr, n + 1, res);
        }
    }
}
