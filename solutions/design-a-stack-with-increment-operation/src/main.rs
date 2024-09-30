use std::cell::RefCell;

fn main() {
    let custom_stack = CustomStack::new(3);
    custom_stack.push(1);
    custom_stack.push(2);
    println!("{}", custom_stack.pop());
    custom_stack.push(2);
    custom_stack.push(3);
    custom_stack.push(4);
    custom_stack.increment(5, 100);
    custom_stack.increment(2, 100);
    println!("{}", custom_stack.pop());
    println!("{}", custom_stack.pop());
    println!("{}", custom_stack.pop());
    println!("{}", custom_stack.pop());

    println!("=====");

    let custom_stack = CustomStack::new(2);
    custom_stack.push(34);
    println!("{}", custom_stack.pop());
    custom_stack.increment(8, 100);
    println!("{}", custom_stack.pop());
    custom_stack.increment(9, 91);
    custom_stack.push(63);
    println!("{}", custom_stack.pop());
    custom_stack.push(84);
    custom_stack.increment(10, 93);
    custom_stack.increment(6, 45);
    custom_stack.increment(10, 4);
}

struct CustomStack {
    max_size: RefCell<usize>,
    current_size: RefCell<usize>,
    nums: RefCell<Vec<i32>>,
}

impl CustomStack {
    fn new(max_size: i32) -> Self {
        Self {
            max_size: RefCell::new(max_size as usize),
            current_size: RefCell::new(0),
            nums: RefCell::new(Vec::with_capacity(max_size as usize)),
        }
    }

    fn push(&self, x: i32) {
        if self.is_full() {
            return;
        }

        self.nums.borrow_mut().push(x);
        *self.current_size.borrow_mut() += 1;
    }

    fn pop(&self) -> i32 {
        if self.is_empty() {
            return -1;
        }

        *self.current_size.borrow_mut() -= 1;
        self.nums.borrow_mut().pop().unwrap()
    }

    fn increment(&self, k: i32, val: i32) {
        if self.is_empty() {
            return;
        }
        for i in 0..k {
            if i as usize > *self.current_size.borrow() - 1 {
                break;
            }
            self.nums.borrow_mut()[i as usize] += val;
        }
    }

    fn is_empty(&self) -> bool {
        *self.current_size.borrow() == 0
    }

    fn is_full(&self) -> bool {
        *self.current_size.borrow() == *self.max_size.borrow()
    }
}
