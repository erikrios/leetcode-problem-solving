fn main() {
    println!("{}", Solution::lemonade_change(vec![5, 5, 5, 10, 20]));
    println!("{}", Solution::lemonade_change(vec![5, 5, 10, 10, 20]));
}

struct Solution;

impl Solution {
    pub fn lemonade_change(bills: Vec<i32>) -> bool {
        let mut five_dollars_count = 0;
        let mut ten_dollars_count = 0;

        for bill in bills {
            match bill {
                5 => five_dollars_count += 1,
                10 => {
                    ten_dollars_count += 1;
                    if five_dollars_count <= 0 {
                        return false;
                    }
                    five_dollars_count -= 1;
                }
                _ => {
                    if ten_dollars_count <= 0 {
                        if five_dollars_count < 3 {
                            return false;
                        }
                        five_dollars_count -= 3;
                    } else {
                        ten_dollars_count -= 1;
                        if five_dollars_count <= 0 {
                            return false;
                        }
                        five_dollars_count -= 1;
                    }
                }
            }
        }

        true
    }
}
