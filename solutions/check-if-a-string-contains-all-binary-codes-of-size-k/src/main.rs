fn main() {
    println!("{}", Solution::has_all_codes("00110110".to_string(), 2));
    println!("{}", Solution::has_all_codes("0110".to_string(), 1));
    println!("{}", Solution::has_all_codes("0110".to_string(), 2));
}

struct Solution;

impl Solution {
    pub fn has_all_codes(s: String, k: i32) -> bool {
        let mut codes = String::new();
        let mut all_codes = vec![false; (2 as i32).pow(k as u32) as usize];

        let s = s.chars();

        for (window_end, character) in s.enumerate() {
            codes.push(character);

            if window_end >= k as usize - 1 {
                let dec = Solution::binary_to_decimal(codes.chars().collect());
                if dec < all_codes.len() {
                    all_codes[dec] = true
                }
                codes.remove(0);
            }
        }

        !all_codes.contains(&false)
    }

    fn binary_to_decimal(binary: Vec<char>) -> usize {
        let mut res: usize = 0;

        for i in 0..binary.len() {
            if binary[binary.len() - 1 - i] == '1' {
                res += (2 as usize).pow(i as u32);
            }
        }

        res
    }
}
