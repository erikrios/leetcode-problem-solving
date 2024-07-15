fn main() {
    println!("{}", Solution::reverse_parentheses("(abcd)".to_string()));
    println!(
        "{}",
        Solution::reverse_parentheses("(u(love)i)".to_string())
    );
    println!(
        "{}",
        Solution::reverse_parentheses("(ed(et(oc))el)".to_string())
    );
    println!(
        "{}",
        Solution::reverse_parentheses("(ed(et(oc)e)el)".to_string())
    );
    println!("{}", Solution::reverse_parentheses("".to_string()));
    println!(
        "{}",
        Solution::reverse_parentheses("ta()usw((((a))))".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn reverse_parentheses(s: String) -> String {
        let s: Vec<char> = s.chars().collect();
        let mut stack = Vec::<usize>::new();
        let mut reversed = String::new();

        for &ch in s.iter() {
            if ch == '(' {
                stack.push(reversed.len());
            } else if ch == ')' {
                let start_index = stack.pop().unwrap();
                let last_index = reversed.len() - 1;
                Self::reverse(start_index, last_index, &mut reversed);
            } else {
                reversed.push(ch);
            }
        }

        reversed
    }

    fn reverse(mut start_index: usize, mut last_index: usize, result: &mut String) {
        while start_index < last_index {
            let tmp = result.chars().nth(start_index).unwrap();
            result.replace_range(
                start_index..start_index + 1,
                &result.chars().nth(last_index).unwrap().to_string(),
            );
            result.replace_range(last_index..last_index + 1, &tmp.to_string());
            start_index += 1;
            last_index -= 1;
        }
    }
}
