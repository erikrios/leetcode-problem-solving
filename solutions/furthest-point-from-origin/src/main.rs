fn main() {
    println!(
        "{}",
        Solution::furthest_distance_from_origin(String::from("L_RL__R"))
    );
    println!(
        "{}",
        Solution::furthest_distance_from_origin(String::from("_R__LL_"))
    );
    println!(
        "{}",
        Solution::furthest_distance_from_origin(String::from("_______"))
    );
}

struct Solution;

impl Solution {
    pub fn furthest_distance_from_origin(moves: String) -> i32 {
        let mut total_left = 0;
        let mut total_right = 0;
        let mut total_empty = 0;

        for i in 0..moves.len() {
            let direction = moves.as_bytes()[i] as char;

            match direction {
                'L' => {
                    total_left += 1;
                }
                'R' => {
                    total_right += 1;
                }
                _ => {
                    total_empty += 1;
                }
            }
        }

        if total_right > total_left {
            total_right += total_empty;
        } else {
            total_left += total_empty;
        }

        let mut distance = total_right - total_left;
        if distance < 0 {
            distance *= -1;
        }

        distance
    }
}
