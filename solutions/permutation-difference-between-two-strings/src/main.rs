use std::collections::HashMap;

fn main() {
    println!("{}", Solution::find_permutation_difference(String::from("abc"), String::from("bac")));
    println!("{}", Solution::find_permutation_difference(String::from("abcde"), String::from("edbac")));
}

struct Solution;

impl Solution {
    pub fn find_permutation_difference(s: String, t: String) -> i32 {
        let mut mapper: HashMap<u8, [usize;2]> = HashMap::new();

        let n = s.len();

        for i in 0..n {
            let char_in_s = s.as_bytes()[i];
            let char_in_t = t.as_bytes()[i];

            match mapper.get_mut(&char_in_s) {
                Some(vals) => vals[0] = i,
                None => {
                    mapper.insert(char_in_s, [i,0]);
                }
            }

            match mapper.get_mut(&char_in_t) {
                Some(vals) => vals[1] = i,
                None => {
                    mapper.insert(char_in_t, [0,i]);
                }
            }
        }

        let mut sum = 0;

        for (_, values) in mapper.iter() {
            let diff = if values[0] > values[1] {
                values[0] - values[1]
            } else {
                values[1] - values[0]
            };
            sum += diff as i32;
        }

        sum
    }
}
