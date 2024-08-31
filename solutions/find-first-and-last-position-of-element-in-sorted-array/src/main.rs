fn main() {
    println!("{:?}", Solution::search_range(vec![5, 7, 7, 8, 8, 10], 8));
    println!("{:?}", Solution::search_range(vec![5, 7, 7, 8, 8, 10], 6));
    println!("{:?}", Solution::search_range(vec![], 0));
    println!("{:?}", Solution::search_range(vec![1], 1));
    println!("{:?}", Solution::search_range(vec![1], 0));
}

struct Solution;

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let n = nums.len();
        let mut start_pos = -1;
        let mut end_pos = -1;

        if n == 0 {
            return vec![start_pos, end_pos];
        }

        let mut start = 0_i32;
        let mut end = n as i32 - 1;

        while start <= end {
            let mid = (start + end) / 2;
            let num = nums[mid as usize];

            match num.cmp(&target) {
                std::cmp::Ordering::Greater => end = mid - 1,
                std::cmp::Ordering::Less => start = mid + 1,
                std::cmp::Ordering::Equal => {
                    start_pos = mid;
                    end_pos = mid;
                    break;
                }
            }
        }

        if start_pos == -1 {
            return vec![start_pos, end_pos];
        }

        let mut prev_start_pos = start_pos - 1;
        let mut next_end_pos = end_pos + 1;

        while prev_start_pos >= 0 || next_end_pos < n as i32 {
            let mut should_break = true;

            if prev_start_pos >= 0 && nums[prev_start_pos as usize] == target {
                start_pos = prev_start_pos;
                should_break = false;
            }

            if next_end_pos < n as i32 && nums[next_end_pos as usize] == target {
                end_pos = next_end_pos;
                should_break = false;
            }

            prev_start_pos -= 1;
            next_end_pos += 1;

            if should_break {
                break;
            }
        }

        vec![start_pos, end_pos]
    }
}
