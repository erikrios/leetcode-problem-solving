fn main() {
    println!("{}", Solution::parse_bool_expr("&(|(f))".to_string()));
    println!("{}", Solution::parse_bool_expr("|(f,f,f,t)".to_string()));
    println!("{}", Solution::parse_bool_expr("!(&(f,t))".to_string()));
    println!(
        "{}",
        Solution::parse_bool_expr("!(&(f,t,|(f,t,f),&(t,t,t)))".to_string())
    );
    println!(
        "{}",
        Solution::parse_bool_expr("|(&(t,f,t),!(t))".to_string())
    );
    println!(
        "{}",
        Solution::parse_bool_expr("!(&(&(!(&(f)),&(t),|(f,f,t)),&(t),&(t,t,f)))".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn parse_bool_expr(expression: String) -> bool {
        Solution::parse(&expression)
    }

    fn parse(expression: &str) -> bool {
        if expression == "t" {
            return true;
        } else if expression == "f" {
            return false;
        }

        let sub_expression = &expression[2..expression.len() - 1];

        let mut sub_expressions = Vec::new();
        let mut start = 0;
        let mut parentheses_counter = 0;

        for (end, &ch) in sub_expression.as_bytes().iter().enumerate() {
            if ch == b',' && parentheses_counter == 0 {
                sub_expressions.push(&sub_expression[start..end]);
                start = end + 1;
            } else if ch == b'(' {
                parentheses_counter += 1;
            } else if ch == b')' {
                parentheses_counter -= 1;
            }
        }

        sub_expressions.push(&sub_expression[start..sub_expression.len()]);

        let operand = expression.as_bytes()[0];

        let mut result = false;
        for (i, &sub_expr) in sub_expressions.iter().enumerate() {
            let res = Self::parse(sub_expr);
            match operand {
                b'!' => {
                    return !res;
                }
                val => match val {
                    b'|' => {
                        if i == 0 {
                            result = res;
                        } else {
                            result |= res;
                        }
                    }
                    b'&' => {
                        if i == 0 {
                            result = res;
                        } else {
                            result &= res;
                        }
                    }
                    _ => {
                        panic!("unrecognized operand");
                    }
                },
            }
        }

        result
    }
}
