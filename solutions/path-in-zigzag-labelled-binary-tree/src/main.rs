fn main() {
    println!("{:?}", Solution::path_in_zig_zag_tree(14));
    println!("{:?}", Solution::path_in_zig_zag_tree(26));
    println!("{:?}", Solution::path_in_zig_zag_tree(1));
}

struct Solution;

impl Solution {
    pub fn path_in_zig_zag_tree(label: i32) -> Vec<i32> {
        let mut visited_labels = Vec::new();
        visited_labels.push(label);

        let mut label = label;

        while label > 1 {
            let level = ((label * 2) as f32).log2() as i32;
            let starting_point = 2_i32.pow(level as u32 - 1);
            let ending_point = 2_i32.pow(level as u32) - 1;
            let should_reverse = level % 2 == 0;

            let actual_index = if should_reverse {
                ending_point - label + starting_point - 1
            } else {
                label - 1
            };

            let parent_index = (actual_index - 1) / 2;
            let parent_should_reverse = !should_reverse;
            label = if parent_should_reverse {
                let starting_point = 2_i32.pow(level as u32 - 1 - 1);
                let ending_point = 2_i32.pow(level as u32 - 1) - 1;
                ending_point - (parent_index - (starting_point - 1))
            } else {
                parent_index + 1
            };
            visited_labels.push(label);
        }

        visited_labels.reverse();

        visited_labels
    }
}
