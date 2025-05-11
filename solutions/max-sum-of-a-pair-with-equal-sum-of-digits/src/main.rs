use std::collections::HashMap;

fn main() {
    println!("{}", Solution::maximum_sum(vec![18, 43, 36, 13, 7]));
    println!("{}", Solution::maximum_sum(vec![10, 12, 19, 14]));
}

struct Solution;

impl Solution {
    pub fn maximum_sum(nums: Vec<i32>) -> i32 {
        let mut mapper = HashMap::new();
        let mut max = -1;

        for num in nums {
            let digit_sum = Self::digit_sum(num as usize);
            mapper
                .entry(digit_sum)
                .and_modify(|nums_vec: &mut Vec<i32>| {
                    if nums_vec.len() == 1 {
                        if num > nums_vec[0] {
                            nums_vec.insert(0, num);
                        } else {
                            nums_vec.push(num);
                        }
                    } else if num > nums_vec[0] {
                        nums_vec.pop();
                        nums_vec.insert(0, num);
                    } else if num > nums_vec[1] {
                        nums_vec.pop();
                        nums_vec.insert(1, num);
                    }

                    let val_sum = nums_vec[0] + nums_vec[1];
                    max = max.max(val_sum);
                })
                .or_insert_with(|| {
                    let mut nums_vec = Vec::with_capacity(2);
                    nums_vec.push(num);
                    nums_vec
                });
        }

        max
    }

    fn digit_sum(num: usize) -> usize {
        let mut num = num;
        let mut sum = 0;

        while num > 0 {
            let digit = num % 10;
            sum += digit;
            num /= 10;
        }

        sum
    }
}
