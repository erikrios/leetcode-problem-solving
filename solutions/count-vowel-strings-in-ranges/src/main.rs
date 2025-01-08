fn main() {
    println!(
        "{:?}",
        Solution::vowel_strings(
            vec![
                "aba".to_string(),
                "bcb".to_string(),
                "ece".to_string(),
                "aa".to_string(),
                "e".to_string()
            ],
            vec![vec![0, 2], vec![1, 4], vec![1, 1],]
        )
    );

    println!(
        "{:?}",
        Solution::vowel_strings(
            vec!["a".to_string(), "e".to_string(), "i".to_string(),],
            vec![vec![0, 2], vec![0, 1], vec![2, 2],]
        )
    );
}

struct Solution;

const VOWELS: [u8; 5] = [b'a', b'e', b'i', b'o', b'u'];

impl Solution {
    pub fn vowel_strings(words: Vec<String>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let n = words.len();
        let mut vowel_counter = 0;
        let mut vowel_counters = Vec::with_capacity(n);

        for word in words {
            let word = word.into_bytes();
            let start_ch = word[0];
            let end_ch = word[word.len() - 1];

            if VOWELS.contains(&start_ch) && VOWELS.contains(&end_ch) {
                vowel_counter += 1;
            }

            vowel_counters.push(vowel_counter);
        }

        let mut results = Vec::with_capacity(queries.len());

        for query in queries {
            let i_start = query[0];
            let i_end = query[1];

            let mut total = vowel_counters[i_end as usize];
            if i_start > 0 {
                total -= vowel_counters[(i_start - 1) as usize];
            }

            results.push(total);
        }

        results
    }
}
