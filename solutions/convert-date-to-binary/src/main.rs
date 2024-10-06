use core::str;

fn main() {
    println!(
        "{}",
        Solution::convert_date_to_binary("2080-02-29".to_string())
    );
    println!(
        "{}",
        Solution::convert_date_to_binary("1900-01-01".to_string())
    );
}

#[test]
fn test_convert_date_to_binary() {
    assert_eq!(
        Solution::convert_date_to_binary("2080-02-29".to_string()),
        "100000100000-10-11101".to_string()
    );
    assert_eq!(
        Solution::convert_date_to_binary("1900-01-01".to_string()),
        "11101101100-1-1".to_string()
    );
}

struct Solution;

impl Solution {
    pub fn convert_date_to_binary(date: String) -> String {
        let mut num_str = Vec::new();
        let mut res = Vec::new();

        for &ch_byte in date.as_bytes() {
            if ch_byte == b'-' {
                if num_str[0] == b'0' {
                    num_str.remove(0);
                }

                let num = Self::utf8_to_num(&num_str);
                Self::num_to_binary(&mut res, num);

                res.push(b'-');
                num_str.clear();
            } else {
                num_str.push(ch_byte);
            }
        }

        let num = Self::utf8_to_num(&num_str);
        Self::num_to_binary(&mut res, num);

        str::from_utf8(&res).unwrap().to_string()
    }

    fn utf8_to_num(utf8: &[u8]) -> u16 {
        str::from_utf8(utf8).unwrap().parse::<u16>().unwrap()
    }

    fn num_to_binary(res: &mut Vec<u8>, mut num: u16) {
        let mut binary_res = Vec::new();

        while num > 0 {
            let binary_digit = (num % 2) as u8;
            if binary_digit == 1 {
                binary_res.insert(0, b'1');
            } else {
                binary_res.insert(0, b'0');
            }
            num /= 2;
        }

        res.append(&mut binary_res)
    }
}
