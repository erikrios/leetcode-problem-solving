fn main() {
    println!("{}", Solution::reverse_bits(43261596));
    println!("{}", Solution::reverse_bits(4294967293));
}

struct Solution;

impl Solution {
    pub fn reverse_bits(x: u32) -> u32 {
        let mut x = x;
        let mut bits = vec![false; 32];

        let mut i = 0;
        while x > 0 {
            bits[i] = x % 2 != 0;
            x /= 2;
            i += 1;
        }

        let mut result = 0;
        for (i, &bit) in bits.iter().rev().enumerate() {
            if bit {
                result += 2_u32.pow(i as u32);
            }
        }

        result
    }
}
