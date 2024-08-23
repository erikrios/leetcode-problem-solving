fn main() {
    println!("{}", Solution::number_of_alternating_groups(vec![1, 1, 1]));
    println!(
        "{}",
        Solution::number_of_alternating_groups(vec![0, 1, 0, 0, 1])
    );
}

struct Solution;

impl Solution {
    pub fn number_of_alternating_groups(colors: Vec<i32>) -> i32 {
        let n = colors.len();
        let mut alt_group_num = 0;

        for (i, &color) in colors.iter().enumerate() {
            let x = color == 1;
            let y = (colors[(i + 1) % n]) == 1;
            let z = (colors[(i + 2) % n]) == 1;

            if Self::is_alternating_group(x, y, z) {
                alt_group_num += 1;
            }
        }

        alt_group_num
    }

    fn is_alternating_group(x: bool, y: bool, z: bool) -> bool {
        (x && !y && z) || (!x && y && !z)
    }
}
