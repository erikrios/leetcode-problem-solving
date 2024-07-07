fn main() {
    println!("{}", Solution::num_water_bottles(9, 3));
    println!("{}", Solution::num_water_bottles(15, 4));
    println!("{}", Solution::num_water_bottles(15, 8));
}

struct Solution;

impl Solution {
    pub fn num_water_bottles(num_bottles: i32, num_exchange: i32) -> i32 {
        let mut num_bottles = num_bottles;
        let mut remaining_empty = 0;
        let mut drinked = 0;

        loop {
            drinked += num_bottles;
            remaining_empty += num_bottles;
            if remaining_empty < num_exchange {
                break;
            }
            num_bottles = remaining_empty / num_exchange;
            remaining_empty %= num_exchange;
        }

        drinked
    }
}
