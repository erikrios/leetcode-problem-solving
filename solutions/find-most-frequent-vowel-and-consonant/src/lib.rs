pub struct Solution;

impl Solution {
    pub fn max_freq_sum_two_loops(s: String) -> i32 {
        const N: usize = 26;

        let mut counters = [0_usize; N];

        for ch in s.into_bytes() {
            counters[(ch - b'a') as usize] += 1;
        }

        let mut max_vocal = usize::MIN;
        let mut max_consonant = usize::MIN;
        for (i, counter) in counters.into_iter().enumerate() {
            let ch = b'a' + i as u8;
            if ch == b'a' || ch == b'i' || ch == b'u' || ch == b'e' || ch == b'o' {
                max_vocal = max_vocal.max(counter);
                continue;
            }

            max_consonant = max_consonant.max(counter);
        }

        (max_vocal + max_consonant) as i32
    }

    pub fn max_freq_sum_one_loop(s: String) -> i32 {
        const N: usize = 26;

        let mut counters = [0_usize; N];
        let mut max_vowel = 0;
        let mut max_consonant = 0;

        for ch in s.bytes() {
            let i = (ch - b'a') as usize;
            counters[i] += 1;

            if ch == b'a' || ch == b'e' || ch == b'i' || ch == b'o' || ch == b'u' {
                max_vowel = max_vowel.max(counters[i]);
            } else {
                max_consonant = max_consonant.max(counters[i]);
            }
        }

        (max_vowel + max_consonant) as i32
    }
}
