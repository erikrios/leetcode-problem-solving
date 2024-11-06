fn main() {
    println!("{}", Solution::can_sort_array(vec![8, 4, 2, 30, 15]));
    println!("{}", Solution::can_sort_array(vec![1, 2, 3, 4, 5]));
    println!("{}", Solution::can_sort_array(vec![3, 16, 8, 4, 2]));
}

struct Solution;

impl Solution {
    pub fn can_sort_array(nums: Vec<i32>) -> bool {
        let mut nums = nums;
        Solution::insertion_sort(&mut nums)
    }

    fn insertion_sort(nums: &mut [i32]) -> bool {
        let n = nums.len();

        for i in 1..n {
            let mut tmp_i = i;

            while tmp_i > 0 {
                let num = nums[tmp_i];
                let prev = nums[tmp_i - 1];

                if num < prev {
                    if Self::count_set_bits(num) != Self::count_set_bits(prev) {
                        return false;
                    }
                    let tmp = num;
                    nums[tmp_i] = prev;
                    nums[tmp_i - 1] = tmp;
                    tmp_i -= 1;
                } else {
                    break;
                }
            }
        }

        true
    }

    fn count_set_bits(num: i32) -> usize {
        let mut num = num;
        let mut counter = 0;

        while num > 0 {
            let digit = num % 2;
            if digit == 1 {
                counter += 1;
            }
            num /= 2;
        }

        counter
    }
}
