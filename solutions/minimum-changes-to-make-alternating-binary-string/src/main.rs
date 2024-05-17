fn main() {
    println!("{}", Solution::min_operations(String::from("0100")));
    println!("{}", Solution::min_operations(String::from("10")));
    println!("{}", Solution::min_operations(String::from("1111")));
}

struct Solution;

impl Solution {
    pub fn min_operations(s: String) -> i32 {
        let mut start_zero_counter = 0;
        let mut start_one_counter = 0;

        let mut start_zero_flipper = false;
        let mut start_one_flipper = true;

        for character in s.chars() {
            match character {
                '0' => {
                    if start_zero_flipper {
                        start_zero_counter += 1;
                    }

                    if start_one_flipper {
                        start_one_counter += 1;
                    }
                }
                _ => {
                    if !start_zero_flipper {
                        start_zero_counter += 1;
                    }

                    if !start_one_flipper {
                        start_one_counter += 1;
                    }
                }
            }

            start_zero_flipper = !start_zero_flipper;
            start_one_flipper = !start_one_flipper;
        }

        if start_zero_counter < start_one_counter {
            return start_zero_counter;
        }

        start_one_counter
    }
}
