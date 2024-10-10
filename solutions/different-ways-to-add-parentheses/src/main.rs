fn main() {
    println!("{:?}", Solution::diff_ways_to_compute("2-1-1".to_string()));
    println!(
        "{:?}",
        Solution::diff_ways_to_compute("2*3-4*5".to_string())
    );
    println!("{:?}", Solution::diff_ways_to_compute("1".to_string()));
}

struct Solution;

impl Solution {
    pub fn diff_ways_to_compute(expression: String) -> Vec<i32> {
        Self::calculate(&expression)
    }

    fn calculate(expression: &str) -> Vec<i32> {
        if !(expression.contains("+") || expression.contains("-") || expression.contains("*")) {
            return vec![expression.parse::<i32>().unwrap()];
        }

        let n = expression.len();

        let mut results = Vec::new();
        for i in 0..n {
            let expr = expression.as_bytes()[i];
            if expr == b'*' || expr == b'+' || expr == b'-' {
                let left_expression = &expression[..i];
                let left_res = Self::calculate(left_expression);
                let right_expression = &expression[i + 1..];
                let right_res = Self::calculate(right_expression);

                for left_num in left_res {
                    for right_num in &right_res {
                        let res = match expr {
                            b'+' => left_num + right_num,
                            b'-' => left_num - right_num,
                            _ => left_num * right_num,
                        };

                        results.push(res);
                    }
                }
            }
        }

        results
    }
}
