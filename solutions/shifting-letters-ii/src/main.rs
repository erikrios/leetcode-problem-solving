fn main() {
    println!(
        "{}",
        Solution::shifting_letters(
            "abc".to_string(),
            vec![vec![0, 1, 0], vec![1, 2, 1], vec![0, 2, 1]]
        )
    );
    println!(
        "{}",
        Solution::shifting_letters("dztz".to_string(), vec![vec![0, 0, 0], vec![1, 1, 1]])
    );
}

struct Solution;

impl Solution {
    pub fn shifting_letters(s: String, shifts: Vec<Vec<i32>>) -> String {
        let n = s.len();
        let mut s = s.into_bytes();
        let mut prefix_sum = vec![0; n + 1];

        for shift in shifts {
            let i_start = shift[0] as usize;
            let i_end = shift[1] as usize;
            let is_forward = shift[2] == 1;
            let counter: i32 = if is_forward { 1 } else { -1 };

            prefix_sum[i_end + 1] += counter;
            prefix_sum[i_start] -= counter;
        }

        let mut diff = 0;
        for (i, pref_sum) in prefix_sum.into_iter().enumerate().rev() {
            if i == 0 {
                continue;
            }

            diff += pref_sum;
            let net_shift = (diff % 26 + 26) % 26;

            s[i - 1] = ((s[i - 1] - b'a' + net_shift as u8) % 26) + b'a';
        }

        String::from_utf8(s).unwrap()
    }
}
