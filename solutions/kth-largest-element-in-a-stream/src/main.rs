use std::cell::RefCell;

fn main() {
    let kth_largest = KthLargest::new(3, vec![4, 5, 8, 2]);
    println!("{}", kth_largest.add(3));
    println!("{}", kth_largest.add(5));
    println!("{}", kth_largest.add(10));
    println!("{}", kth_largest.add(9));
    println!("{}", kth_largest.add(4));

    println!("=====");

    let kth_largest = KthLargest::new(4, vec![7, 7, 7, 7, 8, 3]);
    println!("{}", kth_largest.add(2));
    println!("{}", kth_largest.add(10));
    println!("{}", kth_largest.add(9));
    println!("{}", kth_largest.add(9));
}

struct KthLargest {
    k: usize,
    sorted_nums: RefCell<Vec<i32>>,
}

impl KthLargest {
    fn new(k: i32, nums: Vec<i32>) -> Self {
        let mut nums = nums;
        nums.sort();
        Self {
            k: k as usize,
            sorted_nums: RefCell::new(nums.to_vec()),
        }
    }

    fn add(&self, val: i32) -> i32 {
        if self.sorted_nums.borrow().is_empty() {
            self.sorted_nums.borrow_mut().push(val);
            return val;
        }

        let mut updated_index = -1;
        for (i, &num) in self.sorted_nums.borrow().iter().enumerate() {
            if val <= num {
                updated_index = i as i32;
                break;
            }
        }

        if updated_index == -1 {
            self.sorted_nums.borrow_mut().push(val);
        } else {
            self.sorted_nums
                .borrow_mut()
                .insert(updated_index as usize, val);
        }

        let n = self.sorted_nums.borrow().len();

        self.sorted_nums.borrow_mut()[n - self.k]
    }
}
