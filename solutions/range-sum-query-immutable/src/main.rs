fn main() {
    let num_arr = NumArray::new(vec![-2, 0, 3, -5, 2, -1]);
    println!("{}", num_arr.sum_range(0, 2));
    println!("{}", num_arr.sum_range(2, 5));
    println!("{}", num_arr.sum_range(0, 5));
}

struct NumArray {
    nums: Vec<i32>,
}

impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        NumArray { nums }
    }

    fn sum_range(&self, left: i32, right: i32) -> i32 {
        let mut sum = 0;
        for i in left..=right {
            sum += self.nums[i as usize];
        }
        sum
    }
}
