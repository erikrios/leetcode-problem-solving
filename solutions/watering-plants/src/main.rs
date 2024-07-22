fn main() {
    println!("{}", Solution::watering_plants(vec![2, 2, 3, 3], 5));
    println!("{}", Solution::watering_plants(vec![1, 1, 1, 4, 2, 3], 4));
    println!(
        "{}",
        Solution::watering_plants(vec![7, 7, 7, 7, 7, 7, 7], 8)
    );
}

struct Solution;

impl Solution {
    pub fn watering_plants(plants: Vec<i32>, capacity: i32) -> i32 {
        let mut current_capacity = capacity;
        let mut steps = 0;

        for (i, &amount) in plants.iter().enumerate() {
            if current_capacity >= amount {
                steps += 1_i32;
                current_capacity -= amount;
            } else {
                steps += i as i32;
                current_capacity = capacity;
                steps += i as i32 + 1;
                current_capacity -= amount;
            }
        }

        steps
    }
}
