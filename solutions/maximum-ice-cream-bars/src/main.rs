fn main() {
    println!("{}", Solution::max_ice_cream(vec![1, 3, 2, 4, 1], 7));
    println!("{}", Solution::max_ice_cream(vec![10, 6, 8, 7, 7, 8], 5));
    println!("{}", Solution::max_ice_cream(vec![1, 6, 3, 1, 2, 5], 20));
}

struct Solution;

impl Solution {
    pub fn max_ice_cream(costs: Vec<i32>, coins: i32) -> i32 {
        let mut costs = costs;

        Self::counting_sort(&mut costs);

        let mut counter = 0;

        let mut coins = coins;
        for cost in costs {
            coins -= cost;
            if coins < 0 {
                break;
            }
            counter += 1;
        }

        counter
    }

    fn counting_sort(nums: &mut [i32]) {
        let mut max = 0;

        for &num in &*nums {
            max = max.max(num);
        }

        let mut counts = vec![0_usize; max as usize + 1];

        for &num in &*nums {
            counts[num as usize] += 1;
        }

        let n = counts.len();
        for i in 1..n {
            counts[i] += counts[i - 1];
        }

        let tmp_nums = nums.to_vec();

        for i in tmp_nums {
            counts[i as usize] -= 1;
            nums[counts[i as usize]] = i;
        }
    }
}
