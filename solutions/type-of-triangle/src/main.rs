use std::vec;

fn main() {
    println!("{}", Solution::triangle_type(vec![3, 3, 3]));
    println!("{}", Solution::triangle_type(vec![3, 4, 5]));
}

struct Solution;

impl Solution {
    pub fn triangle_type(nums: Vec<i32>) -> String {
        let mut nums = nums;
        nums.sort();

        let a = nums[0];
        let b = nums[1];
        let c = nums[2];

        if a + b <= c {
            return String::from("none");
        }

        let mut types = "scalene";
        if a == b && a == c {
            types = "equilateral";
        } else if a == b || a == c || b == c {
            types = "isosceles";
        }

        String::from(types)
    }
}
